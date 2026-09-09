package skoleruten

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Kommune struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func readKommuneSourceFile(name string) ([]byte, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("unable to determine current source path")
	}
	path := filepath.Join(filepath.Dir(filename), name)
	return os.ReadFile(path)
}

func LoadKommuner() ([]Kommune, error) {
	data, err := readKommuneSourceFile("kommuner.json")
	if err != nil {
		return nil, err
	}

	var kommuner []Kommune
	if err := json.Unmarshal(data, &kommuner); err != nil {
		return nil, err
	}
	return kommuner, nil
}

func DefaultKommuner() ([]Kommune, error) {
	return LoadKommuner()
}

func FindKommuneByID(id string) (*Kommune, bool) {
	kommuner, err := LoadKommuner()
	if err != nil {
		return nil, false
	}
	for _, k := range kommuner {
		if k.ID == id {
			return &k, true
		}
	}
	return nil, false
}
