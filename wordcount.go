package main

func wordcount(words string) (count int) {
	lastCh := ' '
	for _, ch := range words {
		if ch != ' ' && lastCh == ' ' {
			count++
		}
		lastCh = ch
	}
	return
}
