# Thargo /tɑːr ɡəʊ/ (tar-go) [![Build Status](https://travis-ci.org/EMSSConsulting/Thargo.svg?branch=master)](https://travis-ci.org/EMSSConsulting/Thargo) [![](https://godoc.org/github.com/emssconsulting/thargo?status.svg)](http://godoc.org/github.com/emssconsulting/thargo)
**A tar+gzip compression and extraction library for Golang**

Thargo is a library for Go which provides everything necessary for creating and decompressing
tar archives within your applications. It sports an easy to use and incredibly powerful API
while simplifying the boilerplate necessary for managing archives.

## Using Thargo
Thargo's API is broken up into three primary components, you have an Archive which is pretty
much exactly what you'd expect and targets which are a source of entries.

An archive provides methods for adding items to the archive as well as enumerating the items
contained within it. It is what you will generally be instantiating and using to conduct many
of your management tasks.

Targets provide a high-level abstraction over various data sources, your filesystem for example,
and allow you to quickly and easily add a group of items to your archive. These items are each
represented as an Entry, which consists of a header and data payload.

```go
package main

import (
  "fmt"
  "github.com/EMSSConsulting/thargo"
)

func main() {
  archive, err := thargo.NewArchiveFile("test.tar.gz", nil)
  if err != nil {
    fmt.Fatal(err)
  }
  
  defer archive.Close()
  
  if err := archive.Add(&thargo.FileSystemTarget{
    Path: "test.txt",
  }); err != nil {
    fmt.Fatal(err)
  }
  
  if err := archive.Extract(func(entry thargo.SaveableEntry) error {
    return entry.Save("extracted/")
  }); err != nil {
    fmt.Fatal(err)
  }
}
```