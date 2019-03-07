package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	out = os.Stdout

	l = flag.Bool("little", false, "Little")
)

const (
	// \x1b = ESC
	// ESC[n;mH => move cursor (n, m)
	// ESCK => clear line
	escapeSequenceFormat = "\x1b[%d;%dH\x1bK"

	printDurationVertical   = 30 * time.Millisecond
	printDurationHorizontal = 0 * time.Millisecond
)

func main() {
	flag.Parse()
	if *l {
		printSmallVkgtaro()
		return
	}
	printVkgtaro()
}

func printVkgtaro() {
	clear()
	for y := 80; y > 0; y-- {
		for x := 1; x <= 15; x++ {
			fmt.Fprintf(out, escapeSequenceFormat, x, y)
			fmt.Println(vkgtaroLine(x))
			time.Sleep(printDurationHorizontal)
		}
		time.Sleep(printDurationVertical)
	}
	clear()
}

func printSmallVkgtaro() {
	clear()
	for y := 80; y > 0; y-- {
		fmt.Fprintf(out, escapeSequenceFormat, 15, y)
		fmt.Println("vkgtaro ") // last space is necessary
		time.Sleep(printDurationVertical)
	}
	clear()
}

func clear() {
	// ESC[2J => clear screen
	fmt.Fprintln(out, "\x1b[2J")
}

func vkgtaroLine(line int) string {
	const (
		vkgtaro01 = "         ##                                                   "
		vkgtaro02 = "          #                  #                                "
		vkgtaro03 = "          #                  #                                "
		vkgtaro04 = "          #                  #                                "
		vkgtaro05 = "##   ##   #    #    ### ## ######     ####   ### ##     ###   "
		vkgtaro06 = "#     #   #   #    #   #     #       #    #    ##  #   #   #  "
		vkgtaro07 = "#     #   #  #     #   #     #            #    #   #  #     # "
		vkgtaro08 = " #   #    # ##     #   #     #        #####    #      #     # "
		vkgtaro09 = " #   #    ##  #     ###      #       #    #    #      #     # "
		vkgtaro10 = " #   #    #   #    #         #      #     #    #      #     # "
		vkgtaro11 = "  # #     #    #   ####      #   #  #     #    #      #     # "
		vkgtaro12 = "  # #     #    #  #    #     #   #  #    ##    #       #   #  "
		vkgtaro13 = "   #     ###   ## #     #     ###    #### ## ######     ###   "
		vkgtaro14 = "                  #     #                                     "
		vkgtaro15 = "                   #####                                      "
	)
	switch line {
	case 1:
		return vkgtaro01
	case 2:
		return vkgtaro02
	case 3:
		return vkgtaro03
	case 4:
		return vkgtaro04
	case 5:
		return vkgtaro05
	case 6:
		return vkgtaro06
	case 7:
		return vkgtaro07
	case 8:
		return vkgtaro08
	case 9:
		return vkgtaro09
	case 10:
		return vkgtaro10
	case 11:
		return vkgtaro11
	case 12:
		return vkgtaro12
	case 13:
		return vkgtaro13
	case 14:
		return vkgtaro14
	case 15:
		return vkgtaro15
	default:
		return ""
	}
}
