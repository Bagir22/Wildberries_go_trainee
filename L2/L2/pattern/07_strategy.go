package main

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
	Стратегия определяет семейство схожих алгоритмов и помещает каждый из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

	+
	1. Можно заменить один алгоритм другим, не изменяя структуру контекста. Это обеспечивает гибкость и расширяемость приложения.
	2. Каждый алгоритм инкапсулирован в собственном классе, что делает код более модульным и упрощает добавление новых алгоритмов.

	-
	1. Усложняет программу за счёт дополнительных классов.
*/

import "fmt"

type MessagingStrategy interface {
	SendMessage(message string)
}

type EmailMessaging struct{}

func (em *EmailMessaging) SendMessage(message string) {
	fmt.Println("Sending email:", message)
}

type SMSMessaging struct{}

func (sms *SMSMessaging) SendMessage(message string) {
	fmt.Println("Sending SMS:", message)
}

type Messenger struct {
	Strategy MessagingStrategy
}

func (m *Messenger) SetMessagingStrategy(strategy MessagingStrategy) {
	m.Strategy = strategy
}

func (m *Messenger) Send(message string) {
	m.Strategy.SendMessage(message)
}

func main() {
	messenger := &Messenger{}

	emailMessaging := &EmailMessaging{}
	smsMessaging := &SMSMessaging{}

	messenger.SetMessagingStrategy(emailMessaging)
	messenger.Send("Hello, this is an email message.")

	messenger.SetMessagingStrategy(smsMessaging)
	messenger.Send("Hi, this is an SMS message.")
}
