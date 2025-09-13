package main

import (
    "fmt"
)

type Customer struct {
    name string
    address string
    balance float64
}

func printCustomer(customer Customer) {
    fmt.Printf("Name: %s, Address: %s, Balance: %f\n", customer.name, customer.address, customer.balance)
}

func newCustomerAddresss(customer *Customer, address string) {
    customer.address = address
}

type Rectangle struct {
    length, height float64
}

func (this Rectangle) Area() float64 {
    return this.length * this.height
}

type Contact struct {
    fName string
    lName string
    phone string
}

// Golang does not have inheritance but it supports composition
type Business struct {
    name string
    address string
    contact Contact
}

func (this Business) info() string {
    return fmt.Sprintf("Contact at %s is %s %s", this.name, this.contact.fName, this.contact.lName)
}

func Fn018() {
    var c1 Customer
    c1.name = "John Doe"
    c1.address = "5, main st"
    c1.balance = 123.45
    printCustomer(c1)

    newCustomerAddresss(&c1, "123, second st")

    // You can init with values
    var c2 = Customer { "Jane Smith", "42, third st", 666.66 }

    // When you do this you have to make in the right order (here name and address are inverted)
    var c3 = Customer { "18 main st", "Jane Smith", 666.66 }

    // You can provide the names when init and the order wont matter
    var c4 = Customer { name: "Walter White", address: "420 forth st", balance: 1234.56 }

    printCustomer(c1)
    printCustomer(c2)
    printCustomer(c3)
    printCustomer(c4)

    var rect = Rectangle { 10.0, 15.0 }
    fmt.Println("Rectangle area:", rect.Area())

    var contact1 = Contact { fName: "James", lName: "West", phone: "555-1212" }
    var b1 = Business { name: "ACME", address: "123, Main st", contact: contact1 }
    fmt.Println(b1.info())
}
