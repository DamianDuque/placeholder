package directory

import (
	"errors"
)

var (
	ErrNotFound = errors.New("directory: index file not found")
)

type DirRepository interface {
	SaveIndex(index Index) error
	GetIndexTable() map[string]string
	SearchFile(filename string) (string, error)
}
type defaultMapRepo struct {
	indexTable *map[string]string
}

func NewDefaultRepo(newIndexTable *map[string]string) DirRepository {
	return defaultMapRepo{indexTable: newIndexTable}
}
func (d defaultMapRepo) SaveIndex(index Index) error {
	table := *d.indexTable
	for _, file := range index.Files {
		table[file] = index.Username
	}
	return nil
}

func (d defaultMapRepo) GetIndexTable() map[string]string {
	table := *d.indexTable
	return table
}

func (d defaultMapRepo) SearchFile(filename string) (string, error) {
	table := *d.indexTable
	if _, ok := table[filename]; ok {
		user := table[filename]
		return user, nil
	}
	return "", ErrNotFound
}
