package main

const (
	englishHelloPrefix = "Hello "
	banglaHelloPrefix  = "আসসালামুয়ালাইকুম "
	bangla             = "Bangla"
	spanishHelloPrefix = "Hola "
	spanish            = "Spanish"
)

func message(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case bangla:
		prefix = banglaHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	message("shovon", "")
}
