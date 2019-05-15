package filetest

import (
	"os"
)

// Exists reports whether the path exists.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// IsFile reports whether the path exists and is a file.
func IsFile(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && !stat.IsDir()
}

// IsDir reports whether the path exists and is a directory.
func IsDir(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && stat.IsDir()
}

func IsSymlink(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && stat.Mode()&os.ModeSymlink != 0
}

func IsZero(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && stat.Size() == 0
}

func IsReadable(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && stat.Mode()&(1<<(9-1)) != 0
}

func IsWritable(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && stat.Mode()&(1<<(9-2)) != 0
}

func IsExecutable(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && stat.Mode()&(1<<(9-3)) != 0
}
