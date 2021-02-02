package test

import (
	"fmt"
	"github.com/jinzhu/copier"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	t.Logf("test: %v", time.Now())
	t.Logf("test: %v", time.Now().Add(3 * time.Second))
}

func TestStringArray(t *testing.T)  {
	var str []string
	str = append(str, "aa")
	fmt.Println(str)
	fmt.Println(str == nil)
}

func TestPath(t *testing.T)  {
	file, _ := exec.LookPath(os.Args[0])
	t.Logf("file: %v", file)

	path, _ := filepath.Abs(file)
	t.Logf("path: %v", path)

	index := strings.LastIndex(path, string(os.PathSeparator))
	t.Logf("index: %v", index)

	path = path[:index]
	t.Logf("path: %v", path)

}

func TestString(t *testing.T)  {
	str := "jirenyou"
	index := strings.LastIndex(str, "rena")
	t.Logf("index: %v", index)

	t.Logf("aa: %v",string(os.PathSeparator))

}

func TestCopy(t *testing.T)  {
	acc := Account{
		AccountId: 10,
		Account:   "jirenyou",
	}

	b := new(User)

	if err := deepCopy(b, acc); err != nil {
		t.Error(err)
		return
	}

	t.Logf("----%v", b)
}

func TestCopy2(t *testing.T)  {
	acc := Account{
		AccountId: 10,
		Account:   "jirenyou",
	}

	use := User{}

	if err := copier.Copy(&use, &acc); err != nil {
		t.Errorf("error: %v", err)
		return
	}

	t.Logf("----- %v", use)
}