package scraping

import (
	"regexp"
)

const (
	eInvoicingSubjectPattern = `^[0-9]{9,};.+$`
)

type Message struct {
	id      string
	to      string
	from    string
	subject string
	date    string
}

func NewMessage(id, to, from, subject, date string) *Message {
	if !isMessageFromEInvoicing(subject) {
		return nil
	}

	return &Message{
		id:      id,
		to:      to,
		from:    from,
		subject: subject,
		date:    date,
	}
}

func isMessageFromEInvoicing(subject string) bool {
	expReg, err := regexp.Compile(eInvoicingSubjectPattern)
	if err != nil {
		return false
	}

	return expReg.MatchString(subject)
}
