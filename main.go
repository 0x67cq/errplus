package main

import (
	"errors"
	"fmt"

	"github.com/psychix/errplus/errplus"
)

func main() {
	fmt.Println("hello world")
	err := bllfunc()
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func bllfunc() error {
	err := dotfunc(1, 2, 3)
	if err != nil {
		return errplus.Wrap(err, "name", 1, "stck", 2, "call", 3)
	}
	return nil
}

func dotfunc(...int) error {
	return errors.New("Ooooop!....")
}
