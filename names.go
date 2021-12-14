package validate

import (
	"github.com/whosonfirst/go-whosonfirst-feature/properties"
	"github.com/whosonfirst/go-whosonfirst-names/tags"
)

func ValidateNames(body []byte) (bool, error) {

	names := properties.Names(body)

	for tag, _ := range names {

		_, err := tags.NewLangTag(tag)

		if err != nil {
			return false, err
		}
	}

	return true, nil
}
