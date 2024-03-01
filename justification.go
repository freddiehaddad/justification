package justification

import (
	"strings"
	"sync"
)

type lineDescriptor struct {
	number int
	data   string
}

func processIncomingLines(
	wg *sync.WaitGroup,
	input chan lineDescriptor,
	output []string,
) {
	for line := range input {
		output[line.number] = line.data
	}
	wg.Done()
}

// leftJustifyLine creates a string from words with a single space
// in between. If the length of the words with spaces does not equal
// the width, additional spaces are appended after the last word such
// that the length of the string matches width.
func leftJustifyLine(words []string, width int) string {
	var sb strings.Builder
	var n int
	var spaces int
	var extraSpacesNeeded int

	if len(words) == 0 {
		return strings.Repeat(" ", width)
	}

	for _, word := range words {
		n += len(word)
	}

	spaces = len(words) - 1
	n += spaces

	sb.WriteString(words[0])
	for _, word := range words[1:] {
		sb.WriteByte(' ')
		sb.WriteString(word)
	}

	extraSpacesNeeded = width - n
	sb.WriteString(strings.Repeat(" ", extraSpacesNeeded))

	return sb.String()
}

// fullJustifyLine creates a string from words with a single space
// in between. If the length of the words with spaces does not equal
// the width, additional spaces are evenly distributed between words
// starting from the left until the string length matches width.
func fullJustifyLine(words []string, width int) string {
	var sb strings.Builder
	var n int
	var spaces int
	var extraSpacesNeeded int
	var evenlyDistributedSpaces int
	var oddlyDistributedSpaces int

	if len(words) == 0 {
		return strings.Repeat(" ", width)
	}

	sb.WriteString(words[0])
	if len(words) == 1 {
		spaces = width - len(words[0])
		sb.WriteString(strings.Repeat(" ", spaces))
		return sb.String()
	}

	for _, word := range words {
		n += len(word)
	}

	spaces = len(words) - 1
	n += spaces
	extraSpacesNeeded = width - n
	evenlyDistributedSpaces = extraSpacesNeeded / (len(words) - 1)
	oddlyDistributedSpaces = extraSpacesNeeded % (len(words) - 1)

	i := 1
	for p := 0; p < oddlyDistributedSpaces; p++ {
		sb.WriteByte(' ') // regular space
		sb.WriteByte(' ') // extra space for uneven distribution
		sb.WriteString(strings.Repeat(" ", evenlyDistributedSpaces))
		sb.WriteString(words[i])
		i++
	}

	for _, word := range words[i:] {
		sb.WriteByte(' ') // regular space
		sb.WriteString(strings.Repeat(" ", evenlyDistributedSpaces))
		sb.WriteString(word)
	}

	return sb.String()
}

// FullJustifiy splits words into an array of strings where each string
// contains the maximum number of words separated by a single space
// without the string length exceeding maxWidth.  Extra spaces to satisfy
// maxWidth are evenly distributed starting from the beginning of the string.
// The function behaves similar to strings.Join with the added spaces
// for the width constraint.
func FullJustifiy(words []string, maxWidth int) []string {
	var processingWaitGroup sync.WaitGroup
	var processedWaitGroup sync.WaitGroup

	processedLines := make(chan lineDescriptor)

	var start int
	var lineWidth int
	var lineCount int

	// Assume worst case to memory allocation.  Costs most memory, but
	// improves performance by not requiring append calls and lock
	// acquisition.
	justified := make([]string, len(words))

	processedWaitGroup.Add(1)
	go processIncomingLines(&processedWaitGroup, processedLines, justified)

	lineWidth = len(words[0])
	for i := 1; i < len(words); i++ {
		// accumulate words
		wordWidth := len(words[i]) + 1
		if lineWidth+wordWidth <= maxWidth {
			lineWidth += wordWidth
			continue
		}

		// process the words for the current line
		processingWaitGroup.Add(1)
		go func(line, width, start, end int) {
			var s string
			var ld lineDescriptor

			s = fullJustifyLine(words[start:end], width)

			ld.number = line
			ld.data = s

			processedLines <- ld
			processingWaitGroup.Done()
		}(lineCount, maxWidth, start, i)

		// reset for new line
		start = i
		lineCount++

		// current word to be part of next line
		lineWidth = len(words[i])
	}

	// remaining last words are left justified
	if lineWidth > 0 {
		processingWaitGroup.Add(1)
		go func(line, width, start, end int) {
			var s string
			var ld lineDescriptor

			s = leftJustifyLine(words[start:end], width)

			ld.number = line
			ld.data = s

			processedLines <- ld
			processingWaitGroup.Done()
		}(lineCount, maxWidth, start, len(words))
		lineCount++
	}

	processingWaitGroup.Wait()
	close(processedLines)

	processedWaitGroup.Wait()

	justified = justified[:lineCount]
	return justified
}
