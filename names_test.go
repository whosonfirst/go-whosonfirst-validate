package validate

import (
	"testing"
)

func TestValidateNames(t *testing.T) {

	tests_ok := []string{
		`{"properties": { "name:eng_x_variant": ["SFO", u"KSFO"] }}`,
		`{"properties": { "name:msa_x_preferred": ["Lapangan Terbang Antarabangsa San Francisco"] }}`,
	}

	tests_fail := []string{
		`{"properties": {  "name:123_x_bunk": [ "Testing" ]}}`,
	}

	for idx, str_body := range tests_ok {

		err := ValidateNames([]byte(str_body))

		if err != nil {
			t.Fatalf("Failed to validate names for test %d, %v", idx, err)
		}
	}

	for idx, str_body := range tests_fail {

		err := ValidateNames([]byte(str_body))

		if err == nil {
			t.Fatalf("Expected names for test %d to fail", idx)
		}
	}

}
