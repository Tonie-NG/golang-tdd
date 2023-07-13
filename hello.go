package main

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const french = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix

	case spanish:
		prefix = spanishHelloPrefix
		
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}