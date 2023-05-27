package main

import (
	"fmt"

	"github.com/carlmjohnson/versioninfo"
)

func main() {
	fmt.Println("Version:", versioninfo.Version)
	fmt.Println("Revision:", versioninfo.Revision)
	fmt.Println("DirtyBuild:", versioninfo.DirtyBuild)
	fmt.Println("LastCommit:", versioninfo.LastCommit)
}
