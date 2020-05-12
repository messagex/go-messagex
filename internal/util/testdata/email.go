package testdata

import (
	"vines.smsglobal.local/messagex/v2/sdk/go-messagex/internal/types/api"
)

func GetTestMJR() *api.Email {
	return &api.Email{
		From: &api.Contact{
			Address: "from@example.com",
			Name:    "From Name",
		},
		ReplyTo: &api.Contact{
			Address: "reply@example.com",
			Name:    "Reply Name",
		},
		To: []*api.Contact{
			{
				Address: "to1@example.com",
				Name:    "To Name",
			},
		},
		Cc: []*api.Contact{
			{
				Address: "cc@example.com",
				Name:    "CC Name",
			},
		},
		Bcc: []*api.Contact{
			{
				Address: "bcc@example.com",
				Name:    "BCC Name",
			},
		},
		Subject: "Test Subject",
		Content: []*api.Content{
			{
				Type: "text/plain",
				Body: "My Plain Body",
			},
			{
				Type: "text/html",
				Body: "\n<!DOCTYPE HTML>\n<html>\n    <head>\n        <meta charset=\"UTF-8\">\n        <title>Dude</title>\n    </head>\n    <body>\n        <h1>Shall we go for lunch</h1>\n        <p>\n            Are you doing anything around 1pm?\n        </p>\n    </body>\n</html>\n",
			},
		},
		Headers: []*api.Header{
			{
				Name: "h1",
				Value: "V1",
			},
			{
				Name: "Content-Type",
				Value: "multipart/mixed; boundary=27a9fc1606043bf43fbb12b19e19d47bce0c983229063c7308a5f793ed8c",
			},
			{
				Name: "To",
				Value: "\"To Name\" <to@example.com>",
			},
			{
				Name: "From",
				Value: "\"From Name\" <from@example.com>",
			},
			{
				Name: "Reply-To",
				Value: "\"Reply Name\" <reply@example.com>",
			},
			{
				Name: "Cc",
				Value: "\"CC Name\" <cc@example.com>",
			},
			{
				Name: "Subject",
				Value: "test subject",
			},
			{
				Name: "Mime-Version",
				Value: "1.0",
			},
		},
		Attachments: []*api.Attachment{
			{
				MimeType: "text/plain",
				Filename: "acct-2019-12-28-0000.csv",
				AttachmentType: "attachment",
				Contents: "dHlwZSx0aW1lTG9nZ2VkLHRpbWVRdWV1ZWQsb3JpZyxyY3B0LG9yY3B0LGRzbkFjdGlvbixkc25TdGF0dXMsZHNuRGlhZyxkc25NdGEsYm91bmNlQ2F0LHNyY1R5cGUsc3JjTXRhLGRsdlR5cGUsZGx2U291cmNlSXAsZGx2RGVzdGluYXRpb25JcCxkbHZTaXplLHZtdGEsam9iSWQsZW52SWQscXVldWUsdm10YVBvb2wsaGVhZGVyX1gtR0lELGhlYWRlcl9Gcm9tDQpkLDIwMTktMTItMjggMDM6Mjk6NTMrMTEwMCwyMDE5LTEyLTI4IDAzOjI5OjQ5KzExMDAsbm9yZXBseUBzbXNnbG9iYWwuY29tLHNtc2dsb2JhbDAxQGdtYWlsLmNvbSwscmVsYXllZCwyLjAuMCAoc3VjY2Vzcyksc210cDsyNTAgMi4wLjAgT0sgMTU3NzQ2NDE5MyB6MTRzaTE4NTgwNzk0b3RoLjE1IC0gZ3NtdHAsZ21haWwtc210cC1pbi5sLmdvb2dsZS5jb20gKDc0LjEyNS4yNC4yNyksLHNtdHAsbG9jYWxob3N0ICgxOTIuMTY4LjE3LjExMiksc210cCwxOTIuMTY4LjUxLjIzLDc0LjEyNS4yNC4yNywxNDExLG1haWwyMy1nMSwsLGdtYWlsLmNvbS9tYWlsMjMtZzEsbXhkZWxpdmVybmV0LW9uYm90cmFuczItMTQ0NDgxNTItMjMtMjQsMTRlM2ZlYjQtYzUzOC00NGU2LWIzY2YtMmE2OTg2YjVkY2I2LCIiIiIiIDxub3JlcGx5QHNtc2dsb2JhbC5jb20+Ig0KZCwyMDE5LTEyLTI4IDAzOjM4OjU0KzExMDAsMjAxOS0xMi0yOCAwMzozODo1MSsxMTAwLG5vcmVwbHlAc21zZ2xvYmFsLmNvbSxzbXNnbG9iYWwwMUBnbWFpbC5jb20sLHJlbGF5ZWQsMi4wLjAgKHN1Y2Nlc3MpLHNtdHA7MjUwIDIuMC4wIE9LIDE1Nzc0NjQ3MzQgOTJzaTExNjQzN3BsZS40MSAtIGdzbXRwLGdtYWlsLXNtdHAtaW4ubC5nb29nbGUuY29tICg3NC4xMjUuMjQuMjcpLCxzbXRwLGxvY2FsaG9zdCAoMTkyLjE2OC4xNy4xMTEpLHNtdHAsMTkyLjE2OC41MS4yNCw3NC4xMjUuMjQuMjcsMTQ0MyxtYWlsMjQtZzEsLCxnbWFpbC5jb20vbWFpbDI0LWcxLG14ZGVsaXZlcm5ldC1vbmJvdHJhbnMyLTE0NDQ4MTUyLTIzLTI0LDczMDhmNGYxLTUyMmItNDc5Ni05MTY0LTAyYTczZGFhZmFkNSwiIiIiIiA8bm9yZXBseUBzbXNnbG9iYWwuY29tPiINCmQsMjAxOS0xMi0yOCAwNDozNzo1MSsxMTAwLDIwMTktMTItMjggMDQ6Mzc6NDgrMTEwMCxub3JlcGx5QHNtc2dsb2JhbC5jb20sc21zZ2xvYmFsMDFAZ21haWwuY29tLCxyZWxheWVkLDIuMC4wIChzdWNjZXNzKSxzbXRwOzI1MCAyLjAuMCBPSyAxNTc3NDY4MjcxIGYxOTFzaTMyNzA0MzcycGZhLjAgLSBnc210cCxnbWFpbC1zbXRwLWluLmwuZ29vZ2xlLmNvbSAoNzQuMTI1LjI0LjI2KSwsc210cCxsb2NhbGhvc3QgKDE5Mi4xNjguMTcuMTExKSxzbXRwLDE5Mi4xNjguNTEuMjMsNzQuMTI1LjI0LjI2LDE0ODgsbWFpbDIzLWcxLCwsZ21haWwuY29tL21haWwyMy1nMSxteGRlbGl2ZXJuZXQtb25ib3RyYW5zMi0xNDQ0ODE1Mi0yMy0yNCwwYTY0MGQ2Zi1hM2FlLTRkMzYtODkzYi0zNTY0NTJiN2E5MzksIiIiIiIgPG5vcmVwbHlAc21zZ2xvYmFsLmNvbT4iDQpkLDIwMTktMTItMjggMDU6Mjg6NDIrMTEwMCwyMDE5LTEyLTI4IDA1OjI4OjM5KzExMDAsbm9yZXBseUBzbXNnbG9iYWwuY29tLHNtc2dsb2JhbDAxQGdtYWlsLmNvbSwscmVsYXllZCwyLjAuMCAoc3VjY2Vzcyksc210cDsyNTAgMi4wLjAgT0sgMTU3NzQ3MTMyMiBqN3NpMzEzNzI1MDlwZ2ouNTY2IC0gZ3NtdHAsZ21haWwtc210cC1pbi5sLmdvb2dsZS5jb20gKDc0LjEyNS4yNC4yNyksLHNtdHAsbG9jYWxob3N0ICgxOTIuMTY4LjE3LjExNCksc210cCwxOTIuMTY4LjUxLjI0LDc0LjEyNS4yNC4yNywxMzg5LG1haWwyNC1nMSwsLGdtYWlsLmNvbS9tYWlsMjQtZzEsbXhkZWxpdmVybmV0LW9uYm90cmFuczItMTQ0NDgxNTItMjMtMjQsNjQyYTNlMDMtYjIwYi00MzcxLWFjYTYtNTRhOTA1NTZlYjk2LCIiIiIiIDxub3JlcGx5QHNtc2dsb2JhbC5jb20+Ig0KZCwyMDE5LTEyLTI4IDA1OjM4OjUzKzExMDAsMjAxOS0xMi0yOCAwNTozODo1MCsxMTAwLG5vcmVwbHlAc21zZ2xvYmFsLmNvbSxzbXNnbG9iYWwwMUBnbWFpbC5jb20sLHJlbGF5ZWQsMi4wLjAgKHN1Y2Nlc3MpLHNtdHA7MjUwIDIuMC4wIE9LIDE1Nzc0NzE5MzMgdzdzaTMxMDg0MDIwcGdtLjIyOCAtIGdzbXRwLGdtYWlsLXNtdHAtaW4ubC5nb29nbGUuY29tICg3NC4xMjUuMjQuMjcpLCxzbXRwLGxvY2FsaG9zdCAoMTkyLjE2OC4xNy4xMTEpLHNtdHAsMTkyLjE2OC41MS4yMyw3NC4xMjUuMjQuMjcsMTQyMSxtYWlsMjMtZzEsLCxnbWFpbC5jb20vbWFpbDIzLWcxLG14ZGVsaXZlcm5ldC1vbmJvdHJhbnMyLTE0NDQ4MTUyLTIzLTI0LGNmZDZmMjQ4LTY1ZTEtNGNhNS1hYmZkLTViNzdkYzk1YWQwOCwiIiIiIiA8bm9yZXBseUBzbXNnbG9iYWwuY29tPiINCmQsMjAxOS0xMi0yOCAwNTo1OTowOSsxMTAwLDIwMTktMTItMjggMDU6NTk6MDYrMTEwMCxub3JlcGx5QHNtc2dsb2JhbC5jb20sc21zZ2xvYmFsMDFAZ21haWwuY29tLCxyZWxheWVkLDIuMC4wIChzdWNjZXNzKSxzbXRwOzI1MCAyLjAuMCBPSyAxNTc3NDczMTQ5IDg0c2kzMDU3ODc4OXBnZy4yNDUgLSBnc210cCxnbWFpbC1zbXRwLWluLmwuZ29vZ2xlLmNvbSAoNzQuMTI1LjI0LjI3KSwsc210cCxsb2NhbGhvc3QgKDE5Mi4xNjguMTcuMTExKSxzbXRwLDE5Mi4xNjguNTEuMjQsNzQuMTI1LjI0LjI3LDEzNjgsbWFpbDI0LWcxLCwsZ21haWwuY29tL21haWwyNC1nMSxteGRlbGl2ZXJuZXQtb25ib3RyYW5zMi0xNDQ0ODE1Mi0yMy0yNCw3ZmFmYWNhYS0xYzlkLTQzYzAtYmNlNC05NjgyOTg1YWRmOWYsIiIiIiIgPG5vcmVwbHlAc21zZ2xvYmFsLmNvbT4iDQo=",
			},
			{
				MimeType: "image/png",
				Filename: "hm.png",
				AttachmentType: "attachment",
				Contents: "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",
			},
		},
	}
}
