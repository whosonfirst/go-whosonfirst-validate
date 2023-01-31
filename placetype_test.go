package validate

import (
	"testing"
)

func TestValidatePlacetype(t *testing.T) {

	tests_ok := []string{
		`{"properties": { "wof:placetype": "custom" }}`,
		`{"properties": { "wof:placetype": "localadmin" }}`,
		`{"properties": { "wof:placetype": "microhood" }}`,
		`{"properties": { "wof:placetype": "installation" }}`,
	}

	tests_fail := []string{
		`{"properties": {  "wof:placetype": "" }}`,
		`{"properties": { "wof:placetype": "exhibition" }}`,
	}

	for idx, str_body := range tests_ok {

		err := ValidatePlacetype([]byte(str_body))

		if err != nil {
			t.Fatalf("Failed to validate placetype for test %d, %v", idx, err)
		}
	}

	for idx, str_body := range tests_fail {

		err := ValidatePlacetype([]byte(str_body))

		if err == nil {
			t.Fatalf("Expected placetype for test %d to fail", idx)
		}
	}

}
