package golang

import "fmt"

func echo(s fmt.Stringer) {
	print(s.String())
}

func throw(e error) {
	print(e.Error())
}

type iface interface {
	Name() string
	Age() int
}

func needIF(i iface) {
	fmt.Println(i.Name(), i.Age())
}

type myStringer interface {
	fmt.Stringer
	error
}

func strErr(s myStringer) {
	fmt.Println(s.Error(), s.String())
}
