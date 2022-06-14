package code

const (
	englishHello    = "Hello,"
	inputDefault    = "World"
	spanishLanguage = "Spanish"
	spanishHello    = "Hola,"
	frenchLanguage  = "French"
	frenchHello     = "Bonjour,"
)

func getGreetingPrefix(language string) string {
	prefix := englishHello

	switch language {
	case spanishLanguage:
		prefix = spanishHello
	case frenchLanguage:
		prefix = frenchHello
	}

	return prefix
}

func Hello(input, language string) string {
	if input == "" {
		input = inputDefault
	}

	prefix := getGreetingPrefix(language)
	return prefix + input
}
