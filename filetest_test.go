package filetest

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExists(t *testing.T) {
	setPerm()
	if !Exists(tf("readable")) {
		t.Errorf("Exists(readable): want true")
	}
}

func TestFile(t *testing.T) {
	setPerm()
	if !IsFile(tf("readable")) {
		t.Errorf("IsFile(readable): want true")
	}
}

func TestDir(t *testing.T) {
	setPerm()
	if !IsDir("testfiles/") {
		t.Errorf("IsDir(testfiles)/: want true")
	}
}

func TestZero(t *testing.T) {
	setPerm()
	if !IsZero(tf("readable")) {
		t.Errorf("IsZero(readable): want true")
	}
	if IsZero(tf("nonzero")) {
		t.Errorf("IsZero(nonzero): want false")
	}
}

func TestReadable(t *testing.T) {
	setPerm()
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
	setPerm()
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
	setPerm()
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

func setPerm() {
	os.Chmod(tf("readable"), 0400)
	os.Chmod(tf("writable"), 0200)
	os.Chmod(tf("executable"), 0100)
}

func tf(name string) string {
	return filepath.Join("testfiles", name)
}
