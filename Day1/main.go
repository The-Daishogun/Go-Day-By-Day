package main

import "fmt"

const (
	englishGreetingWord = "Hello, "
	spanishGreetingWord = "Hola, "
	frenchGreetingWord  = "Bonjour, "
	persianGreetingWord = "سلام، "
)

const (
	french  = "French"
	persian = "Persian"
	spanish = "Spanish"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World!"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchGreetingWord
	case spanish:
		prefix = spanishGreetingWord
	case persian:
		prefix = persianGreetingWord
	default:
		prefix = englishGreetingWord
	}
	return
}

func main() {
	fmt.Println(Hello("The_Daishogun", ""))
}
