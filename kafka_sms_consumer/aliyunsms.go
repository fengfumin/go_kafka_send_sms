package kafka_sms_consumer

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)


type Sms struct {
	Phone   string `json:"phone"`
	Content string `json:"content"`
}

type Server struct {
	SMS *SMS
}

type SMS struct {
	KeyId  string
	KeySecret string
	SignName    string
	TemplateCode    string
}

type Resp struct {
	Message string
	RequestId string
	Code string
}

func NewSmsServer(keyId, keySecret,signName, templateCode string) *Server {
	return &Server{
		NewSMS(keyId, keySecret,signName,templateCode),
	}
}

func NewSMS(keyId, keySecret,signName,templateCode string) (m *SMS) {
	m = new(SMS)
	m.init(keyId, keySecret,signName,templateCode)
	return
}

func (m *SMS) init(keyId, keySecret,signName,templateCode string) {
	m.KeyId = keyId
	m.KeySecret = keySecret
	m.SignName = signName
	m.TemplateCode = templateCode
}

func (s *Server) SendSms(value []byte) (err error) {
	var sms Sms
	err = json.Unmarshal(value, &sms)
	if err != nil {
		return
	}
	return s.SMS.SendSMS(sms.Phone, sms.Content)

}


func (m *SMS) SendSMS(phone, content string) (err error) {
	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", m.KeyId, m.KeySecret)
	if err != nil {
		panic(err)
	}

	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-hangzhou"
	request.QueryParams["PhoneNumbers"] = phone
	request.QueryParams["SignName"] = m.SignName
	request.QueryParams["TemplateCode"] = m.TemplateCode
	request.QueryParams["TemplateParam"] = fmt.Sprintf(`{"code":%s}`,content)

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		fmt.Println(err)
		return
	}


	if status:=response.GetHttpStatus();status==200{
		resbytes:=response.GetHttpContentBytes()
		var resp Resp
		err=json.Unmarshal(resbytes, &resp)
		if err != nil {
			return
		}
		fmt.Println("code",":",resp.Code)
		fmt.Println("RequestId",":",resp.RequestId)
		fmt.Println("Message",":",resp.Message)
		if resp.Code !="OK"{
			fmt.Println("通知管理员，发送短信失败")
		}

	}else {
		fmt.Println(status)
	}

	return
}
