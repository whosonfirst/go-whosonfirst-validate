# go-whosonfirst-valida

Go package for validating Who's On First documents

## Install

You will need to have both `Go` (specifically a version of Go more recent than 1.7 so let's just assume you need [Go 1.11](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Tools

### wof-validate-index

This tool will attempt to load all the (principal) WOF documents (using `go-whosonfirst-geojson-v2`) passed to it using a `go-whosonfirst-index` indexer.

```
./bin/wof-validate-index -h
Usage of ./bin/wof-validate-index:
  -liberal
    	Allow go-whosonfirst-geojson-v2 warnings (rather than explicit errors).
  -mode string
    	The mode to use when indexing data. Valid modes are: directory, feature, feature-collection, files, geojson-ls, meta, path, repo, sqlite (default "repo")
  -verbose
    	Be chatty about what's happening.
```

For example:

```
./bin/wof-validate-index /usr/local/data/whosonfirst-data
...time passes
```

Assuming everything loads successfully you won't see any output (unless you've passed the `-verbose` flag (in which case you'll see _a lot_ of output)).

## See also

* https://github.com/whosonfirst/go-whosonfirst-geojson-v2
* https://github.com/whosonfirst/go-whosonfirst-index