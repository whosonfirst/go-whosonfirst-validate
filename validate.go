package validate

import (
	"fmt"
	"github.com/paulmach/orb/geojson"
	"io"
)

type Options struct {
	ValidateNames bool
	ValidateEDTF bool
}

func EnsureValidGeoJSON(r io.Reader) ([]byte, error) {

	body, err := io.ReadAll(r)

	if err != nil {
		return nil, fmt.Errorf("Failed to read body, %w", err)
	}

	_, err = geojson.UnmarshalFeature(body)
	
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal body, %w", err)
	}

	return body, nil
}

func Validate(body []byte, options *Options) error {

	if options.ValidateNames {
		
		_, err := ValidateNames(body)
		
		if err != nil {
			return fmt.Errorf("Failed to parse name tag for body, because %s", err)
		}
	}

	return nil
}
