package favicon

import "fmt"

func svgURL(letter rune) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/FortAwesome/Font-Awesome/6.x/svgs/solid/%c.svg", letter)
}
