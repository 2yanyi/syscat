package internal

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/json-iterator/go"
	"github.com/pkg/errors"
	"io"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"unsafe"
)

func SizeFormat(bytes float64) (_ string) {
	if bytes >= _GB {
		return fmt.Sprintf("%.1fG", bytes/_GB)
	} else if bytes >= _MB {
		return fmt.Sprintf("%.1fM", bytes/_MB)
	} else if bytes >= _KB {
		return fmt.Sprintf("%.1fK", bytes/_KB)
	}
	return
}

func Stderr(err string) {
	if err != "" {
		_, _ = fmt.Fprintf(os.Stderr, "error: %s", err)
	}
}

func String(fp string) string {
	a := Bytes(fp)
	if len(a) > 32 {
		return *(*string)(unsafe.Pointer(&a))
	}
	return string(a)
}

func Bytes(fp string) []byte {
	data, err := os.ReadFile(fp)
	if err != nil {
		Stderr(err.Error())
	}
	return data
}

func Wcl(fp string) int {
	src, err := os.Open(fp)
	if err != nil {
		return 0
	}
	defer src.Close()

	buf := make([]byte, 1024*32)
	sep := []byte{'\n'}
	wcl := 0
	n := 0
	for {
		n, err = src.Read(buf)
		wcl += bytes.Count(buf[:n], sep)
		switch {
		case err == io.EOF:
			return wcl
		case err != nil:
			return wcl
		}
	}
}

func MD5sumChunked(fp string) (os.FileInfo, string, error) {
	src, err := os.Open(fp)
	if err != nil {
		return nil, "", errors.New(err.Error())
	}
	defer src.Close()

	info, _ := src.Stat()
	if info.IsDir() {
		return info, "", errors.New(fmt.Sprintf("%s is a directory", fp))
	}

	// Chunked calculations
	size := info.Size()
	blocks := uint64(math.Ceil(float64(size) / float64(_MB)))
	hash := md5.New()
	for i := uint64(0); i < blocks; i++ {
		blockSize := int(math.Min(_MB, float64(size-int64(i*_MB))))
		buf := make([]byte, blockSize)
		if _, err = src.Read(buf); err != nil {
			return info, "", errors.New(err.Error())
		}
		if _, err = io.WriteString(hash, string(buf)); err != nil {
			return info, "", errors.New(err.Error())
		}
	}
	sum := hex.EncodeToString(hash.Sum(nil))

	return info, sum, nil
}

func FileExist(fp string) bool {
	_, err := os.Lstat(fp)
	return !os.IsNotExist(err)
}

func RandomChars(n int) []byte {
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return buf
}

func Commandline(dir string, args []string) (_ string) {
	if len(args) == 0 {
		return
	}
	buffer := &bytes.Buffer{}
	stdWrite := io.MultiWriter(buffer)
	cmd := exec.Command(args[0])
	cmd.Stdout, cmd.Stderr = stdWrite, stdWrite
	cmd.Args = args
	if dir != "" {
		cmd.Dir = dir
	}
	if err := cmd.Run(); err != nil {
		Stderr(err.Error())
	}
	return strings.TrimSpace(buffer.String())
}

func BashC(dir, sh string) string {
	buffer := &bytes.Buffer{}
	stdWrite := io.MultiWriter(buffer)
	cmd := exec.Command("bash", "-c", sh)
	cmd.Stdout, cmd.Stderr = stdWrite, stdWrite
	if dir != "" {
		cmd.Dir = dir
	}
	if err := cmd.Run(); err != nil {
		Stderr(err.Error())
	}
	return strings.TrimSpace(buffer.String())
}

func JsonFormat(a interface{}) []byte {
	data, err := jsoniter.ConfigFastest.MarshalIndent(a, "", "  ")
	if err != nil {
		Stderr(err.Error())
	}
	return data
}

const _KB = 1024
const _MB = _KB * _KB
const _GB = _MB * _KB

const letterBytes = "0123456789._=abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
