package internal

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

func Default(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

func Set(data ...string) map[string]struct{} {
	var set = make(map[string]struct{})
	for i := range data {
		set[data[i]] = struct{}{}
	}
	return set
}

// IsWritable 目录是否可写
func IsWritable(path string) (isWritable bool, err error) {
	if err = syscall.Access(path, syscall.O_RDWR); err == nil {
		isWritable = true
	}

	return
}

// FileLineNum 获取文件行数
func FileLineNum(path string) (num int, err error) {
	fi, err := os.Open(path)
	if err != nil {
		return 0, err
	}

	br := bufio.NewReader(fi)

	for {
		_, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}

		num++
	}

	return num, nil
}

func FileCopy(dest, source string) error {
	df, err := os.Create(dest)
	if err != nil {
		return err
	}

	f, err := os.Open(source)
	if err != nil {
		return err
	}

	_, err = io.Copy(df, f)
	f.Close()

	return err
}

// 逐行扫描文件,回调返回当前内容及当前行号
func FileScan(path string, f func(b []byte, line int) bool, maxLineLen ...int) error {
	fi, err := os.Open(path)
	if err != nil {
		return err
	}

	ms := 40960 // 40K
	if len(maxLineLen) > 0 {
		ms = maxLineLen[0]
	}

	br := bufio.NewReader(fi)
	line := 0
	eb := string([]byte{0}) // empty bytes
	tb := make([]byte, ms)

	for {
		b, isPrefix, err := br.ReadLine()
		tb = append(tb, b...)

		if isPrefix {
			continue
		}

		if err == io.EOF {
			break
		}

		line++

		tb = bytes.Trim(tb, eb)

		r := f(tb, line)

		tb = []byte{}

		if !r {
			break
		}
	}

	return nil
}

func ScanDir(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}

	res := []string{}

	for _, fi := range files {
		if !fi.IsDir() {
			res = append(res, fi.Name())
		}
	}

	return res, nil
}

// pathinfo()
// -1: all; 1: dirname; 2: basename; 4: extension; 8: filename
// Usage:
// Pathinfo("/home/go/path/src/php2go/php2go.go", 1|2|4|8)
func Pathinfo(path string, options int) map[string]string {
	if options == -1 {
		options = 1 | 2 | 4 | 8
	}

	info := make(map[string]string)
	if (options & 1) == 1 {
		info["dirname"] = filepath.Dir(path)
	}

	if (options & 2) == 2 {
		info["basename"] = filepath.Base(path)
	}

	if ((options & 4) == 4) || ((options & 8) == 8) {
		basename := ""

		if (options & 2) == 2 {
			if tmp, ok := info["basename"]; ok {
				basename = tmp
			}
		} else {
			basename = filepath.Base(path)
		}

		p := strings.LastIndex(basename, ".")
		filename, extension := "", ""

		switch p {
		case 0:
			extension = basename[p+1:]
		case -1:
			filename = basename
		default:
			filename, extension = basename[:p], basename[p+1:]
		}

		if (options & 4) == 4 {
			info["extension"] = extension
		}

		if (options & 8) == 8 {
			info["filename"] = filename
		}
	}

	return info
}

// file_put_contents()
func FilePutContents(filename, data string, mode os.FileMode) error {
	return ioutil.WriteFile(filename, []byte(data), mode)
}

// file_get_contents()
func FileGetContents(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	return string(data), err
}

// unlink()
func Unlink(filename string) error {
	return os.Remove(filename)
}
