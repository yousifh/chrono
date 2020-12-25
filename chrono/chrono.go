package chrono

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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
		return err
	}

	for _, file := range files {
		var err error
		switch file.Name() {
		case config:
			err = c.processConfig(file)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Chrono) processConfig(file os.FileInfo) error {
	data, err := ioutil.ReadFile(filepath.Join(c.sitePath, file.Name()))
	if err != nil {
		return err
	}

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(data, m)
	if err != nil {
		return err
	}

	for key, val := range m {
		fmt.Printf("%v - %v\n", key, val)
	}

	return nil
}
