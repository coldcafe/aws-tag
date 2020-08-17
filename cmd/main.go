package main

import (
	"aws-tag/identity"
	"aws-tag/types"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	config, err := initConfig()
	if err != nil {
		panic(err)
	}
	identity, err := identity.GetInstanceIdentity()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", identity)
	sess, err := newSession(config.AWSAccessKeyID, config.AWSAccessKeySecret, identity.RegionID)
	if err != nil {
		panic(err)
	}
	service := ec2.New(sess)

	response, err := service.DescribeTags(&ec2.DescribeTagsInput{
		MaxResults: aws.Int64(100),
	})

	if err != nil {
		return
	}

	fmt.Printf("%+v", response.Tags)
}

func newSession(id, secret, region string) (*session.Session, error) {
	return session.NewSession(&aws.Config{
		Region:           aws.String(region),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials(id, secret, ""),
		DisableSSL:       aws.Bool(false),
	})
}

// 初始化配置项
func initConfig() (config *types.Configs, err error) {
	configBytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(configBytes, &config)
	return
}
