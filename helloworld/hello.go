package helloworld

const (
	enHelloPrefix = "Hello, "
	deHelloPrefix = "Hallo, "
	frHelloPrefix = "Bonjour, "
)

func Hello(name, lang string) string {
	if name == "" {
		name = defaultGreeting(lang)
	}

	return greetingPrefix(lang) + name
}

func defaultGreeting(lang string) (word string) {
	switch lang {
	case "de":
		word = "Welt"
	case "fr":
		word = "monde"
	default:
		word = "World"
	}
	return
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case "de":
		prefix = deHelloPrefix
	case "fr":
		prefix = frHelloPrefix
	default:
		prefix = enHelloPrefix
	}
	return
}
