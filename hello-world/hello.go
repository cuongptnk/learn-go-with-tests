package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const vietnamese = "Vietnamese"
const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"
const vietnameseHelloPrefix = "Xin chao"

func Hello(name string, language string) string {
	if (name == "") {
		name = "World"
	}

	return fmt.Sprintf("%s, %s", greetingPrefix(language), name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case vietnamese:
		prefix = vietnameseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}