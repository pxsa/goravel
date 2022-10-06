package goravel

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const version = "1.0.0"

type Goravel struct {
	AppName string
	Debug bool
	Version string
}

func (g *Goravel) New(rootPath string) error {
	pathConfig := initPath {
		rootPath: rootPath,
		folderNames: []string {"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}
	err := g.Init(pathConfig)
	if err != nil {
		return err
	}

	err = g.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	// read .env
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	return nil
}

func (g *Goravel) Init(p initPath) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		// create folder if it doesn't exist
		err := g.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Goravel) checkDotEnv(path string) error{
	err := g.CreateFileIfNotExist(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil
}