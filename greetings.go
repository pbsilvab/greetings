package greetings

import "fmt"

func Message(name string) string {
	message := fmt.Sprintf("hello %v wlcome", name)
	return message
}
