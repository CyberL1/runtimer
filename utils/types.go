package utils

type LocalConfigType struct {
	Primary string
	Runtimes []RuntimeType
}

type RuntimeType struct {
	Name string
	Runtime string
	Version string
	Ext string
	Bin string
	Os map[string]OsType
	Arch map[string]string
}

type OsType struct {
	Name string
	Ext string
	Bin string
}