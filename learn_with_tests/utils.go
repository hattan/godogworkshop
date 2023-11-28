package test

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/hattan/learngo/pkg/wcase"
	"golang.org/x/exp/constraints"
)

func writeLine(value any) {
	fmt.Println(value)
}

func runFunctionCaptureStdOut(f func()) string {
	// redirect stdout to a pipe
	oldStdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// invoke function
	f()

	// read output & restore stdout
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdOut

	return string(out)
}

func WordCount(s string, c wcase.Case) map[string]int {
	m := make(map[string]int)
	tokens := strings.Fields(s)
	for _, v := range tokens {
		if c == wcase.INSENSITIVE {
			v = strings.ToLower(v)
		}
		m[v] = (m[v] + 1)
	}

	return m
}

func add(x, y int) int {
	return x + y
}

func addE(x, y int) (int, error) {
	if x < 0 || y < 0 {
		return 0, fmt.Errorf("x and y must be positive")
	}

	return x + y, nil
}

var (
	globalValue int
	globalPtr   *int
)

func sumValue(x, y int) int {
	z := x + y
	return z
}
func sumPtr(x, y int) *int {
	z := x + y
	return &z
}

type Number interface {
	int | int64 | float64
}

func sumGeneric[T Number](a T, b T) T {
	return a + b
}

func sumGenericConstraints[T constraints.Ordered](a T, b T) T {
	return a + b
}

func CreateFileHandle(name string) *os.File {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("error opening file: %v", err))
	}
	return f
}
