package main

import "fmt"

type Person struct {
    Name string
    Address Address
}

type Address struct {
    Province string
    City string
}

func (p *Person) say_hello() {
    fmt.Println("hello, I'm", p.Name)
}

func (p *Person) say_location() {
    fmt.Println("I'm from",p.Address.Province, p.Address.City)
}

func main() {
    p := Person {}
    p.Name = "zyy"
    p.Address.Province = "Sichuan"
    p.Address.City = "Chengdu"
    p.say_hello()
    p.say_location()
}
