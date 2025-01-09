package main

import (
	"fmt"
	"strings"
	"unicode"
)

type TextRange struct {
	Start, End               int
	Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
	return position >= t.Start && position <= t.End
}

type FormattedText struct {
	plainText  string
	formatting []*TextRange
}

func NewFormattedText(plainText string) *FormattedText {
	return &FormattedText{
		plainText: plainText,
	}
}

func (f *FormattedText) Range(start, end int) *TextRange {
	r := &TextRange{start, end, false, false, false}
	f.formatting = append(f.formatting, r)
	return r
}

func (f *FormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(f.plainText); i++ {
		c := f.plainText[i]
		for _, r := range f.formatting {
			if r.Covers(i) && r.Capitalize {
				c = uint8(unicode.ToUpper(rune(c)))

			}
		}
		sb.WriteRune(rune(c))
	}
	return sb.String()
}

/**
 * The main function demonstrates the use of the Flyweight pattern.
 * Insted of converting each character to uppercase, we can now
 * capitalize a range of characters.
 * This is achieved by creating a new TextRange object and setting
 * the Capitalize property to true.
 *
 * Instead of creating a Array of lenght X equal to the length of the string
 * we are now creating a slice of TextRange objects. Which will only be created
 * when needed.
 */
func main() {
	text := "This is a string that will be capitalized"
	ft := NewFormattedText(text)
	ft.Range(10, 20).Capitalize = true // Range returns the pointer to the TextRange object, and we set it to capitalize
	fmt.Println(ft.String())           // Output: This is a STRING that will be capitalized
}
