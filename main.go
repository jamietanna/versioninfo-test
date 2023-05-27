package main

import (
	"fmt"
	"runtime/debug"

	"github.com/carlmjohnson/versioninfo"
)

func main() {
	fmt.Println("Version:", versioninfo.Version)
	fmt.Println("Revision:", versioninfo.Revision)
	fmt.Println("DirtyBuild:", versioninfo.DirtyBuild)
	fmt.Println("LastCommit:", versioninfo.LastCommit)
	fmt.Printf("versioninfo.Short(): %v\n", versioninfo.Short())

	info, ok := debug.ReadBuildInfo()
	if ok {
		fmt.Printf("info: %v\n", info)
		for i, bs := range info.Settings {
			fmt.Printf("%d: %s", i, bs)
		}
	}
}
