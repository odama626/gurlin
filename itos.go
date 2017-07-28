package main

func mapChar(i int) byte {
	var c = byte(0)
	switch {
	case i == 0:
		c = 45
	case i < 11:
		c = byte(i + 47)
	case i < 37:
		c = byte(i + 54)
	case i == 37:
		c = byte(95)
	case i < 64:
		c = byte(i + 59)
	case i == 64:
		c = 126
	}
	return c
}

const base = 64

// ItoS converts an integer into a base64 string with as few characters as possible
func ItoS(i int) string {
	var c int
	var out []byte

	for i != 0 {
		c = i % base
		i = i / base
		out = append(out, mapChar(c))
	}
	return string(out)
}
