package helloworld

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Hello("", ""))
}

const (
	englishHelloPrefix = "Hello, "
	spanish            = "spanish"
	spanishHelloPrefix = "Hola, "
	french             = "french"
	frenchHelloPrefix  = "Bonjour, "
	german             = "german"
	germanHelloPrefix  = "Guten Tag, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return languageGreeting(language) + name
}

func languageGreeting(language string) string {
	prefix := englishHelloPrefix
	switch strings.ToLower(language) {
	case french:
		prefix = frenchHelloPrefix
	case german:
		prefix = germanHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}

	return prefix
}
