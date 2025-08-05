package service

import (
	"bytes"
	"cci-api/internal/config"
	"html/template"

	"github.com/resend/resend-go/v2"
)

type EmailService interface {
	SendEmail(to, subject, templateName string, data interface{}) error
}

type emailService struct {
	cfg *config.Config
}

func NewEmailService(cfg *config.Config) EmailService {
	return &emailService{cfg: cfg}
}

func (s *emailService) SendEmail(to, subject, templateName string, data interface{}) error {
	client := resend.NewClient(s.cfg.ResendAPIKey)

	templatePath := "internal/templates/" + templateName
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return err
	}

	params := &resend.SendEmailRequest{
		From:    s.cfg.ResendFrom,
		To:      []string{to},
		Subject: subject,
		Html:    body.String(),
		Cc:      s.cfg.ResendCc,
		Bcc:     s.cfg.ResendBcc,
	}

	_, err = client.Emails.Send(params)
	return err
}
