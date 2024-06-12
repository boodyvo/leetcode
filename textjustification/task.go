package task

func fullJustify(words []string, maxWidth int) []string {
	rows := make([]string, 0)
	currentRow := make([]string, 0)
	rowSum := 0

	for _, word := range words {
		if len(word)+len(currentRow)+rowSum <= maxWidth {
			currentRow = append(currentRow, word)
			rowSum += len(word)

			continue
		}

		rows = append(rows, extendRow(currentRow, maxWidth))

		currentRow = []string{word}
		rowSum = len(word)
	}

	if len(currentRow) > 0 {
		rows = append(rows, extendLastRow(currentRow, maxWidth))
	}

	return rows
}

func extendLastRow(words []string, maxWidth int) string {
	row := ""
	for i, word := range words {
		row += word
		if i != len(words)-1 {
			row += " "
		}
	}

	for i := len(row); i < maxWidth; i++ {
		row += " "
	}

	return row
}

func extendRow(words []string, maxWidth int) string {
	wordsLen := 0
	for _, word := range words {
		wordsLen += len(word)
	}

	leftSpaces := maxWidth - wordsLen
	numberOfSpaces := len(words) - 1
	row := ""

	for _, word := range words {
		row += word

		necessarySpaces := leftSpaces
		if numberOfSpaces > 0 {
			necessarySpaces = leftSpaces / numberOfSpaces
			if leftSpaces%numberOfSpaces != 0 {
				necessarySpaces++
			}
		}

		for j := 0; j < necessarySpaces; j++ {
			row += " "
		}

		leftSpaces -= necessarySpaces
		numberOfSpaces--
	}

	return row
}
