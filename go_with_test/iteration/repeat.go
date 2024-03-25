package iteration

func Repeat(character string, time int) string {
	var repeat string
	for i := 0; i < time; i++ {
		repeat += character
	}
	return repeat
}
