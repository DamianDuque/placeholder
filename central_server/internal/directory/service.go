package directory

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidFormat = errors.New("directory: invalid file extension")
)

type ServiceDirectory interface {
	Query(filename string) (string, error)
	SendIndex(indexInfo Index) error
	GetIndexTable() map[string]string
}

type ServiceDefaultDir struct {
	repository DirRepository
}

func NewServiceClient(repo DirRepository) *ServiceDefaultDir {
	return &ServiceDefaultDir{repository: repo}

}

// Query searches for a file in the index table
func (s ServiceDefaultDir) Query(filename string) (string, error) {
	return s.repository.SearchFile(filename)
}

// SendIndex saves the index in the repository
func (s ServiceDefaultDir) SendIndex(indexInfo Index) error {
	for _, file := range indexInfo.Files {
		err := validateFileExtension(file)
		if err != nil {
			return err
		}
	}
	err := s.repository.SaveIndex(indexInfo)
	if err != nil {
		return err
	}

	return nil
}

// GetIndexTable returns the index table
func (s ServiceDefaultDir) GetIndexTable() map[string]string {
	return s.repository.GetIndexTable()

}

// validateFileExtension validates that the file extension is allowed and correct
func validateFileExtension(filename string) error {
	allowedExtensions := []string{"jpg", "png", "doc", "docx", "html", "txt"}
	pattern := `(\w+)\.(` + strings.Join(allowedExtensions, "|") + `)$`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	if !re.MatchString(filename) {
		return ErrInvalidFormat
	}
	return nil
}
