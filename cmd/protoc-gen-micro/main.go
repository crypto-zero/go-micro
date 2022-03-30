package main

import (
	"runtime/debug"

	"google.golang.org/protobuf/compiler/protogen"
)

var (
	tag = "master"
)

func main() {
	info, ok := debug.ReadBuildInfo()
	if ok {
		for _, s := range info.Settings {
			if s.Key == "vcs.revision" {
				tag = s.Value
				break
			}
		}
	}
	options := &protogen.Options{}
	options.Run(func(plugin *protogen.Plugin) error {
		for _, f := range plugin.Files {
			if f.Generate {
				generateFile(plugin, f)
			}
		}
		return nil
	})
}
