package main

import (
    "io/ioutil"
    "fmt"
    "github.com/kaml/gocode/input"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    fmt.Println("Hello world! My lucky number is", input.Mul2(privateHelperFunc()))
    dat, err := ioutil.ReadFile("tmp/dat")
    check(err)
    fmt.Println(string(dat))

    //Initalise the Init function with value of A,B
    //e := ellipse.Init{
    //9, 2,
    //}
    //this will give answer as 0.9749960430435691
    //fmt.Println(e.GetEccentricity())
}
