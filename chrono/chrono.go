package chrono

import (
	"fmt"
	"io/ioutil"
)

func BuildSite(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	return nil
}
