package internal

import "errors"

func MsgHelloThereWhatsYourName(lang string) (string, error) {
	msgHello := map[string]string{
		"en": "Hello there, what's your name?",
		"pt": "Olá, qual é o seu nome?",
	}

	return msg(msgHello, lang)
}

func MsgHiNiceToMeetYou(lang string) (string, error) {
	msgHiNice := map[string]string{
		"en": "Hi {1}, nice to meet you!",
		"pt": "Oi {1}, prazer em te conhecer!",
	}

	return msg(msgHiNice, lang)
}

func msg(m map[string]string, l string) (string, error) {
	msg, found := m[l]

	if !found {
		err := errors.New("Language not found")
		return "", err
	}

	return msg, nil
}
