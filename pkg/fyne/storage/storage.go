// Storage is used to manage file storage inside an application sandbox.
package storage

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"fyne.io/fyne/v2"
)

type Saver[T any] interface {
	ToSave() string
	FromSave(string) T
}

func SaveQuery[T Saver[T]](filepath string, value []T) error {
	var toSave string
	for _, v := range value {
		toSave += v.ToSave() + "\n"
	}

	// does it work when the file doesn't exist?
	save, err := fyne.CurrentApp().Storage().Save(filepath)
	if err != nil {
		return errors.New("failed to save in storage")
	}
	save.Write([]byte(toSave))
	return nil
}

func LoadQuery[T Saver[T]](filepath string) ([]T, error) {
	ret := make([]T, 0)

	load, err := fyne.CurrentApp().Storage().Open(filepath)
	if err != nil {
		return ret, errors.New("failed to open file")
	}
	defer load.Close()

	file, err := io.ReadAll(load)
	if err != nil {
		return ret, errors.New("failed to read file")
	}

	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		var v T
		ret = append(ret, v.FromSave(line))
	}

	fmt.Printf("ret: %v\n", ret)
	return ret, nil
}
