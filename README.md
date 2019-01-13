# numeric
[![Go Report Card](https://goreportcard.com/badge/github.com/saromanov/numeric)](https://goreportcard.com/report/github.com/saromanov/numeric)
[![Build Status](https://travis-ci.org/saromanov/numeric.svg?branch=master)](https://travis-ci.org/saromanov/numeric)
[![Coverage Status](https://coveralls.io/repos/github/saromanov/numeric/badge.svg?branch=master)](https://coveralls.io/github/saromanov/numeric?branch=master)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/bb168cb23de449aea9a96d539710d1e6)](https://www.codacy.com/app/saromanov/numeric?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=saromanov/numeric&amp;utm_campaign=Badge_Grade)

Advanced mathematical library. This package contains operations with numbers, vectors, matrices and many others

## Table of Contents
* [Getting Started](#getting-started)
    + [Installing](#installing)
    + [Examples](#examples)
    + [Numeric package](#numeric-package)
    + [Linear package](#linear-package)
* [Roadmap](#roadmap)

## Getting Started

### Installing

```sh
$ go get github.com/saromanov/numeric/...
```

Supporting versions 1.10+

### Examples
Check out `examples` dir

### Numeric package

### Linear package
This package contains operations related with linear algebra. Vectors, Matrices, etc

For create a new vector you should call
```go
import "github.com/saromanov/numeric/linear"

v := linear.Vector([]float64{1, 2, 5, 8, 5, 10})
```
