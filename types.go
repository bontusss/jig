package jig

type Jig struct {
	AppName string
	Version string
	Debug   bool
}

// initPaths is used when initializing the application. It holds the root
// path for the application, and a slice of strings with the names of
// folders that the application expects to find.
type initPaths struct {
	rootPath    string
	folderNames []string
}
