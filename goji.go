package goji

const (
	VERSION = "1.0.0"
)

var projectFolders = []string{
	"models",
	"templates",
	"controllers",
	"config",
	"migrations",
	"tmp",
	"public",
	"logs",
}

// New reads the .env file, creates our application config, populates a goji app with settings
// based on .env values, and creates necessary folders and files if they don't exist
func (g *Goji) New(rootPath string) error {
	g.Version = VERSION
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: projectFolders,
	}
	err := g.Init(pathConfig)
	if err != nil {
		return err
	}

	err = g.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	return nil
}

// Init creates necessary folders for a Goji application
func (g *Goji) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		err := g.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}
