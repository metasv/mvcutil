mvcutil
=======

[![Build Status](https://travis-ci.org/metasv/mvcutil.svg?branch=master)](https://travis-ci.org/metasv/mvcutil)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/metasv/mvcutil)

Package mvcutil provides Bitcoin (BSV) specific convenience functions and types.
A comprehensive suite of tests is provided to ensure proper functionality.  See
`test_coverage.txt` for the gocov coverage report.  Alternatively, if you are
running a POSIX OS, you can run the `cov_report.sh` script for a real-time
report.

This package was developed for mvcd, a full-node implementation of
Bitcoin (BSV). Although it was primarily written for mvcd, this package has intentionally been designed so it
can be used as a standalone package for any projects needing the functionality
provided.

## Installation and Updating

```bash
$ go get -u github.com/metasv/mvcutil
```

## License

Package mvcutil is licensed under the [copyfree](http://copyfree.org) ISC
License.
