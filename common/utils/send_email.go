package utils

import (
	"fmt"
	"ladipage_server/common/configs"
	"net/smtp"
	"strings"
)

func SendEmail(toAddress, subject, content string) error {
	cf := configs.Get()
	email := cf.Email
	pass := cf.AppKey
	smtpHost := cf.SmtpHost
	smtpPort := cf.SmtpPort

	if strings.TrimSpace(toAddress) == "" {
		return fmt.Errorf("recipient email address is empty")
	}
	if strings.TrimSpace(subject) == "" {
		return fmt.Errorf("email subject is empty")
	}
	if strings.TrimSpace(content) == "" {
		return fmt.Errorf("email content is empty")
	}

	message := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", email, toAddress, subject, content)

	auth := smtp.PlainAuth("", email, pass, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, email, []string{toAddress}, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email")
	}

	return nil
}
