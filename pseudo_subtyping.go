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

type Citizen struct {
    Country string
    Person
}

func (p *Person) say_hello() {
    fmt.Println("hello, I'm", p.Name)
}

func (p *Person) say_location() {
    fmt.Println("I'm from",p.Address.Province, p.Address.City)
}

func (c *Citizen) Nationality() {
    fmt.Println(c.Name, "is a citizen of", c.Country)
}

func (p *Citizen) say_hello() {
    fmt.Println("hello, I'm Citizen", p.Name)
}

func main() {
    p := Citizen {}
    p.Name = "zyy"
    p.Address.Province = "Sichuan"
    p.Address.City = "Chengdu"
    p.Country = "China"
    p.say_hello()
    p.Person.say_hello()
    p.say_location()
    p.Nationality()
}
