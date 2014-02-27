package juice

func Parse(file string) Rules {
	var letter byte
	content := []byte(readFile(file))
	state := new(State)

	for letter, content = stripLetter(content); letter != 0; letter, content = stripLetter(content) {
		state.parse(letter)
	}

	return state.rules
}
