package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	"github.com/whosonfirst/go-whosonfirst-index"
	"github.com/whosonfirst/go-whosonfirst-index/utils"
	"github.com/whosonfirst/go-whosonfirst-log"
	"github.com/whosonfirst/warning"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {

	valid_modes := index.Modes()
	sort.Strings(valid_modes)

	str_valid_modes := strings.Join(valid_modes, ", ")
	desc_mode := fmt.Sprintf("The mode to use when indexing data. Valid modes are: %s", str_valid_modes)

	mode := flag.String("mode", "repo", desc_mode)

	liberal := flag.Bool("liberal", false, "Allow go-whosonfirst-geojson-v2 warnings (rather than explicit errors).")
	verbose := flag.Bool("verbose", false, "Be chatty about what's happening.")

	flag.Parse()

	logger := log.SimpleWOFLogger()

	writer := io.MultiWriter(os.Stdout)
	logger.AddLogger(writer, "status")

	cb := func(fh io.Reader, ctx context.Context, args ...interface{}) error {

		path, err := index.PathForContext(ctx)

		if err != nil {
			return err
		}

		ok, err := utils.IsPrincipalWOFRecord(fh, ctx)

		if err != nil {
			return err
		}

		if !ok {
			return nil
		}

		f, err := feature.LoadWOFFeatureFromReader(fh)

		if err != nil {

			logger.Warning("failed to load feature for %s because %s", path, err)

			if warning.IsWarning(err) && *liberal {
				logger.Info("error is warning and -liberal flag enabled so allowing")
			} else {
				return err
			}
		}

		if *verbose {
			logger.Status("OK %s (%s) %s", path, f.Placetype(), f.Name())
		}

		return nil
	}

	indexer, err := index.NewIndexer(*mode, cb)

	if err != nil {
		logger.Fatal("Failed to create new indexer because %s", err)
	}

	err = indexer.IndexPaths(flag.Args())

	if err != nil {
		logger.Fatal("Failed to index paths in %s mode because %s", *mode, err)
	}

	os.Exit(0)
}
