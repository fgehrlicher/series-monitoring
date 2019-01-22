package Maintenance

import (
	"strings"
	"fmt"
	"github.com/fatih/color"
)

func wrapInSuperFancyBox(text []string, boxTitle string) {
	maxLength := len(boxTitle)
	for _, element := range text {
		currentElementLength := len(element)
		if currentElementLength > maxLength {
			maxLength = currentElementLength
		}
	}

	horizontalSideBorder := color.RedString("║")

	topOuterHeaderDivider := color.RedString("╔" + strings.Repeat("═", maxLength+2) + "╗")
	headerDivider := color.RedString("╠" + strings.Repeat("═", maxLength+2) + "╣")
	BottomOuterHeaderDivider := color.RedString("╚" + strings.Repeat("═", maxLength+2) + "╝")

	sidePadding := strings.Repeat(" ", (maxLength-len(boxTitle))/2)
	rightPaddingAddition := ""
	if (maxLength-len(boxTitle))%2 != 0 {
		rightPaddingAddition = " "
	}

	fmt.Println(topOuterHeaderDivider)
	fmt.Println(horizontalSideBorder + sidePadding + " " + color.GreenString(boxTitle) + " " + sidePadding + rightPaddingAddition + horizontalSideBorder)
	fmt.Println(headerDivider)

	for i := 0; i < len(text); i++ {
		currentElementLength := len(text[i])

		sidePadding = strings.Repeat(" ", (maxLength-currentElementLength)/2)
		if (maxLength-currentElementLength)%2 == 0 {
			rightPaddingAddition = ""
		} else {
			rightPaddingAddition = " "
		}
		fmt.Println(horizontalSideBorder + sidePadding + " " + text[i] + " " + sidePadding + rightPaddingAddition + horizontalSideBorder)
	}
	fmt.Println(BottomOuterHeaderDivider)
}
