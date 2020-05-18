
# MessageX SDK - Go

---
[![GoDoc](https://godoc.org/github.com/messagex/go-messagex?status.svg)](https://godoc.org/github.com/messagex/go-messagex)
![Go](https://github.com/messagex/go-messagex/workflows/Go/badge.svg?branch=master)
[![Sourcegraph](https://sourcegraph.com/github.com/messagex/go-messagex/-/badge.svg)](https://sourcegraph.com/github.com/messagex/go-messagex?badge)
![Go](https://github.com/messagex/go-messagex/workflows/Go/badge.svg?branch=master&event=status)

![MessageX Logo](https://raw.githubusercontent.com/messagex/node-messagex/master/img/messagex-logo.png "MessageX")

Company website <https://www.smsglobal.com>.

MessageX SDK provides a library that provides an interface into the MessageX functionality.

---

* [Install Prerequisites](#install-prerequisites)
* [Install MeesageX API package](#install)
* [Examples](#examples)
* [Developers](#developers)
  * [Installing GoDoc for Go v1.1.4](#installing-godoc-for-go-v1.1.4)
  * [Running Documentation Server](#running-documentation-server)
  * [Running Tests](#running-tests)

---

## Install Prerequisites

The following are the prerequisites for this package

### Go v1.14

Use following page to install go <https://golang.org/doc/install?download=go1.14.2.linux-amd64.tar.gz>

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/messagex/go-messagex
```

---

## Examples

### Sending email

   Import the library into your package

```go
import (
    "github.com/messagex/go-messagex"
)
```

   The following example shows how to send an email

 ```go
package main

import (
  "bufio"
  "encoding/base64"
  "fmt"
  "io/ioutil"
  "os"

  "github.com/messagex/go-messagex"
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
  f, _ := os.Open("./image.png")

  // Read entire JPG into byte slice.
  reader := bufio.NewReader(f)
  content, _ := ioutil.ReadAll(reader)

  // Encode as base64.
  attPic := base64.StdEncoding.EncodeToString(content)

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
```

---

## Developers

### Installing GoDoc for Go v1.1.4

Once Go v1.14 has been installed, execute the following commands to install
GoDoc from go-messagex project's root directory.

```shell script
GOBIN=`pwd`/bin && mkdir bin && go get -u golang.org/x/tools/cmd/godoc
```

Make sure you use the godoc tool that is installed in go-messagex/bin directory

### Running documentation server

./dev-bin/docs

Navigate here to see the documentation: <http://localhost:6060/pkg/github.com/messagex/go-messagex>

Note: Until the release to github, the address is <http://localhost:6060/pkg/github.com/messagex/go-messagex/#APIClient>

Note: We need to find out how to publish the documentation on godoc.org

### Running Tests

To run unit tests execute the following command

```shell script
./dev-bin/test
```
