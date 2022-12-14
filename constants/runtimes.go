package constants

import "fmt"

var Runtimes = []*RuntimesType{
	{
		Name: "node",
		Url: "https://nodejs.org/dist/v$v/node-v$v-$o-$a.$e",
		Version: "19.1.0",
		Ext: "tar.xz",
		Bin: "/node-v$v-$o-$a/bin/node",
		Os: map[string]OsType{
			"windows": {
				Name: "win",
				Ext: "zip",
				Bin: "/node-v$v-$o-$a/node.exe",
			},
		},
		Arch: map[string]string{
			"amd64": "x64",
		},
	},
	{
		Name: "deno",
		Url: "https://github.com/denoland/deno/releases/$v/download/deno-$a-$o.$e",
		Version: "latest",
		Ext: "zip",
		Bin: "/deno",
		Os: map[string]OsType{
			"windows": {
				Name: "pc-windows-msvc",
			},
		},
		Arch: map[string]string{
			"amd64": "x86_64",
		},
	},
}

func GetDefinedRuntime(name string) (*RuntimesType, error) {
	for i, r := range Runtimes {
		if r.Name == name {
			return Runtimes[i], nil
		}
	}
	return nil, fmt.Errorf("cannot find runtime %s", name)
}