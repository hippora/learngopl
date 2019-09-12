package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type limitReader struct {
	r io.Reader
	n int64
}

func (l *limitReader) Read(p []byte) (n int, err error) {
	if l.n <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.n {
		p = p[0:l.n]
	}
	n, err = l.r.Read(p)
	l.n -= int64(n)
	return
}

// LimitReader returns a Reader that reads from r
// but stops with EOF after n bytes.
func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r, n}
}

func main() {
	s := strings.NewReader("hello,World!")
	lr1 := LimitReader(s, 9)
	lr2 := LimitReader(lr1, 2)
	n, err := ioutil.ReadAll(lr2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(n))
	n, err = ioutil.ReadAll(lr1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(n))
	fmt.Println(s)
}
