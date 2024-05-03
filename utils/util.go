package utils

import (
	"net/url"
	"strconv"
)

func GetPagination(query url.Values) (int64, int64, int64, int64) {
	str_page := query.Get("page")
	str_per_page := query.Get("perPage")

	page, err := strconv.Atoi(str_page)
	if err != nil {
		page = 1
	}

	per_page, err := strconv.Atoi(str_per_page)
	if err != nil {
		per_page = 10
	}

	skip := int64(per_page * (page - 1))
	limit := int64(per_page)

	return int64(page), int64(per_page), skip, limit
}
