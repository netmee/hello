package main

import (
	"fmt"
)

const spaninsh = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spaninshHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello return an hello string in a specific language.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spaninsh:
		prefix = spaninshHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Printf(Hello("Nico", "French"))
}
