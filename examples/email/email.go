package main

import (
	"fmt"
	"os"

	"vines.smsglobal.local/messagex/v2/sdk/go-messagex"
)

var apiKey = "qsP8QCUF9q4Lmmv6dKDJukOIZ9vA1BVIRiy91RGjKaU42ovWtdYb2gwYaqhYY788"
var apiSecret = "ezEONICwdyMyPkLdU553wXz5xk0OoV4y6HnngRcAimj1hoid7ZLl61k9pV1bMDoY"

func main() {

	apiClient := messagex.CreateAPIClient(apiKey, apiSecret)

	err := apiClient.Login()
	if err != nil {
		fmt.Printf("Error logging in to MessageX API: %s", err.Error())
		os.Exit(0)
	}

	htmlBody :=  `
<!DOCTYPE HTML>
<html>
    <head>
        <meta charset="UTF-8">
        <title>Test Email Body</title>
    </head>
    <body>
        <h1>Test Email Body Heading</h1>
        <p>
            Test email body paragraph
        </p>
    </body>
</html>
`
	attPic := "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"

	email := apiClient.CreateEmail()
	email.SetFrom("Sender Name", "from@example.com")
	email.SetReplyTo("ReplyTo Name", "reply@example.com")
	email.AddTo("To Name", "to@example.com")
	email.AddCc("CC Name", "cc@example.com")
	email.AddBcc("BCC Name", "bcc@example.com")
	email.AddHeader("Content-Type", "application/json")
	email.SetSubject("Test email subject")
	email.SetHTMLBody(htmlBody)
	email.SetPlainBody("Example plain email body")
	email.AddAttachment("filename.jpg", "image/png", attPic)

	err = apiClient.SendEmail(email)
	if err != nil {
		fmt.Printf("Error while sending an email to MessageX: %s", err.Error())
		os.Exit(0)
	}
}