package searcher

import (
	"go.uber.org/zap"
	"io/ioutil"
	"strings"
	"sync"
	"word-search-in-files/pkg/dir"
)

type Searcher struct {
	Dir    *dir.Files
	logger *zap.Logger
}

func NewSearcher(dir *dir.Files, logger *zap.Logger) *Searcher {
	return &Searcher{Dir: dir, logger: logger}
}

func (s *Searcher) Search(word string) ([]string, error) {
	s.logger.Info("Starting search", zap.String("word", word))

	files, err := s.Dir.List(".")
	if err != nil {
		return nil, err
	}

	var (
		resultFiles []string
		mutex       sync.Mutex
	)

	var wg sync.WaitGroup
	wg.Add(len(files))

	for _, file := range files {
		go func(filename string) {
			defer wg.Done()

			content, err := ioutil.ReadFile(filename)
			if err != nil {
				s.logger.Error("Error reading file", zap.String("file", filename), zap.Error(err))
				return
			}

			if strings.Contains(string(content), word) {
				mutex.Lock()
				resultFiles = append(resultFiles, filename)
				mutex.Unlock()
			}
		}(file)
	}

	wg.Wait()

	s.logger.Info("Search completed", zap.Int("files_found", len(resultFiles)))
	if len(resultFiles) == 0 {
		return nil, nil
	}

	return resultFiles, nil
}
