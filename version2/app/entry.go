package main

import "fmt"

func init() {
	fmt.Println("app/fetch-version.go ==> init()")
}

var myVersion = fetchVersion()

func main() {
	fmt.Println("app/fetch-version.go ==> fetchVersion()")
	fmt.Println("version ===> ", myVersion)
}
