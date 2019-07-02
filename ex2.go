package main

import (
    
    "fmt"
    
)

func main() {
type student := struct {

    name   string

    code   int

    gpa    float64

    status bool

}{

    name:   "Panumars",

    code:   1002,

    gpa:    3.59,

    status: true,

}

s := student{

     name:   "Nareenart",

     code:   1044,

     gpa:    4.00,

     status: true,

}

fmt.Println("Name", s.name)

fmt.Println("Code", s.code)

fmt.Println("GPA", s.gpa)

fmt.Println("Status", s.status)

}