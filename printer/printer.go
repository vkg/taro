package printer

import (
	"fmt"
	"io"
	"time"
	"unicode/utf8"

	"github.com/vkg/taro/window"
)

const (
	escapeSequenceFormat = "\x1b[%d;%dH\x1bK"

	printDurationVertical   = 30 * time.Millisecond
	printDurationHorizontal = 0 * time.Millisecond
)

// Source is print source characters
type Source struct {
	source   []string
	maxWidth int
	height   int
}

// NewSource returns source
func NewSource(source []string) *Source {
	s := Source{source: source, height: len(source)}
	maxWidth := 0
	for _, line := range source {
		width := utf8.RuneCountInString(line)
		if width > maxWidth {
			maxWidth = width
		}
	}
	s.maxWidth = maxWidth
	return &s
}

func (s *Source) line(line int) string {
	if line < 1 || s.height < line {
		return ""
	}
	return s.source[line-1]
}

// Printer can print given strings
type Printer struct {
	out io.Writer

	screenHeight int // $LINES
	screenWidth  int // $COLUMNS
}

// New returns printer instance
func New(out io.Writer) *Printer {
	screenHeight, screenWidth, _ := window.GetSize()
	return &Printer{out: out, screenHeight: screenHeight, screenWidth: screenWidth}
}

// PrintAndMove runs given string
func (p *Printer) PrintAndMove(s *Source) {
	p.clear()
	height := s.height
	if p.screenHeight < s.height {
		height = p.screenHeight
	}
	for x := p.screenWidth; x > 0; x-- {
		for y := 1; y <= height; y++ { // (x, y)
			fmt.Fprintf(p.out, escapeSequenceFormat, y, x)
			fmt.Fprintf(p.out, s.line(y))
			time.Sleep(printDurationHorizontal)
		}
		time.Sleep(printDurationVertical)
	}
	p.clear()
}

// Clear clears screen
func (p *Printer) clear() {
	fmt.Fprintln(p.out, "\x1b[2J")
}
