package main

import "fmt"

const (
	french  = "French"
	spanish = "Spanish"

	englishPrefix = "Hello"
	frenchPrefix  = "Bonjour"
	spanishPrefix = "Hola"
)

func Hello(s string, language string) string {
	if s == "" {
		s = "World"
	}
	return greeting(language) + ", " + s
}

func greeting(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("world", "english"))
}
