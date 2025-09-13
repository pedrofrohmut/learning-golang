package main

import(
    "fmt"
)

type Animal interface {
    AngrySound()
    HappySound()
}

type Cat string

func (this Cat) Attack() {
    fmt.Println("Cat attacks its prey")
}

func (this Cat) Name() string {
    return string(this)
}

func (this Cat) AngrySound() {
    fmt.Println("Hissss")
}

func (this Cat) HappySound() {
    fmt.Println("Purrrrr")
}

func main() {
    var kitty Animal
    kitty = Cat("Kitty")
    kitty.AngrySound()
    kitty.HappySound()
}
