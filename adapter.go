package multicsvadapter

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

type Adapter struct {
	dirPath string
}

func NewAdapter(dirPath string) *Adapter {
	return &Adapter{dirPath}
}

func (a *Adapter) LoadPolicy(model model.Model) error {
	countLoadFile := 0

	if err := filepath.Walk(a.dirPath, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(strings.ToLower(f.Name()), ".csv") && !f.IsDir() {
			if err := a.loadPolicyFile(path, model, persist.LoadPolicyLine); err != nil {
				return err
			}

			countLoadFile++
		}

		return nil
	}); err != nil {
		return err
	}

	if countLoadFile == 0 {
		return errors.New(errInvalidDirPath)
	}

	return nil
}

func (a *Adapter) loadPolicyFile(filepath string, model model.Model, handler func(string, model.Model)) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		handler(line, model)
	}

	return scanner.Err()
}

func (a *Adapter) SavePolicy(model model.Model) error {
	return errors.New(errNotImplemented)
}

func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	return errors.New(errNotImplemented)
}

func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return errors.New(errNotImplemented)
}

func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return errors.New(errNotImplemented)
}
