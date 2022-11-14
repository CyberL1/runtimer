package constants

var Runtimes = []RuntimesType{
	{
		Name: "nodejs",
		Url:  "https://nodejs.org/dist/v$v/node-v$v-$o-$a.$e",
	},
	{
		Name: "deno",
		Url:  "https://github.com/denoland/deno/releases/download/v$v/deno-$a-$o.$e",
	},
}

func GetDefinedRuntime(name string) RuntimesType {
		for i, r := range Runtimes {
			if r.Name == name {
				return Runtimes[i]
			}
		}
		return Runtimes[0]
}