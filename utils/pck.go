package utils

import "fmt"

func BuildGreeting(name string, count int) string {
	return fmt.Sprintf("Hello, %s! You have %d new messages.", name, count)
}
