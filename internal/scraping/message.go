package scraping

import (
	"regexp"
)

const (
	eInvoicingSubjectPattern = `^[0-9]{9,};.+$`
)

type Message struct {
	id             string
	to             string
	from           string
	subject        string
	date           string
	attachmentName string
	attachment     []byte
}

func (m *Message) From() string {
	return m.from
}

func (m *Message) Subject() string {
	return m.subject
}

func (m *Message) Date() string {
	return m.date
}

func (m *Message) AttachmentName() string {
	return m.attachmentName
}

func NewMessage(id, to, from, subject, date string, attachmentName string, attachment []byte) *Message {
	if !isMessageFromEInvoicing(subject) {
		return nil
	}

	return &Message{
		id:             id,
		to:             to,
		from:           from,
		subject:        subject,
		date:           date,
		attachmentName: attachmentName,
		attachment:     attachment,
	}
}

func isMessageFromEInvoicing(subject string) bool {
	expReg, err := regexp.Compile(eInvoicingSubjectPattern)
	if err != nil {
		return false
	}

	return expReg.MatchString(subject)
}
