package main

import (
	"fmt"
	"os"
	"strings"
)

// Color codes
var ColorCodes = map[string]string{
	"Reset":        "\033[0m",
	"Black":        "\033[30m",
	"Red":          "\033[31m",
	"Green":        "\033[32m",
	"Yellow":       "\033[33m",
	"Blue":         "\033[34m",
	"Magenta":      "\033[35m",
	"Cyan":         "\033[36m",
	"LightGray":    "\033[37m",
	"Gray":         "\033[90m",
	"LightRed":     "\033[91m",
	"LightGreen":   "\033[92m",
	"LightYellow":  "\033[93m",
	"LightBlue":    "\033[94m",
	"LightMagenta": "\033[95m",
	"LightCyan":    "\033[96m",
	"White":        "\033[97m",
	"Orange":       "\033[38;5;208m",
}

func main() {
	if len(os.Args) > 2 {
		ColorFromOutside := os.Args[1]
		textColor := ""
		if len(ColorFromOutside) > 8 {
			textColor = ColorFromOutside[8:]
		}

		if len(os.Args) == 3 {
			textFromOutside := os.Args[2]
			printColored(textFromOutside, textColor)
			fmt.Println("")

		} else if len(os.Args) == 4 {
			textFromOutside := os.Args[3]
			subStringFromOutside := os.Args[2]

			letterCount := strings.Count(textFromOutside, subStringFromOutside)
			
			for i := 0; i < letterCount; i++ {
				lenOfSubStr := len(subStringFromOutside)
				idxOfSubStrInStr := strings.Index(textFromOutside, subStringFromOutside)
				fmt.Print(textFromOutside[:idxOfSubStrInStr])
				printColored(subStringFromOutside, textColor)
				textFromOutside = textFromOutside[idxOfSubStrInStr+lenOfSubStr:]
			}
			fmt.Println(textFromOutside)
		}
	}

}

func printColored(text, textColor string) {
	titleCase := strings.Title(textColor)
	fmt.Print(ColorCodes[titleCase] + text + ColorCodes["Reset"])
}
