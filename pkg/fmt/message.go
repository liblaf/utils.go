package fmt

import "fmt"

func On(action string, src string) string {
	return fmt.Sprintf("%s %s", action, src)
}

func To(action string, src string, dst string) string {
	return fmt.Sprintf("%s %s -> %s", action, src, dst)
}
