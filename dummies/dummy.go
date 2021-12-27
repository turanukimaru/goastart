package dummies

import (
	"fmt"
	"github.com/turanukimaru/goastart/dummydb"
)

type Dummy struct {
}

func (d *Dummy) Hello() (err error) {
	fmt.Println("Hello Dummies !")
	err = dummydb.DbAccess()
	return
}

func (d *Dummy) Allow() (err error) {
	fmt.Println("Hello Dummies !")
	err = dummydb.DbAccess()
	return
}
