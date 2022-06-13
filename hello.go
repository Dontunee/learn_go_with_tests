package main

const namePrefix = "Hello,"

func Hello(input string) string {
	if input == "" {
		input = "World"
	}
	return namePrefix + input
}

func main() {
}
