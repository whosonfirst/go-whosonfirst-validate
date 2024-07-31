package validate

import (
	"fmt"

	"github.com/whosonfirst/go-whosonfirst-feature/properties"
)

func ValidateId(body []byte) error {

	_, err := properties.Id(body)

	if err != nil {
		return fmt.Errorf("Failed to derive wof:name from body, %w", err)
	}

	return nil
}
