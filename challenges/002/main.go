package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" || mToSend.sender.number == 0 || mToSend.recipient.name == "" || mToSend.recipient.number == 0 {
		return false
	}
	return true
}

func main() {
	mToSend := messageToSend{
		message: "Hello World",
		sender: user{
			name:   "John Doe",
			number: 42,
		},
		recipient: user{
			name:   "Doe John",
			number: 24,
		},
	}
	canSendMessage(mToSend)
}
