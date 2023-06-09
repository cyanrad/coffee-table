// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CoffeesColumns holds the columns for the "coffees" table.
	CoffeesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "sugar", Type: field.TypeInt},
		{Name: "coffee", Type: field.TypeInt},
		{Name: "powdered_milk", Type: field.TypeInt},
		{Name: "coffee_mate", Type: field.TypeInt},
		{Name: "milk", Type: field.TypeInt},
		{Name: "water", Type: field.TypeInt},
		{Name: "rating", Type: field.TypeFloat64},
	}
	// CoffeesTable holds the schema information for the "coffees" table.
	CoffeesTable = &schema.Table{
		Name:       "coffees",
		Columns:    CoffeesColumns,
		PrimaryKey: []*schema.Column{CoffeesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CoffeesTable,
	}
)

func init() {
}
