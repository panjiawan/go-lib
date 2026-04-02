package psms

import (
	"encoding/json"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v5/client"
	"github.com/alibabacloud-go/tea/tea"
	"strings"
)

type AliYunProvider struct {
	config  *Config
	client  *dysmsapi.Client
	request *dysmsapi.SendSmsRequest
}

func (a *AliYunProvider) Init(config *Config) {
	a.config = config
	a.client, _ = dysmsapi.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(config.SecretId),
		AccessKeySecret: tea.String(config.SecretKey),
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
	})
}

func (a *AliYunProvider) Send(phone string, message []string) (string, error) {
	req := &dysmsapi.SendSmsRequest{
		PhoneNumbers: tea.String(phone),
		SignName:     tea.String(a.config.Sign),
		TemplateCode: tea.String(a.config.TemplateId),
	}
	par, err := json.Marshal(map[string]interface{}{ //定义短信模板参数（具体需要几个参数根据自己短信模板格式）
		"code": message[0],
	})
	if err != nil {
		return "", err
	}
	pm := string(par)
	req.TemplateParam = &pm

	rep, err := a.client.SendSms(req)

	res, _ := json.Marshal(rep)

	return string(res), err
}

func (a *AliYunProvider) SendMultiple(phones []string, message []string) (string, error) {
	req := &dysmsapi.SendSmsRequest{
		PhoneNumbers: tea.String(strings.Join(phones, ",")),
		SignName:     tea.String(a.config.Sign),
		TemplateCode: tea.String(a.config.TemplateId),
	}
	par, err := json.Marshal(map[string]interface{}{ //定义短信模板参数（具体需要几个参数根据自己短信模板格式）
		"code": message[0],
	})
	if err != nil {
		return "", err
	}
	pm := string(par)
	req.TemplateParam = &pm

	rep, err := a.client.SendSms(req)

	res, _ := json.Marshal(rep)

	return string(res), err
}
