package main

func main() {
	words := []string{
		"go", "java", "go", "python",
		"go", "java", "rust", "go",
	}
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	for _, word := range words {
		println(wordCount[word])
	}
	
	println(wordCount)
	alias := wordCount
	println(alias)

	alias["go"] = 100
	println(wordCount["go"])

}
