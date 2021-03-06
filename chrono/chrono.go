package chrono

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	config  = "config.yaml"
	layouts = "layouts"
	pages   = "pages"
)

type Chrono struct {
	sitePath string
	template *template.Template
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
		case layouts:
			err = c.processLayouts(file)
		case pages:
			err = c.processPages(file)
		default:
			err = c.processOther(file)
		}
		if err != nil {
			return errors.Wrap(err, "processing site directory")
		}
	}

	//err = c.outputSite()
	if err != nil {
		return errors.Wrap(err, "outputting site")
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

func (c *Chrono) processLayouts(file os.FileInfo) error {
	layoutsPath := filepath.Join(c.sitePath, file.Name())
	var templateFiles []string
	err := filepath.Walk(layoutsPath, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && strings.HasSuffix(info.Name(), ".html") {
			templateFiles = append(templateFiles, path)
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "walking layouts dir")
	}
	temp, err := template.ParseFiles(templateFiles...)
	if err != nil {
		return errors.Wrap(err, "parsing templates")
	}
	c.template = temp

	return nil
}

func (c *Chrono) processPages(file os.FileInfo) error {
	pagesPath := filepath.Join(c.sitePath, file.Name())
	var pages []string
	err := filepath.Walk(pagesPath, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			pages = append(pages, path)
		}
		return nil
	})

	type Data struct {
		Content string
	}

	for _, page := range pages {
		println(page)
		data := Data{Content: "test"}
		err = c.template.ExecuteTemplate(os.Stdout, "base.html", data)
	}

	return err
}

func (c *Chrono) processOther(file os.FileInfo) error {

	return nil
}

func (c *Chrono) outputSite() error {
	return c.template.ExecuteTemplate(os.Stdout, "base.html", nil)
}
