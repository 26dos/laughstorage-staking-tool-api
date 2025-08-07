package service

import (
	"api/internal/consts"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	cloudauth20190307 "github.com/alibabacloud-go/cloudauth-20190307/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	credential "github.com/aliyun/credentials-go/credentials"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

var Ali = new(ali)

type ali struct {
	Error         error
	CredentialCli *cloudauth20190307.Client
}

func (a *ali) InitCredential(ctx context.Context) *ali {
	credentialConfig := new(credential.Config)
	accessKeyId, _ := g.Cfg().Get(ctx, "ali.accessKeyId")
	accessKeySecret, _ := g.Cfg().Get(ctx, "ali.accessKeySecret")
	credentialConfig.SetType("access_key").SetAccessKeyId(accessKeyId.String()).SetAccessKeySecret(accessKeySecret.String())

	credential, err := credential.NewCredential(credentialConfig)
	if err != nil {
		g.Log().Error(ctx, "aliyun credential error", err)
		Ali.Error = CustomError.ServerError(ctx, "Kyc verify failed,please try again")
		return a
	}
	config := &openapi.Config{
		Credential: credential,
	}
	config.Endpoint = tea.String("cloudauth.aliyuncs.com")
	client, err := cloudauth20190307.NewClient(config)
	if err != nil {
		g.Log().Error(ctx, "aliyun cloudauth client error", err)
		Ali.Error = CustomError.ServerError(ctx, "Kyc verify failed,please try again")
		return a
	}
	Ali.CredentialCli = client
	return a
}

func (a *ali) Id2MetaVerify(ctx context.Context, idCardNo string, name string) error {
	if Ali.Error != nil {
		return Ali.Error
	}
	id2MetaVerifyRequest := &cloudauth20190307.Id2MetaVerifyRequest{
		IdentifyNum: tea.String(idCardNo),
		UserName:    tea.String(name),
		ParamType:   tea.String("normal"),
	}
	runtime := &util.RuntimeOptions{}
	res, err := Ali.CredentialCli.Id2MetaVerifyWithOptions(id2MetaVerifyRequest, runtime)
	if err != nil {
		g.Log().Error(ctx, "aliyun id2meta verify error：", err)
		return err
	}
	if *res.Body.Code != "200" {
		return CustomError.ParameterError(ctx, "kyc verify failed,please try again")
	}
	g.Log().Info(ctx, "aliyun id2meta verify success", res)
	return nil
}

func (a *ali) CredentialVerify(ctx context.Context, idCardNo string, name string, imageContext string) error {
	if Ali.Error != nil {
		return Ali.Error
	}
	ossUrl, fileName, err := a.UploadBase64ImageOss(ctx, imageContext)
	if err != nil {
		return err
	}
	credentialVerifyRequest := &cloudauth20190307.CredentialVerifyV2Request{
		IdentifyNum: tea.String(idCardNo),
		CredName:    tea.String("0101"),
		CredType:    tea.String("01"),
		IsCheck:     tea.String("1"),
		UserName:    tea.String(name),
		ImageUrl:    tea.String(ossUrl),
	}
	runtime := &util.RuntimeOptions{}
	res, err := Ali.CredentialCli.CredentialVerifyV2WithOptions(credentialVerifyRequest, runtime)
	if err != nil {
		g.Log().Error(ctx, "aliyun credential verify error：", err)
		a.DeleteOssFile(ctx, fileName)
		return consts.ServerErr
	}
	if *res.Body.Code != "200" {
		a.DeleteOssFile(ctx, fileName)
		g.Log().Error(ctx, "aliyun credential verify error：", res)
		return CustomError.ParameterError(ctx, "kyc verify failed,please try again,error!")
	}
	a.DeleteOssFile(ctx, fileName)
	return nil
}

func (a *ali) Id2MetaVerifyWithOCR(ctx context.Context, Front string, back string) error {
	if Ali.Error != nil {
		return Ali.Error
	}
	frontOssUrl, frontFileName, err := a.UploadBase64ImageOss(ctx, Front)
	if err != nil {
		return err
	}
	backOssUrl, backFileName, err := a.UploadBase64ImageOss(ctx, back)
	if err != nil {
		return err
	}
	id2MetaVerifyWithOCRRequest := &cloudauth20190307.Id2MetaVerifyWithOCRRequest{
		CertUrl:         tea.String(frontOssUrl),
		CertNationalUrl: tea.String(backOssUrl),
	}
	runtime := &util.RuntimeOptions{}
	res, err := Ali.CredentialCli.Id2MetaVerifyWithOCRWithOptions(id2MetaVerifyWithOCRRequest, runtime)
	if err != nil {
		g.Log().Error(ctx, "aliyun id2meta verify with ocr error：", err)
		a.DeleteOssFile(ctx, frontFileName)
		a.DeleteOssFile(ctx, backFileName)
		return err
	}
	if *res.Body.Code != "200" {
		a.DeleteOssFile(ctx, frontFileName)
		a.DeleteOssFile(ctx, backFileName)
		g.Log().Error(ctx, "aliyun id2meta verify with ocr error：", res)
		return CustomError.ParameterError(ctx, "kyc verify failed,please try again,error!")
	}
	a.DeleteOssFile(ctx, frontFileName)
	a.DeleteOssFile(ctx, backFileName)
	return nil

}

func (a *ali) DeleteOssFile(ctx context.Context, fileName string) {
	region, _ := g.Cfg().Get(ctx, "ali.region")
	bucketName, _ := g.Cfg().Get(ctx, "ali.ossBucketName")
	accessKeyId, _ := g.Cfg().Get(ctx, "ali.accessKeyId")
	accessKeySecret, _ := g.Cfg().Get(ctx, "ali.accessKeySecret")
	provider := credentials.NewStaticCredentialsProvider(accessKeyId.String(), accessKeySecret.String())
	cfg := oss.LoadDefaultConfig().WithCredentialsProvider(provider).WithRegion(region.String())
	ossClient := oss.NewClient(cfg)
	request := &oss.DeleteObjectRequest{
		Bucket: tea.String(bucketName.String()),
		Key:    tea.String(fileName),
	}
	res, err := ossClient.DeleteObject(ctx, request)
	if err != nil {
		g.Log().Error(ctx, "aliyun oss delete error：", err)
	}
	g.Log().Info(ctx, "aliyun oss delete success", res)
}

func (a *ali) UploadBase64ImageOss(ctx context.Context, base64Image string) (string, string, error) {
	region, _ := g.Cfg().Get(ctx, "ali.region")
	bucketName, _ := g.Cfg().Get(ctx, "ali.ossBucketName")
	ossUrl, _ := g.Cfg().Get(ctx, "ali.ossUrl")
	accessKeyId, _ := g.Cfg().Get(ctx, "ali.accessKeyId")
	accessKeySecret, _ := g.Cfg().Get(ctx, "ali.accessKeySecret")
	provider := credentials.NewStaticCredentialsProvider(accessKeyId.String(), accessKeySecret.String())
	cfg := oss.LoadDefaultConfig().WithCredentialsProvider(provider).WithRegion(region.String())
	ossClient := oss.NewClient(cfg)

	// 处理 base64 字符串
	base64Image = strings.TrimSpace(base64Image)
	if strings.Contains(base64Image, "base64,") {
		base64Image = strings.Split(base64Image, "base64,")[1]
	}
	// 替换可能的特殊字符
	base64Image = strings.ReplaceAll(base64Image, "\n", "")
	base64Image = strings.ReplaceAll(base64Image, "\r", "")
	base64Image = strings.ReplaceAll(base64Image, " ", "")

	// 解码 base64 字符串
	imageBytes, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		g.Log().Error(ctx, "base64 decode error：", err, "base64 string:", base64Image[:min(100, len(base64Image))])
		return "", "", consts.ServerErr
	}
	fileName := fmt.Sprintf("%s.jpg", uuid.New().String())
	result, err := ossClient.PutObject(ctx, &oss.PutObjectRequest{
		Bucket: oss.Ptr(bucketName.String()),
		Key:    oss.Ptr(fileName),
		Body:   bytes.NewReader(imageBytes),
	})
	if err != nil || result == nil {
		g.Log().Error(ctx, "aliyun oss upload error：", err)
		return "", "", consts.ServerErr
	}
	if result.StatusCode != 200 {
		g.Log().Error(ctx, "aliyun oss upload error：", result.StatusCode)
		return "", "", consts.ServerErr
	}
	return fmt.Sprintf("%s/%s", ossUrl.String(), fileName), fileName, nil
}
