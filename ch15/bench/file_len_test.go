package filelen

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

const expectedFileLen = 65204

var dataFile string

func TestMain(m *testing.M) {
	os.Exit(runTests(m))
}

// runTests handles setup and cleanup for all tests and benchmarks
// in this package.
func runTests(m *testing.M) (exitCode int) {
	dir, err := os.MkdirTemp("", "filelen-test-*")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	defer func() {
		if err := os.RemoveAll(dir); err != nil {
			fmt.Fprintln(os.Stderr, err)
			exitCode = 1
		}
	}()

	dataFile = filepath.Join(dir, "data.txt")

	if err := makeData(dataFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return m.Run()
}

// makeData generates a deterministic test file rather than storing
// a large fixture in source control.
//
// A local pseudorandom number generator with a fixed seed ensures
// that the same data is generated on every run.
//
// The resulting file contains 10,000 newline-terminated words
// and is exactly 65,204 bytes long.
func makeData(path string) error {
	rng := rand.New(rand.NewSource(1))

	var data bytes.Buffer

	for range 10_000 {
		word := makeWord(rng, rng.Intn(10)+1)
		data.Write(word)
	}

	return os.WriteFile(path, data.Bytes(), 0600)
}

// makeWord generates a word of length l consisting of lowercase
// ASCII letters, followed by a newline character.
func makeWord(rng *rand.Rand, l int) []byte {
	out := make([]byte, l+1)

	for i := range l {
		out[i] = 'a' + byte(rng.Intn(26))
	}

	out[l] = '\n'
	return out
}

func TestFileLen(t *testing.T) {
	result, err := FileLen(dataFile, 1)
	if err != nil {
		t.Fatal(err)
	}

	if result != expectedFileLen {
		t.Errorf("expected %d, got %d", expectedFileLen, result)
	}
}

// BenchmarkFileLen measures the performance of FileLen using
// different buffer sizes.
func BenchmarkFileLen(b *testing.B) {
	for _, size := range []int{1, 10, 100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("BufferSize-%d", size), func(b *testing.B) {
			// we don't need to use the previous i < b.N condition
			// Thus it's cleaner and more concise.
			for b.Loop() {
				_, err := FileLen(dataFile, size)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
