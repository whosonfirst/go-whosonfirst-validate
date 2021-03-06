# go-whosonfirst-validate

Go package for validating Who's On First documents

## Install

You will need to have both `Go` (version [1.12](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make tools
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Background

So far this is a tool built and developed out of expediency (translation: why is `x` broken?). So far it validates the things that have needed to be validated.

Going forward it would be nice to imagine this as both a general-purpose validation tool and a core piece of a Go "export" package for Who's On First documents. At present exporting WOF documents is handled exclusively by the `py-mapzen-whosonfirst-export` and `py-mapzen-whosonfirst-geojson` libraries and it would be nice to have something in another language (and one that can generate pre-compiled binaries).

## Tools

### wof-validate-index

This tool will attempt to load all the (principal) WOF documents (using `go-whosonfirst-geojson-v2`) passed to it using a `go-whosonfirst-index` indexer.

```
$> ./bin/wof-validate-index -h
Usage of ./bin/wof-validate-index:
  -liberal
    	Allow go-whosonfirst-geojson-v2 warnings (rather than explicit errors).
  -mode string
    	The mode to use when indexing data. Valid modes are: directory, feature, feature-collection, files, geojson-ls, meta, path, repo, sqlite (default "repo")
  -names
    	Validate WOF/RFC 5646 names.	
  -verbose
    	Be chatty about what's happening.
```

For example:

```
$> ./bin/wof-validate-index /usr/local/data/whosonfirst-data
...time passes
```

Assuming everything loads successfully you won't see any output (unless you've passed the `-verbose` flag (in which case you'll see _a lot_ of output)).

Or this:

```
> ./bin/wof-validate-index -names /usr/local/data/whosonfirst-data
error: Failed to parse name tag for /usr/local/data/whosonfirst-data/data/112/585/728/5/1125857285.geojson, because Failed to parse language tag 'eng_v_variant'
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-geojson-v2
* https://github.com/whosonfirst/go-whosonfirst-index