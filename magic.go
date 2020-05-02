package magic

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Manager is the main interaction point with the magic package. It coordinates
// all underlying functions and calls to external services
type Manager struct {
	cursor *cursor
	out    io.Writer
}

// NewManager initializes a new manager instance
func NewManager(src io.Reader, srcDec Decoder, out io.Writer) *Manager {
	err := srcDec.FromReader(src)
	if err != nil {
		fmt.Printf("[manager] unable to set reader of srcDecoder: %v\n", err)
		os.Exit(-1)
	}

	return &Manager{
		cursor: newCursor(srcDec),
		out:    out,
	}
}

func NewManagedLevelIterator(level int, exercises [2]int, tester string, testMustPass bool, inSuffix, outSuffix string, dec Decoder, levelExecutor func(m *Manager)) {
	fmt.Printf("[ManagedLevelIterator] Starting with Level %d\n", level)

	doesExists := func(path string) bool {
		_, err := os.Stat(path)
		if err != nil && os.IsNotExist(err) {
			return false
		} else if err != nil {
			fmt.Printf("[ManagedLevelIterator]: error getting os.Stat for solution directory: %v\n", err)
			os.Exit(-1)
		}

		return true
	}

	makeDir := func(path string) {
		if doesExists(path) {
			return
		}

		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Printf("[ManagedLevelIterator]: error making directory: %v\n", err)
			os.Exit(-1)
		}
	}

	solPath := "sol" + strconv.Itoa(level)
	oldIndicator := "old"
	solOldPath := solPath + "/" + oldIndicator

	makeDir(solPath)
	makeDir(solOldPath)

	levelPath := "level" + strconv.Itoa(level)

	getManager := func(exercise string) (*Manager, *bytes.Buffer) {
		b := &bytes.Buffer{}
		m := NewManager(
			FromFile(fmt.Sprintf("%s/%s_%s.%s", levelPath, levelPath, exercise, inSuffix)),
			dec,
			b)

		return m, b
	}

	hash := func(buf []byte) string {
		h := sha256.New()

		_, err := h.Write(buf)
		if err != nil {
			fmt.Printf("[ManagedLevelIterator] Failed to write buffer to sha256 hasher: %v\n", err)
			os.Exit(-1)
		}

		return hex.EncodeToString(h.Sum(nil))[:10]
	}

	hashFile := func(p string) string {
		buf, err := ioutil.ReadFile(p)
		if err != nil {
			fmt.Printf("[ManagedLevelIterator] Failed to read file %q into memory: %v\n", p, err)
			os.Exit(-1)
		}

		return hash(buf)
	}

	saveSolution := func(exercise string, buf []byte, hashStr string) {
		pWild := fmt.Sprintf("%s/%s_%s_*.%s", solPath, levelPath, exercise, outSuffix)
		pHash := strings.ReplaceAll(pWild, "*", hashStr)

		if doesExists(pHash) {
			fmt.Printf("[ManagedLevelIterator] Output hasn't changed since last try of %s_%s\n", levelPath, exercise)
			return
		}

		matches, err := filepath.Glob(pWild)
		if err != nil {
			fmt.Printf("[ManagedLevelIterator] Failed to search for old files at %q: %v\n", pWild, err)
			os.Exit(-1)
		}

		for _, p := range matches {
			last := strings.LastIndex(p, "/")

			err = os.Rename(p, p[:last]+"/"+oldIndicator+"/"+p[last:])
			if err != nil {
				fmt.Printf("[ManagedLevelIterator] Failed to move changed file %q: %v\n", p, err)
				os.Exit(-1)
			}
		}

		f, err := os.OpenFile(pHash, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("[ManagedLevelIterator] Failed to create new solution file %q: %v\n", pHash, err)
			os.Exit(-1)
		}

		_, err = f.Write(buf)
		if err != nil {
			fmt.Printf("[ManagedLevelIterator] Failed to write solution for %q: %v\n", pHash, err)
			os.Exit(-1)
		}
	}

	// test example case
	if testMustPass {
		m, b := getManager(tester)
		levelExecutor(m)

		testerSolFound := b.Bytes()
		testerSolPath := fmt.Sprintf("%s/%s_%s.%s", levelPath, levelPath, tester, outSuffix)

		if hash(testerSolFound) != hashFile(testerSolPath) {
			fmt.Printf("[ManagedLevelIterator] tester case %q failed\n", tester)
			fmt.Println("========================================")
			fmt.Println("============ Wrong Solution ============")
			fmt.Println("========================================")
			fmt.Println(string(testerSolFound))
			os.Exit(-1)
		}

		fmt.Printf("[ManagedLevelIterator] Tester case %q passed successfully. Running exercises now\n", tester)
	} else {
		fmt.Printf("[ManagedLevelIterator] Skipping tester case %q. Running exercises now\n", tester)
	}

	for i := exercises[0]; i <= exercises[1]; i++ {
		e := strconv.Itoa(i)

		m, b := getManager(e)
		levelExecutor(m)

		bBuf := b.Bytes()
		saveSolution(e, bBuf, hash(bBuf))
	}
}
