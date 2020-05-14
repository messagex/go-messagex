package api

import "github.com/messagex/go-messagex/internal/types/constants"

// MailJob - this is the format in which the API layer passes the campaign mail job to the Core
type Email struct {
	From        *Contact      `json:"from"`
	ReplyTo     *Contact      `json:"replyTo,omitempty"`
	To          []*Contact    `json:"to"`
	Cc          []*Contact    `json:"cc,omitempty"`
	Bcc         []*Contact    `json:"bcc,omitempty"`
	Subject     string        `json:"subject"`
	Content     []*Content    `json:"content"`
	Headers     []*Header     `json:"headers,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

// Headers - email headers to be inserted into each email
type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Attachment - an object that holds information about an attachment
type Attachment struct {
	Contents       string `json:"contentEncoded"` // Base64 encoded string
	MimeType       string `json:"mimeType"`
	Filename       string `json:"filename"`
	AttachmentType string `json:"attachmentType"`
}

// Contact - holds contact information for From, Reply-To, To, Cc and Bcc
type Contact struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

// Content - holds email body
type Content struct {
	Type    string `json:"type"`
	Body    string `json:"body"`
}

type MailSendResponse struct {
	Success bool `json:"success"`
}

func (e *Email) SetFrom(name, address string) {
	e.From = &Contact{
		Name:    name,
		Address: address,
	}
}

func (e *Email) SetReplyTo(name, address string) {
	e.ReplyTo = &Contact{
		Name:    name,
		Address: address,
	}
}

func (e *Email) AddTo(name, address string) {
	e.To = append(e.To, &Contact{
		Name:    name,
		Address: address,
	})
}

func (e *Email) AddCc(name, address string) {
	e.Cc = append(e.Cc, &Contact{
		Name:    name,
		Address: address,
	})
}

func (e *Email) AddBcc(name, address string) {
	e.Bcc = append(e.Bcc, &Contact{
		Name:    name,
		Address: address,
	})
}

func (e *Email) SetSubject(subject string) {
	e.Subject = subject
}

func (e *Email) SetHTMLBody(html string) {
	e.Content = append(e.Content, &Content{
		Type:    "text/html",
		Body:    html,
	})
}

func (e *Email) SetPlainBody(text string) {
	e.Content = append(e.Content, &Content{
		Type:    "text/plain",
		Body:    text,
	})
}

func (e *Email) AddHeader(key, value string) {
	e.Headers = append(e.Headers, &Header{
		Name:  key,
		Value: value,
	})
}

func (e *Email) AddAttachment(fn, mt, ct string) {
	e.Attachments = append(e.Attachments, &Attachment{
		Filename:       fn,
		MimeType:       mt,
		Contents:       ct,
		AttachmentType: constants.AttachmentType,
	})
}
