package validate

import (
	"testing"
)

func TestValidateRepo(t *testing.T) {

	tests_ok := []string{
		`{"properties": { "wof:repo": "whosonfirst-data-admin-us" }}`,
		`{"properties": { "wof:repo": "sfomuseum-data-publicart" }}`,
	}

	tests_fail := []string{
		`{"properties": {  "wof:repo": "" }}`,
		`{"properties": { }}`,
	}

	for idx, str_body := range tests_ok {

		err := ValidateRepo([]byte(str_body))

		if err != nil {
			t.Fatalf("Failed to validate repo for test %d, %v", idx, err)
		}
	}

	for idx, str_body := range tests_fail {

		err := ValidateRepo([]byte(str_body))

		if err == nil {
			t.Fatalf("Expected repo for test %d to fail", idx)
		}
	}

}
