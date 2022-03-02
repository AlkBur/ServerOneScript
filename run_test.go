package ServerOneScript

import (
	"github.com/AlkBur/ServerOneScript/runes"
	"github.com/AlkBur/ServerOneScript/vm"
	"os"
	"path/filepath"
	"testing"
)

const root = "."

func TestRun(t *testing.T) {
	//var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".os" {
			s, err := runes.NeStreamFromFile(path)
			if err != nil {
				t.Error(err)
			}
			program, err := Compile(s)
			if err != nil {
				t.Error(err)
			}
			output, err := vm.Run(program)
			if err != nil {
				t.Error(err)
			}
			t.Log(output)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
