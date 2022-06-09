package main

type Pessoa struct {
	name string
	age  uint
}

func (p Pessoa) sayHello() string {
	return "Hello my name is " + p.name + "!"
}
