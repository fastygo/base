package title

import "fmt"

func titleTag(order int) string {
	if order < 1 || order > 6 {
		return "h2"
	}
	return fmt.Sprintf("h%d", order)
}
