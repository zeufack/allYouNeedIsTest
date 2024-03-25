package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour"
)

func Hello(name string, langage string) string {
	if name == "" {
		name = "world"
	}

	// if langage == "Spanish" {
	// 	return spanishHelloPrefix + name
	// }

	// if langage == "french" {
	// 	return frenchHelloPrefix + name
	// }

	return greetingPrefix(langage) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", spanish))
}