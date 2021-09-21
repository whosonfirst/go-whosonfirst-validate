package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	"github.com/whosonfirst/go-whosonfirst-iterate/emitter"
	"github.com/whosonfirst/go-whosonfirst-iterate/iterator"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"github.com/whosonfirst/go-whosonfirst-validate"
	"github.com/whosonfirst/warning"
	"io"
	"log"
	"os"
)

func main() {

	iterator_uri := flag.String("iterator-uri", "repo://", "A valid whosonfirst/go-whosonfirst-iterate URI")

	check_names := flag.Bool("names", false, "Validate WOF/RFC 5646 names.")

	liberal := flag.Bool("liberal", false, "Allow go-whosonfirst-geojson-v2 warnings (rather than explicit errors).")
	verbose := flag.Bool("verbose", false, "Be chatty about what's happening.")

	flag.Parse()

	ctx := context.Background()

	iter_cb := func(ctx context.Context, fh io.ReadSeeker, args ...interface{}) error {

		path, err := emitter.PathForContext(ctx)

		if err != nil {
			return fmt.Errorf("Failed to derive path from context, %w", err)
		}

		_, uri_args, err := uri.ParseURI(path)

		if err != nil {
			return fmt.Errorf("Failed to parse URI '%s', %w", path, err)
		}

		if uri_args.IsAlternate {
			return nil
		}

		f, err := feature.LoadWOFFeatureFromReader(fh)

		if err != nil {

			if warning.IsWarning(err) && *liberal {
				// log.Printf("error is warning and -liberal flag enabled so allowing")
			} else {
				return fmt.Errorf("Failed to load feature for '%s', %w", path, err)
			}
		}

		// START OF put this in a package method (with options (that reads []byte instead of f))

		if *check_names {

			_, err := validate.ValidateNames(f)

			if err != nil {
				return fmt.Errorf("Failed to parse name tag for %s, because %s", path, err)
			}
		}

		// END OF put this in a package method (with options (that reads []byte instead of f))

		if *verbose {
			log.Printf("OK %s (%s) %s", path, f.Placetype(), f.Name())
		}

		return nil
	}

	iter, err := iterator.NewIterator(ctx, *iterator_uri, iter_cb)

	if err != nil {
		log.Fatalf("Failed to create new indexer because %v", err)
	}

	iterator_sources := flag.Args()

	err = iter.IterateURIs(ctx, iterator_sources...)

	if err != nil {
		log.Fatalf("Failed to iterate URIs, %v", err)
	}

	os.Exit(0)
}
