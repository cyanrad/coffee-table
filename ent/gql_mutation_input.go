// Code generated by ent, DO NOT EDIT.

package ent

// CreateCoffeeInput represents a mutation input for creating coffees.
type CreateCoffeeInput struct {
	Sugar        int
	Coffee       int
	PowderedMilk int
	CoffeeMate   int
	Milk         int
	Water        int
	Rating       float64
}

// Mutate applies the CreateCoffeeInput on the CoffeeMutation builder.
func (i *CreateCoffeeInput) Mutate(m *CoffeeMutation) {
	m.SetSugar(i.Sugar)
	m.SetCoffee(i.Coffee)
	m.SetPowderedMilk(i.PowderedMilk)
	m.SetCoffeeMate(i.CoffeeMate)
	m.SetMilk(i.Milk)
	m.SetWater(i.Water)
	m.SetRating(i.Rating)
}

// SetInput applies the change-set in the CreateCoffeeInput on the CoffeeCreate builder.
func (c *CoffeeCreate) SetInput(i CreateCoffeeInput) *CoffeeCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateCoffeeInput represents a mutation input for updating coffees.
type UpdateCoffeeInput struct {
	Sugar        *int
	Coffee       *int
	PowderedMilk *int
	CoffeeMate   *int
	Milk         *int
	Water        *int
	Rating       *float64
}

// Mutate applies the UpdateCoffeeInput on the CoffeeMutation builder.
func (i *UpdateCoffeeInput) Mutate(m *CoffeeMutation) {
	if v := i.Sugar; v != nil {
		m.SetSugar(*v)
	}
	if v := i.Coffee; v != nil {
		m.SetCoffee(*v)
	}
	if v := i.PowderedMilk; v != nil {
		m.SetPowderedMilk(*v)
	}
	if v := i.CoffeeMate; v != nil {
		m.SetCoffeeMate(*v)
	}
	if v := i.Milk; v != nil {
		m.SetMilk(*v)
	}
	if v := i.Water; v != nil {
		m.SetWater(*v)
	}
	if v := i.Rating; v != nil {
		m.SetRating(*v)
	}
}

// SetInput applies the change-set in the UpdateCoffeeInput on the CoffeeUpdate builder.
func (c *CoffeeUpdate) SetInput(i UpdateCoffeeInput) *CoffeeUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateCoffeeInput on the CoffeeUpdateOne builder.
func (c *CoffeeUpdateOne) SetInput(i UpdateCoffeeInput) *CoffeeUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
