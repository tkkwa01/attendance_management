package usecase

import "attendance-management/resource/mail_body"

type Mail interface {
	Send(to string, body mail_body.MailBody) error
}
