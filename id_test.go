package validate

import (
	"testing"
)

func TestValidateId(t *testing.T) {

	tests_ok := []string{
		`{"properties": { "wof:id": 1234 }}`,
		`{"properties": { "wof:id": -4 }}`,
		`{"properties": {  }}`,
	}

	tests_fail := []string{
		`{"properties": { "wof:id": "abscd" }}`,
		`{"properties": { "wof:id": "1234" }}`,
	}

	for idx, str_body := range tests_ok {

		err := ValidateId([]byte(str_body))

		if err != nil {
			t.Fatalf("Failed to validate wof:id for test #%d, %v", idx, err)
		}
	}

	for idx, str_body := range tests_fail {

		err := ValidateId([]byte(str_body))

		if err == nil {
			t.Fatalf("Expected wof:id for test #%d to fail", idx)
		}
	}

}
