package utils

type LocalConfig struct {
	Primary string
	Runtimes []Runtime
}

type Runtime struct {
	Name string
	Runtime string
	Version string
	Ext string
	Bin string
	Os map[string]Os
	Arch map[string]string
}

type Os struct {
	Name string
	Ext string
	Bin string
}