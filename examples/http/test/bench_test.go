package test

import (
	"bytes"
	"encoding/gob"
	"github.com/jinzhu/copier"
	"testing"
)

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

type Account struct {
	AccountId int
	Account string
}

type User struct {
	AccountId int
	Account string
	Name string
}

func BenchmarkCopy1(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		acc := Account{
			AccountId: 10,
			Account:   "jirenyou",
		}

		usr := User{}

		if err := deepCopy(&usr, acc); err != nil {
			b.Error(err)
			return
		}

		//b.Logf("--- %v", usr)
	}
}

func BenchmarkCopy2(b *testing.B) {
	for i :=0; i < b.N; i++ {
		acc := Account{
			AccountId: 10,
			Account:   "jirenyou",
		}

		usr := User{}

		if err := copier.Copy(&usr, &acc); err != nil {
			b.Error(err)
			return
		}
	}
}
