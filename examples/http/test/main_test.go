package test

import (
	"fmt"
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