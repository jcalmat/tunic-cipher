// Storage is used to manage file storage inside an application sandbox.
package storage

import (
	"errors"
	"io"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

type Saver[T any] interface {
	ToSave() string
	FromSave(string) T
}

func Save[T Saver[T]](filepath string, value []T) error {
	var toSave string
	for _, v := range value {
		toSave += v.ToSave() + "\n"
	}

	_, err := fyne.CurrentApp().Storage().Create(filepath)
	if err != nil {
		// if the file already exists, proceed
		if !errors.Is(err, storage.ErrAlreadyExists) {
			return errors.New("failed to create file")
		}
	}

	// does it work when the file doesn't exist?
	save, err := fyne.CurrentApp().Storage().Save(filepath)
	if err != nil {
		return errors.New("failed to save in storage")
	}
	_, err = save.Write([]byte(toSave))
	return err
}

func Load[T Saver[T]](filepath string) ([]T, error) {
	ret := make([]T, 0)

	load, err := fyne.CurrentApp().Storage().Open(filepath)
	if err != nil {
		return ret, ErrNotFound
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

	return ret, nil
}
