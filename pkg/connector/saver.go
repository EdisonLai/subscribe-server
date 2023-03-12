package connector

import (
	"bufio"
	"fmt"
	"os"
)

type PersistenceStore interface {
	Save(data string) error
	Load() ([]string, error)
}

type FilePersistenceStore struct {
	filepath string
}

func NewFilePersistenceStore(filepath string) *FilePersistenceStore {
	return &FilePersistenceStore{filepath: filepath}
}

func (f *FilePersistenceStore) Save(data string) error {
	// 打开文件，如果不存在则创建
	file, err := os.OpenFile(f.filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open file %s: %s\n", f.filepath, err)
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(data))
	if err != nil {
		return err
	}

	return nil
}

func (f *FilePersistenceStore) Load() ([]string, error) {
	file, err := os.Open(f.filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
