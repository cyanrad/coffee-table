// Code generated by ent, DO NOT EDIT.

package ent

import (
	"main/ent/coffee"
	"main/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	coffeeFields := schema.Coffee{}.Fields()
	_ = coffeeFields
	// coffeeDescSugar is the schema descriptor for sugar field.
	coffeeDescSugar := coffeeFields[0].Descriptor()
	// coffee.SugarValidator is a validator for the "sugar" field. It is called by the builders before save.
	coffee.SugarValidator = coffeeDescSugar.Validators[0].(func(int) error)
	// coffeeDescCoffee is the schema descriptor for coffee field.
	coffeeDescCoffee := coffeeFields[1].Descriptor()
	// coffee.CoffeeValidator is a validator for the "coffee" field. It is called by the builders before save.
	coffee.CoffeeValidator = coffeeDescCoffee.Validators[0].(func(int) error)
	// coffeeDescPowderedMilk is the schema descriptor for powdered_milk field.
	coffeeDescPowderedMilk := coffeeFields[2].Descriptor()
	// coffee.PowderedMilkValidator is a validator for the "powdered_milk" field. It is called by the builders before save.
	coffee.PowderedMilkValidator = coffeeDescPowderedMilk.Validators[0].(func(int) error)
	// coffeeDescCoffeeMate is the schema descriptor for coffee_mate field.
	coffeeDescCoffeeMate := coffeeFields[3].Descriptor()
	// coffee.CoffeeMateValidator is a validator for the "coffee_mate" field. It is called by the builders before save.
	coffee.CoffeeMateValidator = coffeeDescCoffeeMate.Validators[0].(func(int) error)
	// coffeeDescMilk is the schema descriptor for milk field.
	coffeeDescMilk := coffeeFields[4].Descriptor()
	// coffee.MilkValidator is a validator for the "milk" field. It is called by the builders before save.
	coffee.MilkValidator = coffeeDescMilk.Validators[0].(func(int) error)
	// coffeeDescWater is the schema descriptor for water field.
	coffeeDescWater := coffeeFields[5].Descriptor()
	// coffee.WaterValidator is a validator for the "water" field. It is called by the builders before save.
	coffee.WaterValidator = coffeeDescWater.Validators[0].(func(int) error)
	// coffeeDescRating is the schema descriptor for rating field.
	coffeeDescRating := coffeeFields[6].Descriptor()
	// coffee.RatingValidator is a validator for the "rating" field. It is called by the builders before save.
	coffee.RatingValidator = coffeeDescRating.Validators[0].(func(float64) error)
}