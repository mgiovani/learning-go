package main

import "fmt"

const spanish = "Spanish"
const portuguese = "Portuguese"

const englishHello = "Hello, "
const spanishHello = "Hola, "
const portugueseHello = "Olá, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHello
	case portuguese:
		prefix = portugueseHello
	default:
		prefix = englishHello
	}
	return
}

func main(){
	fmt.Println(Hello("world", ""))
}