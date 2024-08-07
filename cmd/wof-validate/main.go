// wof-validate is a command line tool to validate the contents of one or more whosonfirst/go-whosonfirst-iterate/v2 data sources.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/whosonfirst/go-whosonfirst-iterate/v2/iterator"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"github.com/whosonfirst/go-whosonfirst-validate"
)

func main() {

	iterator_uri := flag.String("iterator-uri", "repo://", "A valid whosonfirst/go-whosonfirst-iterate/v2 URI")

	check_id := flag.Bool("id", true, "Validate wof:id property.")
	check_name := flag.Bool("name", true, "Validate wof:name property.")
	check_placetype := flag.Bool("placetype", true, "Validate wof:placetype property.")
	check_repo := flag.Bool("repo", true, "Validate wof:repo property.")
	check_edtf := flag.Bool("edtf", true, "Validate edtf: properties.")
	check_iscurrent := flag.Bool("is-current", false, "Validate mz:is_current property.")

	check_names := flag.Bool("names", false, "Validate WOF/RFC 5646 names.")

	check_all := flag.Bool("all", false, "Enable all validation checks.")

	verbose := flag.Bool("verbose", false, "Be chatty about what's happening.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Validate the contents of one or more whosonfirst/go-whosonfirst-iterate/v2 data sources.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s path(N) path(N)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid arguments are:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *check_all {
		*check_id = true
		*check_names = true
		*check_name = true
		*check_placetype = true
		*check_repo = true
		*check_edtf = true
		*check_iscurrent = true
	}

	opts := &validate.Options{
		ValidateId:        *check_id,
		ValidateName:      *check_name,
		ValidatePlacetype: *check_placetype,
		ValidateRepo:      *check_repo,
		ValidateEDTF:      *check_edtf,
		ValidateIsCurrent: *check_iscurrent,
		ValidateNames:     *check_names,
	}

	ctx := context.Background()

	iter_cb := func(ctx context.Context, path string, fh io.ReadSeeker, args ...interface{}) error {

		_, uri_args, err := uri.ParseURI(path)

		if err != nil {
			return fmt.Errorf("Failed to parse URI '%s', %w", path, err)
		}

		body, err := validate.EnsureValidGeoJSON(fh)

		if err != nil {
			return fmt.Errorf("Failed to ensure GeoJSON for '%s', %w", path, err)
		}

		if uri_args.IsAlternate {
			return nil
		}

		err = validate.ValidateWithOptions(body, opts)

		if err != nil {
			return fmt.Errorf("Failed to validate '%s', %w", path, err)
		}

		if *verbose {
			log.Printf("OK %s\n", path)
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
