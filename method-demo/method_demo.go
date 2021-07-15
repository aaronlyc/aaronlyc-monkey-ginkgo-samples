package method_demo

import (
	"errors"
	"fmt"
)

var (
	ErrElemExsit    = errors.New("elem already exist")
	ErrElemNotExsit = errors.New("elem not exist")
)

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}

func (this *Slice) Add(elem int) error {
	for _, v := range *this {
		if v == elem {
			fmt.Printf("Slice: Add elem: %v already exist\n", elem)
			return ErrElemExsit
		}
	}
	*this = append(*this, elem)
	fmt.Printf("Slice: Add elem: %v succ\n", elem)
	return nil
}

func (this *Slice) Remove(elem int) error {
	found := false
	for i, v := range *this {
		if v == elem {
			if i == len(*this)-1 {
				*this = (*this)[:i]

			} else {
				*this = append((*this)[:i], (*this)[i+1:]...)
			}
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Slice: Remove elem: %v not exist\n", elem)
		return ErrElemNotExsit
	}
	fmt.Printf("Slice: Remove elem: %v succ\n", elem)
	return nil
}
