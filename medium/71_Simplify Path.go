package medium

import (
	"strings"
)

func simplifyPath(path string) string {
	items := strings.Split(path, "/")
	tmp := []string{}

	for _, item := range items {
		if item == "." || item == "" {
			continue
		}
		if item == ".." {
			if len(tmp) > 0 {
				tmp = tmp[:len(tmp) - 1]
			}
			continue
		}
		tmp = append(tmp, item)
	}

	re := strings.Join(tmp, "/")
	re = "/" + re
	return re
}
