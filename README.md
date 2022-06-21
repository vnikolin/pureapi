<!--
Title: Storage API Go
Description: Storage api to support Pure and Solidfire Storage Arrays
Author: kgrvamsi
Updated by: vnikolin
-->

# StorageAPI

This Library supports Pure Storage array

[![GoDoc](https://godoc.org/github.com/vnikolin/pureapi?status.svg)](https://godoc.org/github.com/vnikolin/pureapi)
[![Go Report Card](https://goreportcard.com/badge/github.com/vnikolin/pureapi)](https://goreportcard.com/report/github.com/vnikolin/pureapi)

## Usage

```go
//main.go
package main

import "fmt"
import "github.com/vnikolin/pureapi"

func main() {
    client, err := NewStorageClient("192.168.1.115", "admin", "admin", "", "1.17")

	array, err := client.FetchArrayInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(array)
}
```
