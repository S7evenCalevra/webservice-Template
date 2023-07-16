package internals

type Say struct {
	greeting string
}

func Hello() Say {
	return Say{
		greeting: "Greetings. This is just a Example",
	}
}
