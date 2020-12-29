package chrono

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	config = "config.yaml"
)

type Chrono struct {
	sitePath string
}

func NewChrono(sitePath string) *Chrono {
	return &Chrono{
		sitePath: sitePath,
	}
}

func (c *Chrono) BuildSite() error {
	files, err := ioutil.ReadDir(c.sitePath)
	if err != nil {
		return errors.Wrap(err, "reading sitePath")
	}

	for _, file := range files {
		var err error
		switch file.Name() {
		case config:
			err = c.processConfig(file)
		}
		if err != nil {
			return errors.Wrap(err, "processing site config")
		}
	}

	return nil
}

func (c *Chrono) processConfig(file os.FileInfo) error {
	data, err := ioutil.ReadFile(filepath.Join(c.sitePath, file.Name()))
	if err != nil {
		return errors.Wrap(err, "reading site config yaml")
	}

	var conf SiteConfig

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return errors.Wrap(err, "unmarshalling site config yaml")
	}

	for key, val := range conf {
		fmt.Printf("%v - %v\n", key, val)
	}

	return nil
}
