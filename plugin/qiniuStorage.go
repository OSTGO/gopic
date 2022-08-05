package plugin

import (
	"bytes"
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"gopic/conf"
	"gopic/utils"
)

const (
	qiniuPluginName = "qiniu"
)

var qiniuConfig map[string]interface{}

type QiniuStorage struct {
	*utils.MetaStorage
}

func (g *QiniuStorage) Upload(im *utils.Image) (string, error) {
	accessKey := qiniuConfig["accesskey"].(string)
	secretKey := qiniuConfig["secretkey"].(string)
	bucket := qiniuConfig["bucket"].(string)
	responseURL := qiniuConfig["responseurl"].(string)
	return responseURL + im.OutSuffix, uploadPictureToQiniu(accessKey, secretKey, bucket, im.OutBytes, im.OutSuffix)
}

func NewQiniuStorage() *QiniuStorage {
	return &QiniuStorage{utils.NewMetaStorage()}
}

func uploadPictureToQiniu(accessKey, secretKey, bucket string, data []byte, suffix string) error {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	//cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	//cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	//cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		//Params: map[string]string{
		//	"x:name": "github logo",
		//},
	}
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, "pics/"+suffix, bytes.NewReader(data), dataLen, &putExtra)
	return err
}

func init() {
	utils.StroageHelp[qiniuPluginName] = qiniuHelp()
	qiniuConfig = conf.Viper.GetStringMap(qiniuPluginName)
	if qiniuConfig == nil {
		return
	}
	active := qiniuConfig["active"]
	if active == nil {
		return
	}
	if active == true {
		utils.StroageMap[qiniuPluginName] = NewQiniuStorage()
	}
}

func qiniuHelp() string {
	return "accessKey,secretKey,bucket,responseURL"
}
