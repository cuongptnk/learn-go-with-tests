package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Address Address
}

type Address struct {
	Street string
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct{
		Name string
		Input interface{}
		Expected []string
	} {
		{
			"simple case with 1 string field",
			struct {
				Name string
			}{"A"},
			[]string{"A"},
		},
		{
			"2 string fields",
			struct {
				FName string
				LName string
			}{
				"FN",
				"LN",
			},
			[]string{"FN", "LN"},
		},
		{
			"string in nested field",
			Person{
				"Name",
				Address{
					"Street",
					"City",
				},
			},
			[]string{"Name", "Street", "City"},
		},
		{
			"struct with non-string field",
			struct {
				Name string
				Age int
			}{
				"Name",
				21,
			},
			[]string{"Name"},
		},
		{
			"pointers",
			&Person{
				"Name",
				Address{
					"Street",
					"City",
				},
			},
			[]string{"Name", "Street", "City"},
		},
		{
			"slices",
			[] Person {
				{"Name1", Address{"S1", "C1"}},
				{"Name2", Address{"S2", "C2"}},
			},
			[]string{"Name1", "S1", "C1", "Name2", "S2", "C2"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(s string) {
				got = append(got, s)
			})
			if ! reflect.DeepEqual(test.Expected, got) {
				t.Errorf("want %v got %v", test.Expected, got)
			}
		})
	}
}
