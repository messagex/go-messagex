package api

// Email - Object that represents the email payload to the MessageX API
type Email struct {
	From        *Contact      `json:"from"`                  // From contact. See the api.Contact struct for details.
	ReplyTo     *Contact      `json:"replyTo,omitempty"`     // ReplyTo contact. See the api.Contact struct for details.
	To          []*Contact    `json:"to"`                    // To contact. See the api.Contact struct for details.
	Cc          []*Contact    `json:"cc,omitempty"`          // List of CC contacts. See the api.Contact struct for details.
	Bcc         []*Contact    `json:"bcc,omitempty"`         // List of BCC contacts. See the api.Contact struct for details.
	Subject     string        `json:"subject"`               // Email subject.
	Content     []*Content    `json:"content"`               // Stores the html and plain email bodies.
	Headers     []*Header     `json:"headers,omitempty"`     // List of headers. See the api.Header struct for details.
	Attachments []*Attachment `json:"attachments,omitempty"` // List of attachments, see the api.Attachment struct for details.
}

// Headers - Email headers to be inserted into each email
type Header struct {
	Name  string `json:"name"`  // Name of the field
	Value string `json:"value"` // Value of the field
}

// Attachment - Object that holds information about an attachment
type Attachment struct {
	Contents       string `json:"contentEncoded"` // Base64 encoded string
	MimeType       string `json:"mimeType"`       // Mime type of the file, e.g.: "image/png"
	Filename       string `json:"filename"`       // The name of the file as it appears in the email client
}

// Contact - Holds contact information for From, Reply-To, To, Cc and Bcc
type Contact struct {
	Address string `json:"address"` // Contact email address
	Name    string `json:"name"`    // Contact name
}

// Content - Holds email body
type Content struct {
	Type string `json:"type"` // Type of body, i.e.: text/plain, text/html
	Body string `json:"body"` // Body content
}

// MailSendResponse - Represents the response from the MessageX API server for the mail/send request
type MailSendResponse struct {
	Success bool `json:"success"` // Indicates whether the API request was successful or not
}

// SetFrom - Sets the From contact
//
// Parameters:
//
//  name    - string representing the contact name
//  address - string representing the contact email address
func (e *Email) SetFrom(name, address string) {
	e.From = &Contact{
		Name:    name,
		Address: address,
	}
}

// SetReplyTo - Sets the ReplyTo contact
//
// Parameters:
//
//  name    - string representing the contact name
//  address - string representing the contact email address
func (e *Email) SetReplyTo(name, address string) {
	e.ReplyTo = &Contact{
		Name:    name,
		Address: address,
	}
}

// AddTo - Adds a contact to the To list
//
// Parameters:
//
//  name    - string representing the contact name
//  address - string representing the contact email address
func (e *Email) AddTo(name, address string) {
	e.To = append(e.To, &Contact{
		Name:    name,
		Address: address,
	})
}

// AddCc - Adds a contact to the CC list
//
// Parameters:
//
//  name    - string representing the contact name
//  address - string representing the contact email address
func (e *Email) AddCc(name, address string) {
	e.Cc = append(e.Cc, &Contact{
		Name:    name,
		Address: address,
	})
}

// AddBcc - Adds a contact to the BCC list
//
// Parameters:
//
//  name    - string representing the contact name
//  address - string representing the contact email address
func (e *Email) AddBcc(name, address string) {
	e.Bcc = append(e.Bcc, &Contact{
		Name:    name,
		Address: address,
	})
}

// SetSubject - Sets the email subject
//
// Parameters:
//
//  subject - representing the email subject
func (e *Email) SetSubject(subject string) {
	e.Subject = subject
}

// SetHTMLBody - Sets the email HTML body
//
// Parameters:
//
//  html - string representing the email HTML body
func (e *Email) SetHTMLBody(html string) {
	e.Content = append(e.Content, &Content{
		Type: "text/html",
		Body: html,
	})
}

// SetHTMLBody - Sets the email plain text body
//
// Parameters:
//
//  text - string representing the email plain text body
func (e *Email) SetPlainBody(text string) {
	e.Content = append(e.Content, &Content{
		Type: "text/plain",
		Body: text,
	})
}

// AddHeader - Adds a key/value pair to the list of email headers
//
// Parameters:
//
//  key   - string representing the key of the header
//  value - string representing the value of the header
func (e *Email) AddHeader(key, value string) {
	e.Headers = append(e.Headers, &Header{
		Name:  key,
		Value: value,
	})
}

// AddAttachment - Adds an attachment to the email
//
// Parameters:
//
//  fn - string representing the filename of the attachment
//  mt - string representing the mime type of the attachment
//  ct - string representing the base64 encoded content of the attachment
func (e *Email) AddAttachment(fn, mt, ct string) {
	e.Attachments = append(e.Attachments, &Attachment{
		Filename:       fn,
		MimeType:       mt,
		Contents:       ct,
	})
}
