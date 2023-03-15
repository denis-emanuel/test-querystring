package querystring

type Group []string

type Order []struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Skip int
type Limit int
type Page int
type PageSize int
