package test

import (
	"github.com/AlkBur/ServerOneScript"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

const root = "."

func TestRun(t *testing.T) {
	srv := ServerOneScript.NewWorker()

	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".os" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			t.Error(err)
		}

		str := string(b)
		srv.RunScript(str)
	}
	srv.Stop()
}
