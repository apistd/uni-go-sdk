package unisms

import (
	"errors"

	"github.com/apistd/uni-go-sdk"
)

type UniSMSClient struct {
	Client	*uni.UniClient
}

type UniMessage struct {
	To				*[] string
	Signature		*string
	TemplateId		*string
	TemplateData	*map[string]string
	Content			*string
	Text			*string
}

func NewClient(params ...string) *UniSMSClient {
	return &UniSMSClient{
		Client: uni.NewClient(params...),
	}
}

func BuildMessage() *UniMessage {
	return &UniMessage{}
}

func (m *UniMessage) SetTo(phoneNumbers ...string) *UniMessage {
	m.To = &phoneNumbers
	return m
}

func (m *UniMessage) SetSignature(signature string) *UniMessage {
	m.Signature = &signature
	return m
}

func (m *UniMessage) SetTemplateId(templateId string) *UniMessage {
	m.TemplateId = &templateId
	return m
}

func (m *UniMessage) SetTemplateData(templateData map[string]string) *UniMessage {
	m.TemplateData = &templateData
	return m
}

func (m *UniMessage) SetContent(content string) *UniMessage {
	m.Content = &content
	return m
}

func (m *UniMessage) SetText(text string) *UniMessage {
	m.Text = &text
	return m
}

func (c *UniSMSClient) Send(m *UniMessage) (response *uni.UniResponse, err error) {
	data := make(map[string]interface{})

	if (m.To == nil) {
		return nil, errors.New("To phone number is required")
	} else {
		data["to"] = m.To
	}

	if (m.Signature != nil) {
		data["signature"] = m.Signature
	}

	if (m.TemplateId != nil) {
		data["templateId"] = m.TemplateId
	}

	if (m.TemplateData != nil) {
		data["templateData"] = m.TemplateData
	}

	if (m.TemplateData != nil) {
		data["templateData"] = m.TemplateData
	}

	if (m.Content != nil) {
		data["content"] = m.Content
	}

	if (m.Text != nil) {
		data["text"] = m.Text
	}

	return c.Client.Request("sms.message.send", data)
}
