package querystring

import (
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("should parse action=test,example", func(t *testing.T) {
		qs := "action=test,example"
		type Filter struct {
			Where struct {
				Action []string
			}
		}

		var expected Filter
		expected.Where = struct {
			Action []string
		}{Action: []string{"test", "example"}}

		var actual Filter
		err := Parse(qs, &actual)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v, but got %v", expected, actual)
		}
	})

	t.Run("should parse order=action,category", func(t *testing.T) {
		qs := "order=action,category"

		type Filter struct {
			Order Order `json:"order"`
		}

		var expected Filter
		expected.Order = Order{{Name: "action", Value: "desc"}, {Name: "category", Value: "desc"}}

		var actual Filter
		err := Parse(qs, &actual)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v, but got %v", expected, actual)
		}
	})

	t.Run("should parse order=action.asc", func(t *testing.T) {
		qs := "order=action.asc"

		type Filter struct {
			Order Order `json:"order"`
		}

		var expected Filter
		expected.Order = Order{{Name: "action", Value: "asc"}}

		var actual Filter
		err := Parse(qs, &actual)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v, but got %v", expected, actual)
		}
	})

	t.Run("should parse group=action,category", func(t *testing.T) {
		qs := "group=action,category"

		type Filter struct {
			Group Group `json:"group"`
		}

		expected := Filter{
			Group: []string{"action", "category"},
		}

		var actual Filter
		err := Parse(qs, &actual)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v, but got %v", expected, actual)
		}
	})

	t.Run("should parse limit=10", func(t *testing.T) {
		qs := "limit=10"

		type Filter struct {
			Limit
		}

		expected := Filter{
			Limit: 10,
		}

		var actual Filter
		err := Parse(qs, &actual)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v, but got %v", expected, actual)
		}
	})

	t.Run("should parse pageSize=10", func(t *testing.T) {
		qs := "pageSize=10"

		type Filter struct {
			PageSize
		}

		expected := Filter{
			PageSize: 10,
		}

		var actual Filter
		err := Parse(qs, &actual)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v, but got %v", expected, actual)
		}
	})

	t.Run("should parse page=1", func(t *testing.T) {
		qs := "page=1"

		type Filter struct {
			Page
		}

		expected := Filter{
			Page: 1,
		}

		var actual Filter
		err := Parse(qs, &actual)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v, but got %v", expected, actual)
		}
	})

	t.Run("should parse skip=10", func(t *testing.T) {
		qs := "skip=10"

		type Filter struct {
			Skip
		}

		expected := Filter{
			Skip: 10,
		}

		var actual Filter
		err := Parse(qs, &actual)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v, but got %v", expected, actual)
		}
	})

	t.Run("should parse ?action=test,example&order=action,category&group=action&limit=10&skip=5", func(t *testing.T) {
		qs := "action=test,example&order=action,category&group=action&limit=10&skip=5"

		type Filter struct {
			Where struct {
				Action []string `json:"action"`
			} `json:"where"`
			Order
			Group
			Skip
			Limit
		}

		expected := Filter{
			Where: struct {
				Action []string `json:"action"`
			}{Action: []string{"test", "example"}},
			Order: Order{
				{Name: "action", Value: "desc"},
				{Name: "category", Value: "desc"},
			},
			Group: Group{"action"},
			Limit: 10,
			Skip:  5,
		}

		var actual Filter
		err := Parse(qs, &actual)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v, but got %v", expected, actual)
		}
	})

	t.Run("should parse urlEncoded ?action=test,example&order=action,category&group=action&limit=10&skip=5", func(t *testing.T) {
		qs := "action%3Dtest%2Cexample%26order%3Daction%2Ccategory%26group%3Daction%26limit%3D10%26skip%3D5"

		type Filter struct {
			Where struct {
				Action []string `json:"action"`
			} `json:"where"`
			Order
			Group
			Skip
			Limit
		}

		expected := Filter{
			Where: struct {
				Action []string `json:"action"`
			}{Action: []string{"test", "example"}},
			Order: Order{
				{Name: "action", Value: "desc"},
				{Name: "category", Value: "desc"},
			},
			Group: Group{"action"},
			Limit: 10,
			Skip:  5,
		}

		var actual Filter
		err := Parse(qs, &actual)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v, but got %v", expected, actual)
		}
	})
}
