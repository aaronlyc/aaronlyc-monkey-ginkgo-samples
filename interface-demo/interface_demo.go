package interface_demo

import "fmt"

type Db interface {
	Retrieve(url string) (string, error)
}

type Etcd struct {
}

func (this *Etcd) Retrieve(url string) (string, error) {
	output := fmt.Sprintf("%s, %s!", "Hello", "Etcd")
	return output, nil
}

var Marshal = func(v interface{}) ([]byte, error) {
	return nil, nil
}

type Mysql struct {
}

func (this *Mysql) Retrieve(url string) (string, error) {
	output := fmt.Sprintf("%s, %s!", "Hello", "Mysql")
	return output, nil
}

func NewDb(style string) Db {
	if style == "etcd" {
		return new(Etcd)
	} else {
		return new(Mysql)
	}
}
