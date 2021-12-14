# go-whosonfirst-validate

Go package for validating Who's On First documents

## Documentation

Documentation is incomplete.

[![Go Reference](https://pkg.go.dev/badge/github.com/whosonfirst/go-whosonfirst-validate.svg)](https://pkg.go.dev/github.com/whosonfirst/go-whosonfirst-validate)

## Background

So far this is a tool built and developed out of expediency (translation: why is `x` broken?). So far it validates the things that have needed to be validated.

Going forward it would be nice to imagine this as both a general-purpose validation tool and a core piece of a Go "export" package for Who's On First documents. At present exporting WOF documents is handled exclusively by the `py-mapzen-whosonfirst-export` and `py-mapzen-whosonfirst-geojson` libraries and it would be nice to have something in another language (and one that can generate pre-compiled binaries).

## Tools

### wof-validate

This tool will attempt to load all the (principal) WOF documents (using `go-whosonfirst-geojson-v2`) passed to it using a `go-whosonfirst-iterate` iterator.

```
$> ./bin/wof-validate -h
Validate the contents of one or more whosonfirst/go-whosonfirst-iterate/v2 data sources.
Usage:
	 ./bin/wof-validate path(N) path(N)
Valid arguments are:
  -all
    	Enable all validation checks.
  -iterator-uri string
    	A valid whosonfirst/go-whosonfirst-iterate/v2 URI (default "repo://")
  -names
    	Validate WOF/RFC 5646 names.
  -verbose
    	Be chatty about what's happening.
```

For example:

```
$> ./bin/wof-validate /usr/local/data/whosonfirst-data
...time passes
```

Assuming everything loads successfully you won't see any output (unless you've passed the `-verbose` flag (in which case you'll see _a lot_ of output)).

Or this:

```
> ./bin/wof-validate -names /usr/local/data/whosonfirst-data
error: Failed to parse name tag for /usr/local/data/whosonfirst-data/data/112/585/728/5/1125857285.geojson, because Failed to parse language tag 'eng_v_variant'
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-iterate/v2
* https://github.com/whosonfirst/go-whosonfirst-feature
* https://github.com/whosonfirst/go-whosonfirst-names