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

func stimulate(p *Person) {
    // after stimulation, he should speak
    p.say_hello()
}

func main() {
    p := Person{Name: "zyy"}
    c := Citizen{Person: Person{Name: "yzz"},Country: "China"}
    stimulate(&p)
    stimulate(&c)
}
