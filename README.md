# Uni Go SDK

[UniSMS](https://unisms.apistd.com/) - 高可用聚合短信服务平台官方 Go SDK.

## 文档

查看完整产品介绍与 API 文档请访问 [UniSMS Documentation](https://unisms.apistd.com/docs).

## 安装

Uni Go SDK 提供 Go Module，可从公共 [Github 仓库](https://github.com/apistd/uni-go-sdk) 中获得。

在项目中添加 `uni-go-sdk` 作为依赖：

```bash
go get github.com/apistd/uni-go-sdk
```

## 使用示例

以下示例展示如何使用 Uni Go SDK 快速调用服务。

### 发送短信

```go

package main

import (
    "fmt"
    unisms "github.com/apistd/uni-go-sdk/sms"
)

func main() {
    // 初始化
    client := unisms.NewClient("your access key id", "your access key secret")

    // 创建信息
    message := unisms.BuildMessage()
    message.SetTo("your phone number")
    message.SetSignature("UniSMS")
    message.SetTemplateId("login_tmpl")
    message.SetTemplateData(map[string]string {"code": "6666"}) // 设置自定义参数 (变量短信)

    // 发送短信
    res, err := client.Send(message)
    if (err != nil) {
        fmt.Println(err)
        return
    }
    fmt.Println(res)
}

```

## 相关参考

### 其他语言 SDK

- [Java](https://github.com/apistd/uni-java-sdk)
- [Node.js](https://github.com/apistd/unisms-node-sdk)
- [Python](https://github.com/apistd/uni-python-sdk)
- [PHP](https://github.com/apistd/uni-php-sdk/)
