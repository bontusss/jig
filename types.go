package jig

import "github.com/bontusss/jig/log"

type Jig struct {
	AppName string
	Version string
	// The Debug flag indicates whether the application should run in debug mode or not.
	Debug   bool
	Logger *log.Logger
	Handlers []string
}

// initPaths is used when initializing the application. It holds the root
// path for the application, and a slice of strings with the names of
// folders that the application expects to find.
type initPaths struct {
	rootPath    string
	folderNames []string
}
