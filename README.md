# go-fakeit

[![Go Report Card](https://goreportcard.com/badge/github.com/the0rem/go-fakeit)](https://goreportcard.com/report/github.com/the0rem/go-fakeit)
[![Coverage Status](https://coveralls.io/repos/github/the0rem/go-fakeit/badge.svg?branch=master)](https://coveralls.io/github/the0rem/go-fakeit?branch=master)
[![Godoc](https://godoc.org/github.com/the0rem/go-fakeit?status.svg)](https://godoc.org/github.com/the0rem/go-fakeit)
[![Release](https://img.shields.io/github/release/the0rem/go-fakeit.svg)](https://github.com/the0rem/go-fakeit/releases/latest)

Go Fakeit provides a way to automatically fake data from Go Structs. It is designed to work with both builtin Go Data Types as well as go-openapi/strfmt Data Types.

It comes with support for Struct Field Tag interpolation to set specific fakers for Struct Fields.

## TODO
- [ ] Handle faking Interface, Map & Chanel Struct Fields
- [ ] Complete test coverage for handlers and fakers
- [ ] Add godocs for public functions
- [X] Add handling for limiting recursion depth of circular dependencies
- [X] Add Struct Field Tag option to omit Field from being faked

## Available Fakers
Fakeit implements the available fakers in the https://github.com/icrowley/fake package. Available fakers can be found  on [Godoc](https://godoc.org/github.com/the0rem/go-fakeit/fakers/)

## Available Handlers
Fakeit comes with support for all [builtin Go Types](https://golang.org/ref/spec#Types), as well as all [go-openapi/strfmt Types](https://github.com/go-openapi/strfmt). The default fakers which are used for handling each Data Type can be found in the [Godoc](https://godoc.org/github.com/the0rem/go-fakeit/handlers/)

## Using tags to specify faker rules
The `fakeit` Tag can be used on Struct Fields to specify faker behaviour. The examples below describe how to utilise the functionality.

### Example 1 - Overriding the default faker:

```
type AwesomeThing struct {
  ID            strfmt.UUID
  Name          string `fakeit:"password,min=25,max=30"`
  Description   *string `fakeit:"paragraph"`
}
```

### Example 2 - Omitting a Field form being faked:
There are cases where you don't want to populate a certain field in a struct (such as when generating data for inserting into a database). In this case the field can be annotated with a hyphen (`-`). This follows the convention set the the JSON package in Go.
```
type AwesomeThing struct {
  ID        strfmt.UUID
  Name      string
  DeletedAt *string `fakeit:"-"`
}
```
The result will be:
```
thing := AwesomeThing{}
fakeit.Fakeit(&thing)
fmt.Printf("%+v\n", thing)

// Prints output:
// &{ID:34b7a0d3-c753-11e7-9dd1-dca90491e849 Name:Stagorchid DeletedAt:<nil>}
```
