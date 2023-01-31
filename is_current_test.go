package validate

import (
	"testing"
)

func TestValidateIsCurrent(t *testing.T) {

	tests_ok := []string{
		`{"properties": { "mz:is_current": 1 }}`,
		`{"properties": { "mz:is_current": 0 }}`,
		`{"properties": { "mz:is_current": -1 }}`,
	}

	tests_fail := []string{}

	for idx, str_body := range tests_ok {

		err := ValidateIsCurrent([]byte(str_body))

		if err != nil {
			t.Fatalf("Failed to validate mz:is_current for test %d, %v", idx, err)
		}
	}

	for idx, str_body := range tests_fail {

		err := ValidateIsCurrent([]byte(str_body))

		if err == nil {
			t.Fatalf("Expected mz:is_current for test %d to fail", idx)
		}
	}

}
