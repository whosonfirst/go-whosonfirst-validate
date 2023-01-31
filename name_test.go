package validate

import (
	"testing"
)

func TestValidateName(t *testing.T) {

	tests_ok := []string{
		`{"properties": { "wof:name": "Gowanus Yacht Club" }}`,
		`{"properties": { "wof:name": "Latin American Club" }}`,
	}

	tests_fail := []string{
		`{"properties": {  "wof:name": "" }}`,
		`{"properties": { }}`,
	}

	for idx, str_body := range tests_ok {

		err := ValidateName([]byte(str_body))

		if err != nil {
			t.Fatalf("Failed to validate name for test %d, %v", idx, err)
		}
	}

	for idx, str_body := range tests_fail {

		err := ValidateName([]byte(str_body))

		if err == nil {
			t.Fatalf("Expected name for test %d to fail", idx)
		}
	}

}
