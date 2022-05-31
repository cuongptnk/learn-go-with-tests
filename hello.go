package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"

func Hello(name string, language string) string {
	if (name == "") {
		name = "World"
	}
	if language == spanish {
		return fmt.Sprintf("%s, %s", spanishHelloPrefix, name)
	}
	if language == french {
		return fmt.Sprintf("%s, %s", frenchHelloPrefix, name)
	}
	return fmt.Sprintf("%s, %s", englishHelloPrefix, name)
}

func main() {
	fmt.Println(Hello("World", ""))
}