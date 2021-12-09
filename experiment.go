package main

import (
	"fmt"

	"gorm.io/gorm"
)

type Person struct {
	Gm gorm.Model // avoid name conflict with Model
	Mm Model      // avoid name conflict with gorm.Model

	Coffee
	Name string `gorm:"column:name"`
	Type int
}

type Model struct {
	Name string
	Date string
}

type Coffee struct {
	Type string
}

func (c *Coffee) Brew() string {
	return fmt.Sprintf("Brewing %s", c.Type)
}

func RunExperiment() {
	p := Person{Name: "Rizal Dwi Andhika"}
	p.Coffee.Type = "Arabica"
	p.Type = 10

	fmt.Println(p.Brew())
}
