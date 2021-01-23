package test

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	t.Logf("test: %v", time.Now())
	t.Logf("test: %v", time.Now().Add(3 * time.Second))
}
