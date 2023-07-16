package internals

import "testing"

func TestHello(t *testing.T) {

	// Create our test table
	testTable := []struct {
		greeting string
	}{
		{
			greeting: "Greetings. This is just a Example",
		},
	}
	for _, test := range testTable {
		if test.greeting != Hello().greeting {
			t.Errorf("Hello() = %v, want %v", Hello().greeting, test.greeting)
		}
	}
}
