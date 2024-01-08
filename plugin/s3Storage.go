package plugin

import (
	"bytes"
	"fmt"
	"github.com/OSTGO/gopic/conf"
	"github.com/OSTGO/gopic/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Storage struct {
	*utils.MetaStorage
}

const (
	s3PluginName = "s3"
)

var s3Config map[string]interface{}

func (g *S3Storage) Upload(im *utils.Image) (string, error) {
	accessKey := s3Config["access_key"].(string)
	secretKey := s3Config["secret_key"].(string)
	endpoint := s3Config["endpoint"].(string)
	sslEnable := s3Config["disable_ssl"].(bool)
	forcePathStyle := s3Config["force_path_style"].(bool)
	prefix := s3Config["prefix"].(string)
	bucket := s3Config["bucket"].(string)
	region := s3Config["region"].(string)
	if len(bucket) == 0 {
		return "", fmt.Errorf("bucket name is empty")
	}
	s3Session, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:         aws.String(endpoint),
		DisableSSL:       aws.Bool(sslEnable),
		S3ForcePathStyle: aws.Bool(forcePathStyle),
		Region:           aws.String(region),
	})
	if err != nil {
		return "", err
	}
	s3Filename := im.OutSuffix
	if len(prefix) != 0 {
		s3Filename = fmt.Sprintf("%s/%s", prefix, im.OutSuffix)
	}
	imageUrl := fmt.Sprintf("s3://%s/%s/%s", endpoint, bucket, s3Filename)
	err = uploadPictureToS3(im, s3Session, bucket, s3Filename)
	return imageUrl, err
}

func uploadPictureToS3(im *utils.Image, s3Session *session.Session, bucket string, s3Filename string) error {
	uploader := s3manager.NewUploader(s3Session)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(s3Filename),
		Body:   bytes.NewReader(im.OutBytes),
	})
	if err != nil {
		return fmt.Errorf("upload %s err: %s", im.Path, err.Error())
	}
	return nil
}

func NewS3Storage() *S3Storage {
	return &S3Storage{utils.NewMetaStorage()}
}

func init() {
	utils.StroageHelp[s3PluginName] = s3Help()
	s3Config = conf.Viper.GetStringMap(s3PluginName)
	if s3Config == nil {
		return
	}
	active := s3Config["active"]
	if active == nil {
		return
	}
	if active == true {
		utils.StroageMap[s3PluginName] = NewS3Storage()
	}
}

func s3Help() string {
	return "s3 plugin need this parameters:\nactive: false or true\naccess_key: access key id string\nsecret_key: secret key string\nendpoint: s3 service entrypoint\ndisable_ssl: bool, if false, use ssl to upload\nforce_path_style: bool, path style\nprefix: string, file prefix which upload to s3\nbucket: string, the name of bucket which file upload to it\nregion: the bucket area, if not aws. use default value[us-en-1]."
}
