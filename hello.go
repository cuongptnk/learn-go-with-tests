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

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case vietnamese:
		prefix = vietnameseHelloPrefix
	}

	return fmt.Sprintf("%s, %s", prefix, name)
}

func main() {
	fmt.Println(Hello("World", ""))
}