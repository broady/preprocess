package preprocess

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGolden(t *testing.T) {
	files, err := filepath.Glob("testdata/*.in")
	if err != nil {
		t.Fatal(err)
	}

	for _, fn := range files {
		t.Run(fn, func(t *testing.T) {
			f, err := os.Open(fn)
			if err != nil {
				t.Fatal(err)
			}

			out, err := Process(f, nil, "//#")
			if err != nil {
				t.Fatal(err)
			}

			goldenFn := "testdata/golden/" + filepath.Base(fn)
			golden, err := ioutil.ReadFile(goldenFn)
			if os.IsNotExist(err) {
				t.Log("Writing golden file")
				if err := ioutil.WriteFile(goldenFn, out, 0700); err != nil {
					t.Fatal(err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(out, golden) {
				os.MkdirAll("testdata/fail", 0700)
				failFn := "testdata/fail/" + filepath.Base(fn)
				t.Logf("Doesn't match golden file. Writing %s", failFn)
				t.Logf("\tdiff -y testdata/*/%s", filepath.Base(fn))
				if err := ioutil.WriteFile(failFn, out, 0700); err != nil {
					t.Fatal(err)
				}
				t.Fail()
			}
		})
	}
}
