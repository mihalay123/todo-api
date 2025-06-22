package utils

import (
	"net/http"
	"strconv"
	"strings"
)

func ParseIDFromPath(r *http.Request, prefix string) (int, error) {
	idStr := strings.TrimPrefix(r.URL.Path, prefix)
	return strconv.Atoi(idStr)
}
