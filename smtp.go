package helpers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"
)

type Mail struct {
	Host           string
	Port           int
	User           string
	Password       string
	From           string
	To             []string
	Cc             []string
	Bcc            []string
	Subject        string
	Msg            []byte
	AttachmentName string
	Attachment     []byte
}

func (m Mail) Builder() (body []byte) {
	buf := new(bytes.Buffer)

	// Mail Header
	buf.WriteString(fmt.Sprintf("From: %s\r\n", m.From))
	buf.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(m.To, ";")))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", m.Subject))

	if len(m.Cc) != 0 {
		buf.WriteString(fmt.Sprintf("Cc: %s\r\n", strings.Join(m.Cc, ";")))
	}

	if len(m.Bcc) != 0 {
		buf.WriteString(fmt.Sprintf("Bcc: %s\r\n", strings.Join(m.Bcc, ";")))
	}

	boundary := "zercle-mail-boundary"
	// buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString(fmt.Sprintf("Content-Type: multipart/alternative; boundary=%s\n", boundary))

	// Mail Body
	buf.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	buf.WriteString("Content-Type: text/plain; charset=\"UTF-8\"\r\n")
	buf.WriteString("Content-Transfer-Encoding: quoted-printable\r\n")
	buf.WriteString(fmt.Sprintf("\r\n%s", m.Msg))

	// Mail Attachment
	if len(m.Attachment) != 0 {
		buf.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
		buf.WriteString("Content-Type: text/plain; charset=\"UTF-8\"\r\n")
		buf.WriteString("Content-Transfer-Encoding: base64\r\n")
		buf.WriteString("Content-Disposition: attachment; filename=" + m.AttachmentName + "\r\n")
		buf.WriteString("Content-ID: <" + m.AttachmentName + ">\r\n\r\n")

		attachmentByte := make([]byte, base64.StdEncoding.EncodedLen(len(m.Attachment)))
		base64.StdEncoding.Encode(attachmentByte, m.Attachment)
		buf.Write(attachmentByte)
	}
	buf.WriteString(fmt.Sprintf("\r\n--%s", boundary))

	// Mail End
	buf.WriteString("--")

	return buf.Bytes()
}

func (m Mail) Send() (err error) {
	if m.Port == 0 {
		m.Port = 587
	}
	
	addr := fmt.Sprintf("%s:%d", m.Host, m.Port)

	auth := smtp.PlainAuth("", m.User, m.Password, m.Host)

	err = smtp.SendMail(addr, auth, m.From, m.To, m.Builder())
	return
}
