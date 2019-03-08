package vkgtaro

import (
	"fmt"
	"io"
	"time"
)

const (
	escapeSequenceFormat = "\x1b[%d;%dH\x1bK"

	printDurationVertical       = 30 * time.Millisecond
	printDurationHorizontal     = 0 * time.Millisecond
	printDurationHorizontalSlow = 30 * time.Millisecond
)

var (
	vkgtaroLines = []string{
		"         ##                                                   ",
		"          #                  #                                ",
		"          #                  #                                ",
		"          #                  #                                ",
		"##   ##   #    #    ### ## ######     ####   ### ##     ###   ",
		"#     #   #   #    #   #     #       #    #    ##  #   #   #  ",
		"#     #   #  #     #   #     #            #    #   #  #     # ",
		" #   #    # ##     #   #     #        #####    #      #     # ",
		" #   #    ##  #     ###      #       #    #    #      #     # ",
		" #   #    #   #    #         #      #     #    #      #     # ",
		"  # #     #    #   ####      #   #  #     #    #      #     # ",
		"  # #     #    #  #    #     #   #  #    ##    #       #   #  ",
		"   #     ###   ## #     #     ###    #### ## ######     ###   ",
		"                  #     #                                     ",
		"                   #####                                      ",
	}
)

var _ Vkgtaro = (*vkgtaroImpl)(nil)

// Vkgtaro is vkg service
type Vkgtaro interface {
	Move()
	MoveLittle()
	MoveSlow()
}

type vkgtaroImpl struct {
	out io.Writer
}

// New returns vkgtaro
func New(out io.Writer) Vkgtaro {
	return &vkgtaroImpl{out: out}
}

// Move moves normally
func (v *vkgtaroImpl) Move() {
	v.clear()
	for y := 80; y > 0; y-- {
		for x := 1; x <= 15; x++ {
			fmt.Fprintf(v.out, escapeSequenceFormat, x, y)
			fmt.Println(vkgtaroLine(x))
			time.Sleep(printDurationHorizontal)
		}
		time.Sleep(printDurationVertical)
	}
	v.clear()
}

// MoveLittle little version
func (v *vkgtaroImpl) MoveLittle() {
	v.clear()
	for y := 80; y > 0; y-- {
		fmt.Fprintf(v.out, escapeSequenceFormat, 15, y)
		fmt.Println("vkgtaro ") // last space is necessary
		time.Sleep(printDurationVertical)
	}
	v.clear()
}

// MoveSlow slowly
func (v *vkgtaroImpl) MoveSlow() {
	v.clear()
	for y := 80; y > 0; y-- {
		for x := 1; x <= 15; x++ {
			fmt.Fprintf(v.out, escapeSequenceFormat, x, y)
			fmt.Println(vkgtaroLine(x))
			time.Sleep(printDurationHorizontalSlow)
		}
		time.Sleep(printDurationVertical)
	}
	v.clear()
}

func (v *vkgtaroImpl) clear() {
	fmt.Fprintln(v.out, "\x1b[2J")
}

func vkgtaroLine(line int) string {
	if line < 1 || 15 < line {
		return ""
	}
	return vkgtaroLines[line-1]
}
