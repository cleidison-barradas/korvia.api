package service

type WhatsappSender interface {
	SendText(to string, text string) error
	SendButtons(to string, text string, options []string) error
}