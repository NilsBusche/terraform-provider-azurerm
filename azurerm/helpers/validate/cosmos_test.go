package validate

import "testing"

func TestCosmosAccountName(t *testing.T) {
	cases := []struct {
		Value  string
		Errors int
	}{
		{
			Value:  "foo-bar",
			Errors: 0,
		},
		{
			Value:  "foo",
			Errors: 0,
		},
		{
			Value:  "fu",
			Errors: 1,
		},
		{
			Value:  "foo_bar",
			Errors: 1,
		},
		{
			Value:  "fooB@r",
			Errors: 1,
		},
		{
			Value:  "foo-bar-foo-bar-foo-bar-foo-bar-foo-bar-foo-bar-foo-bar",
			Errors: 1,
		},
	}

	for _, tc := range cases {
		_, errors := CosmosAccountName(tc.Value, "throughput")
		if len(errors) != tc.Errors {
			t.Fatalf("Expected DatabaseCollation to trigger '%d' errors for '%s' - got '%d'", tc.Errors, tc.Value, len(errors))
		}
	}
}

func TestCosmosEntityName(t *testing.T) {
	cases := []struct {
		Value  string
		Errors int
	}{
		{
			Value:  "",
			Errors: 1,
		},
		{
			Value:  "someEntityName",
			Errors: 0,
		},
		{
			Value:  "someEntityNamesomeEntityNamesomeEntityNamesomeEntityNamesomeEntityNamesomeEntityNamesomeEntityNamesomeEntityNamesomeEntityNamesomeEntityNamesomeEntityNamesomeEntityNamesomeEntityNamesomeEntityNamesomeEntityNamesomeEntityNamesomeEntityNamesomeEntityNamesomeEntityName",
			Errors: 1,
		},
	}

	for _, tc := range cases {
		_, errors := CosmosEntityName(tc.Value, "throughput")
		if len(errors) != tc.Errors {
			t.Fatalf("Expected DatabaseCollation to trigger '%d' errors for '%s' - got '%d'", tc.Errors, tc.Value, len(errors))
		}
	}
}

func TestThroughput(t *testing.T) {
	cases := []struct {
		Value  int
		Errors int
	}{
		{
			Value:  400,
			Errors: 0,
		},
		{
			Value:  300,
			Errors: 1,
		},
		{
			Value:  450,
			Errors: 1,
		},
		{
			Value:  10000,
			Errors: 0,
		},
	}

	for _, tc := range cases {
		_, errors := Throughput(tc.Value, "throughput")
		if len(errors) != tc.Errors {
			t.Fatalf("Expected DatabaseCollation to trigger '%d' errors for '%d' - got '%d'", tc.Errors, tc.Value, len(errors))
		}
	}
}
