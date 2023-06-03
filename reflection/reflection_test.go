package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	t.Run("cases", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"struct with one string field",
				struct {
					Name string
				}{"John"},
				[]string{"John"},
			},
			{
				"struct with two string fields",
				struct {
					Name string
					City string
				}{"Anna", "Berlin"},
				[]string{"Anna", "Berlin"},
			},
			{
				"struct with two string fields",
				struct {
					Name string
					City string
					Age  int
				}{"Anna", "Berlin", 25},
				[]string{"Anna", "Berlin"},
			},
			{
				"nested fields",
				Person{
					"John",
					Profile{33, "London"},
				},
				[]string{"John", "London"},
			},
			{
				"pointers to things",
				&Person{
					"John",
					Profile{33, "London"},
				},
				[]string{"John", "London"},
			},
			{
				"slices",
				[]Profile{
					{33, "London"},
					{34, "Berlin"},
				},
				[]string{"London", "Berlin"},
			},
			{
				"arrays",
				[2]Profile{
					{33, "London"},
					{34, "Berlin"},
				},
				[]string{"London", "Berlin"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				Walk(test.Input, func(input string) {
					got = append(got, input)
				})
				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("got %v want %v", got, test.ExpectedCalls)
				}
			})
		}
	})

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		channel := make(chan Profile)
		go func() {
			channel <- Profile{27, "London"}
			channel <- Profile{35, "Berlin"}
			close(channel)
		}()

		var got []string
		want := []string{"London", "Berlin"}
		Walk(channel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		Walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, arr []string, targ string) {
	t.Helper()
	contains := false
	for _, val := range arr {
		if val == targ {
			contains = true
			break
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", arr, targ)
	}
}
