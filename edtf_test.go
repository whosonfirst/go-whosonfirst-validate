package validate

import (
	"testing"
)

func TestValidateEDTF(t *testing.T) {

	tests_ok := []string{
		`{"properties": { "edtf:inception": "2023-01-31", "edtf:cessation": ".." }}`,
		`{"properties": { "edtf:inception": "2022-01-31", "edtf:cessation": "", "deprecated": "2023-01-01" }}`,
	}

	tests_fail := []string{
		`{"properties": { "edtf:inception": "2023-01-31", "edtf:cessation": "..." }}`,
		`{"properties": { "edtf:inception": "December 21", "edtf:cessation": "", "deprecated": "2023-01-01" }}`,
	}

	for idx, str_body := range tests_ok {

		err := ValidateEDTF([]byte(str_body))

		if err != nil {
			t.Fatalf("Failed to validate EDTF for test %d, %v", idx, err)
		}
	}

	for idx, str_body := range tests_fail {

		err := ValidateEDTF([]byte(str_body))

		if err == nil {
			t.Fatalf("Expected EDTF for test %d to fail", idx)
		}
	}

}
