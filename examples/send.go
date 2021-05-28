package main

import (
	"fmt"

	unisms "github.com/apistd/uni-go-sdk/sms"
)

func main() {
	client := unisms.NewClient("your access key id", "your access key secret")

	message := unisms.BuildMessage()
	message.SetTo("your phone number")
	message.SetSignature("UniSMS")
	message.SetTemplateId("login_tmpl")
	message.SetTemplateData(map[string]string {"code": "6666"})

	res, err := client.Send(message)
	if (err != nil) {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
