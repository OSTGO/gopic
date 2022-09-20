package plugin

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/OSTGO/gopic/conf"
	"github.com/OSTGO/gopic/utils"
)

type GithubStorage struct {
	*utils.MetaStorage
}

const (
	githubPluginName = "github"
)

var githubConfig map[string]interface{}

func (g *GithubStorage) Upload(im *utils.Image) (string, error) {
	responseURL := githubConfig["responseurl"].(string)
	requestURL := githubConfig["requesturl"].(string)
	token := githubConfig["token"].(string)
	return responseURL + im.OutSuffix, uploadPictureToGithub(requestURL, token, im.OutBase64, im.OutSuffix)
}

func NewGithubStorage() *GithubStorage {
	return &GithubStorage{utils.NewMetaStorage()}
}

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

func init() {
	utils.StroageHelp[githubPluginName] = githubHelp()
	githubConfig = conf.Viper.GetStringMap(githubPluginName)
	if githubConfig == nil {
		return
	}
	active := githubConfig["active"]
	if active == nil {
		return
	}
	if active == true {
		utils.StroageMap[githubPluginName] = NewGithubStorage()
	}
}

func githubHelp() string {
	return "github plugin need this parameters:\nactive: false or true\nresponseURL: like https://gcore.jsdelivr.net/gh/yourUserName/pics@main/\nrequestURL: like https://api.github.com/repos/yourUserName/pics/contents/\ntoke: your github token"
}
