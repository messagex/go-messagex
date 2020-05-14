package api

import (
	"testing"
	"github.com/messagex/go-messagex/internal/types/constants"

	"github.com/stretchr/testify/assert"
)

func TestSetFrom(t *testing.T) {
	em := &Email{}
	em.SetFrom("first last", "test@example.com")
	assert.ObjectsAreEqual(em.From, &Contact{
		Name: "first last",
		Address: "test@example.com",
	})
}

func TestSetReplyTo(t *testing.T) {
	em := &Email{}
	em.SetReplyTo("first last", "test@example.com")
	assert.ObjectsAreEqual(em.ReplyTo, &Contact{
		Name: "first last",
		Address: "test@example.com",
	})
}

func TestAddTo(t *testing.T) {
	em := &Email{}
	em.AddTo("first last", "test@example.com")
	assert.ObjectsAreEqual(em.To, []*Contact{
		{
		Name: "first last",
		Address: "test@example.com",
	}})
}

func TestAddCc(t *testing.T) {
	em := &Email{}
	em.AddCc("first last", "test@example.com")
	assert.ObjectsAreEqual(em.Cc, []*Contact{
		{
			Name: "first last",
			Address: "test@example.com",
		}})
}

func TestAddBcc(t *testing.T) {
	em := &Email{}
	em.AddBcc("first last", "test@example.com")
	assert.ObjectsAreEqual(em.Bcc, []*Contact{
		{
			Name: "first last",
			Address: "test@example.com",
		}})
}

func TestSetSubject(t *testing.T) {
	em := &Email{}
	em.SetSubject("subject string")
	assert.Equal(t, "subject string", em.Subject, "Subject is as expected")
}

func TestAddHeader(t *testing.T) {
	em := &Email{}
	em.AddHeader("key", "value")
	assert.ObjectsAreEqual(em.Headers, []*Header{
		{
			Name: "key",
			Value: "value",
		}})
}

func TestSetHTMLBody(t *testing.T) {
	em := &Email{}
	em.SetHTMLBody("<html></html>")
	assert.Equal(t, em.Content, []*Content{
		{
			Type: "text/html",
			Body: "<html></html>",
		}})
}

func TestSetPlainBody(t *testing.T) {
	em := &Email{}
	em.SetPlainBody("plain text")
	assert.Equal(t, em.Content, []*Content{
		{
			Type: "text/plain",
			Body: "plain text",
		}})
}

func TestAddAttachment(t *testing.T) {
	em := &Email{}
	em.AddAttachment("filename.txt", "text/plain", "test content")
	assert.ObjectsAreEqual(em.Attachments, []*Attachment{
		{
			Filename: "filename.txt",
			MimeType: "text/plain",
			Contents: "test content",
			AttachmentType: constants.AttachmentType,
		}})
}