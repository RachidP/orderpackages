package main

import (
	"fmt"

	"github.com/RachidP/orderpackages/version2/version"
)

func init() {
	fmt.Println("app/fetch-version.go ==> init() A")
}
func fetchVersion() string {
	fmt.Println("app/fetch-version.go ==> fetchVersion()")
	return version.Version
}
