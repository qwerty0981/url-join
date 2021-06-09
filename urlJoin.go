package urlJoin

import (
	"path"
	"strings"
)

func Join(paths ...string) string {
	urlString := ""

	for _, pathPart := range paths {
		// Clean the line
		cleanedPath := path.Clean(pathPart)
		if urlString == "" || strings.HasSuffix(urlString, "/") {
			urlString += cleanedPath
		} else {
			urlString += "/" + cleanedPath
		}
	}

	return urlString
}
