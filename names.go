package validate

import (
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-names/tags"
)

func ValidateNames(f geojson.Feature) (bool, error) {

	names := whosonfirst.Names(f)

	for tag, _ := range names {

		_, err := tags.NewLangTag(tag)

		if err != nil {
			return false, err
		}
	}

	return true, nil
}
