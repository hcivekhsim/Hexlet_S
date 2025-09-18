package utils

import (
	"fmt"
	"strings"
)


func BuildGreeting(name string, count int) string {
	return fmt.Sprintf("Hello, %s! You have %d new messages.", name, count)
	//"You have " + strconv.Itoa(count) + " new messages"
	
}
func GetHiddenCard(card string, stars int ) string{
	var c string = strings.Repeat("*", stars)
	var v string = card[len(card)-4:]
	return fmt.Sprintf("%s%s", c, v)
}