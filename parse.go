package juice

func ParseFile(path string) Rules {
	var letter byte
	content := []byte(readFile(path))
	state := new(State)

	for letter, content = stripLetter(content); letter != 0; letter, content = stripLetter(content) {
		state.parse(letter)
	}

	return state.rules
}
