package filetest

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExists(t *testing.T) {
	if !Exists(tf("readable")) {
		t.Errorf("Exists(readable): want true")
	}
}

func TestFile(t *testing.T) {
	if !IsFile(tf("readable")) {
		t.Errorf("IsFile(readable): want true")
	}
}

func TestDir(t *testing.T) {
	if !IsDir("testfiles/") {
		t.Errorf("IsDir(testfiles)/: want true")
	}
}

func TestZero(t *testing.T) {
	if !IsZero(tf("readable")) {
		t.Errorf("IsZero(readable): want true")
	}
	if IsZero(tf("nonzero")) {
		t.Errorf("IsZero(nonzero): want false")
	}
}

func TestReadable(t *testing.T) {
	stat, _ := os.Stat(tf("writable"))
	t.Logf(stat.Mode().String())
	if !IsReadable(tf("readable")) {
		t.Errorf("IsReadable(readable): want true")
	}
	if IsReadable(tf("writable")) {
		t.Errorf("IsReadable(writable): want false")
	}
	if IsReadable(tf("executable")) {
		t.Errorf("IsReadable(executable): want false")
	}
}

func TestWritable(t *testing.T) {
	if !IsWritable(tf("writable")) {
		t.Errorf("IsWritable(writable): want true")
	}
	if IsWritable(tf("readable")) {
		t.Errorf("IsWritable(readable): want false")
	}
	if IsWritable(tf("executable")) {
		t.Errorf("IsWritable(executable): want false")
	}
}

func TestExecutable(t *testing.T) {
	if !IsExecutable(tf("executable")) {
		t.Errorf("IsExecutable(executable): want true")
	}
	if IsExecutable(tf("readable")) {
		t.Errorf("IsExecutable(readable): want false")
	}
	if IsExecutable(tf("writable")) {
		t.Errorf("IsExecutable(writable): want false")
	}
}

func tf(name string) string {
	return filepath.Join("testfiles", name)
}
