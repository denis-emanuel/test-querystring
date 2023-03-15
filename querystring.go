package querystring

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"strings"
)

func Parse(query string, v any) error {
	if query == "" {
		return nil
	}

	unescapedQuery, err := url.QueryUnescape(query)
	if err != nil {
		return err
	}

	URL, err := url.Parse("?" + unescapedQuery)
	if err != nil {
		return err
	}
	if URL.RawQuery == "" {
		return errors.New("invalid querystring")
	}

	// split by &
	queryParams := strings.Split(URL.RawQuery, "&")

	queryParamsMap := struct {
		Where    map[string][]string `json:"where"`
		Order    []map[string]string `json:"order"`
		Group    []string            `json:"group"`
		Limit    int                 `json:"limit"`
		Skip     int                 `json:"skip"`
		Page     int                 `json:"page"`
		PageSize int                 `json:"pageSize"`
	}{
		Where: map[string][]string{},
		Order: []map[string]string{},
		Group: []string{},
	}

	for _, queryParam := range queryParams {
		splitQueryParam := strings.Split(queryParam, "=")

		if len(splitQueryParam) < 2 {
			continue
		}

		switch splitQueryParam[0] {
		case "order":
			queryParamsMap.Order = parseOrder(splitQueryParam[1])
		case "group":
			queryParamsMap.Group = strings.Split(splitQueryParam[1], ",")
		case "limit":
			limit, err := strconv.Atoi(splitQueryParam[1])
			if err != nil {
				return err
			}
			queryParamsMap.Limit = limit
		case "pageSize":
			pageSize, err := strconv.Atoi(splitQueryParam[1])
			if err != nil {
				return err
			}
			queryParamsMap.PageSize = pageSize
		case "skip":
			skip, err := strconv.Atoi(splitQueryParam[1])
			if err != nil {
				return err
			}
			queryParamsMap.Skip = skip
		case "page":
			page, err := strconv.Atoi(splitQueryParam[1])
			if err != nil {
				return err
			}
			queryParamsMap.Page = page
		default:
			queryParamsMap.Where[splitQueryParam[0]] = strings.Split(splitQueryParam[1], ",")
		}
	}

	jsonBytes, err := json.Marshal(queryParamsMap)
	if err != nil {
		return err
	}

	json.Unmarshal(jsonBytes, &v)

	return nil
}

func parseOrder(orderStr string) []map[string]string {
	defaultDirection := "desc"
	orderItems := strings.Split(orderStr, ",")
	orderMap := make([]map[string]string, len(orderItems))
	for i, orderField := range orderItems {
		if strings.ContainsRune(orderField, '.') {
			// in case of ?orderStr=orderField.direction
			orderMap[i] = map[string]string{
				"name":  strings.Split(orderField, ".")[0],
				"value": strings.Split(orderField, ".")[1],
			}
			continue
		}
		orderMap[i] = map[string]string{"name": orderField, "value": defaultDirection}
	}
	return orderMap
}
