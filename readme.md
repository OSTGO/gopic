# 简介

gopic 是一个图床工具，用来上传到七牛云或者github 等，使用简单，容易开发插件，全并发，充分解耦

# 使用

## 交互

| gopic命令 | 说明                                     |
| --------- | ---------------------------------------- |
| init      | 初始化配置，配置生成在`home/.gopic.json` |
| conf      | 打印配置信息等                           |
| update    | 从github更新自己                         |
| upload    | 图床主命令，下表有详细参数介绍           |
| convert   | 图片转换命令                             |

| gopic upload 参数 | 说明                                                  | 默认值                             |
| ----------------- | ----------------------------------------------------- | ---------------------------------- |
| -a，--all         | bool类型，若为true,将上传到配置中激活的所有存储插件中 | false，若为false，-s内容为空会报错 |
| -o, --out         | string类型，选用哪一个存储插件作为返回值              | 默认为第一个存储                   |
| -s, --storage     | string列表，选择要上传到哪些存储插件中，用`,`分割     | 默认为空，若为空，-a未配置会报错   |
| -p, --path        | string列表，图片可以是本地地址，也可以是网络地址      | 必选值                             |

例如`gopic upload -a -p ./1.gif ./2.png https://baidu.com/img.png -o qiniu`

代表上传` ./1.gif ./2.png https://baidu.com/img.png `三个图片到所有已经激活的插件，并使用`qiniu`插件的返回值

结果为:

`https://pic.longtao.fun/pics/22/8520716594215125117496217642356398137175_1.png
https://pic.longtao.fun/pics/22/230182152169544190214107228101001761655235_2.png
https://pic.longtao.fun/pics/22/254555454555441902154101556543454354543545_img.png
`

其中22为年份，文件名为图片内容的md5结果与原本文件名使用`_`连接，能有效去重

| gopic convert 参数 | 说明                                                  | 默认值                             |
| ------------------ | ----------------------------------------------------- | ---------------------------------- |
| -c, --covertPath   | string类型，要操作的文件夹或文件                      | 当前文件夹                         |
| -a, --all          | bool类型，若为true,将上传到配置中激活的所有存储插件中 | false，若为false，-s内容为空会报错 |
| -r, --recurse      | bool类型，递归替换                                    | 默认递归(非递归没实现)             |
| -s, --storage      | string列表，选择要上传到哪些存储插件中，用`,`分割     | 默认为空，若为空，-a未配置会报错   |
| -f, --format       | string类型，选用哪一个存储插件作为返回值              | 默认为第一个存储                   |
| -d, --dir          | string类型, 替换后输出的位置                          | 必选                               |

例如，如命令: `./gopic-linux-amd64 convert -c /home/longtao/temp/blog/ -d /home/longtao/temp/blog2 -f samba -s samba`

将会把`/home/longtao/temp/blog/`目录下递归的把所有md文件中的图片转换到samb存储中，并把转换后的md文件存储在 `/home/longtao/temp/blog2`中

对比`tree /home/longtao/temp/blog/ ` 和`tree /home/longtao/temp/blog2 ` 我们发现是一样的，但是每个md文件中的图片都转换到samb存储中路

![image-20220805150255551](https://pic.longtao.fun/pics/22/17699932065516729233109210501292428820286_image-20220805150255551.png)

## 结合Typora使用

1.
进入`文件->偏好设置`![image-20220720161734836](https://pic.longtao.fun/pics/22/11820711612046216230142202241441926412985_image-20220720161734836.png)
2. 在`偏好配置->图像`选项卡中选择插入图片时上传图片，并使用`Custom Command`
   自定义上传服务，在命令中填写需要的gopic命令即可![image-20220720162157408](https://pic.longtao.fun/pics/22/111222126119019917210931716018085239178195_image-20220720162157408.png)

3. 使用验证图片上传成功后，即可开心的任意粘贴图片，文章中图片会自动保存到激活插件的存储中，本说明的图片就使用gopic自动上出处理

# 开发

## 添加插件

插件需要添加到`plugin`目录中，需要满足以下条件

1. 定义这个插件的类，其中要包含基础类`*utils.BaseStorage`；
2. 实现`Upload(im *utils.Image)(string,error)`方法,在`Upload`中实现单个图片的上传，返回值分别为：图片返回地址、错误信息；
3. 定义初始化函数`init()`
4. 将插件名和帮助文档写入`utils.StroageHelp`
5. 将插件名和实例写入`utils.StroageMap` ,注意需要判断是否`active`

例子：

```go
package plugin

import (
	"bytes"
	"encoding/json"
	"gopic/conf"
	"gopic/utils"
	"io/ioutil"
	"net/http"
)
//定义插件类
type GithubStorage struct {
   *utils.MetaStorage
}
//定义插件名
const (
	githubPluginName = "github"
)
//声明插件配置
var githubConfig map[string]interface{}

//实现Upload方法
func (g *GithubStorage) Upload(im *utils.Image) (string, error) {
	responseURL := githubConfig["responseurl"].(string)
	requestURL := githubConfig["requesturl"].(string)
	token := githubConfig["token"].(string)
	return responseURL + im.OutSuffix, uploadPictureToGithub(requestURL, token, im.OutBase64, im.OutSuffix)
}

//实现插件实例构造函数
func NewGithubStorage() *GithubStorage {
    return &GithubStorage{utils.NewMetaStorage()}
}

//upload具体的实现，用来上传单个图片
func uploadPictureToGithub(requestURL, token, data, suffix string) error {
	url := requestURL + suffix
	body := make(map[string]interface{})
	body["branch"] = "main"
	body["message"] = "markdown upload picture"
	body["content"] = data
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	return err
}
//初始化函数
func init() {
    //将帮助文档写入
	utils.StroageHelp[githubPluginName] = githubHelp()
    //读取插件配置
	githubConfig = conf.Viper.GetStringMap(githubPluginName)
    //判断插件是否在配置文件中存在
	if githubConfig == nil {
		return
	}
    //判断配置中的插件是否处于激活状态
	active := githubConfig["active"]
	if active == nil {
		return
	}
    //如果插件处于激活状态，将插件的实例传入utils.StroageMap
	if active == true {
		utils.StroageMap[githubPluginName] = NewGithubStorage()
	}
}
//帮助文档具体实现
func githubHelp() string {
	return "github plugin need this parameters:\nactive: false or true\nresponseURL: like https://gcore.jsdelivr.net/gh/yourUserName/pics@main/\nrequestURL: like https://api.github.com/repos/yourUserName/pics/contents/\ntoke: your github token"
}
```

## 基本框架

### 全并行

不同图片之间以及不同插件之间全部并行
单例模式：相同图片只会处理一次，节省内存及时间消耗，对大目录来说能快速转换！

`使用一个插件传输一张图片`的时间T11与`使用n个插件传输n张图片`的时间Tnn, 在算力满足的条件下，Tnn远远小于`T11*n`,更远远小于`T11*n*n`

### 充分解耦

插件与框架完全解耦，可自行编写插件
