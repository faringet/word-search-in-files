package dir

import (
	"go.uber.org/zap"
	"io/fs"
)

type Files struct {
	FS     fs.FS
	logger *zap.Logger
}

func NewFilesFS(fsys fs.FS, logger *zap.Logger) *Files {
	return &Files{FS: fsys, logger: logger}
}

func (f *Files) List(dir string) ([]string, error) {
	f.logger.Info("Listing files in directory", zap.String("directory", dir))

	var fileNames []string
	err := fs.WalkDir(f.FS, dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			f.logger.Error("Error walking directory", zap.Error(err))
			return err
		}
		if !d.IsDir() {
			fileNames = append(fileNames, path)
		}
		return nil
	})
	if err != nil {
		f.logger.Error("Error listing files", zap.Error(err))
		return nil, err
	}

	f.logger.Info("Finished listing files", zap.Int("count", len(fileNames)))
	return fileNames, nil
}
