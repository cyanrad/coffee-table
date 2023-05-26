# coffee-table
An application that will help me create the perfect coffee one day...one day **<stares into the far distance>**

to run make sure you have all the modules installed then (localhost:3000 by default):
```bash
go run .
```
## Exmaples
creating coffee mmmmmmmmm yeash cofvee
```graphql
mutation {
  createCoffee(
    input: {sugar: 1, coffee: 1, coffeeMate: 1, powderedMilk: 1, milk: 1, water: 1, rating: 6.5}
  ) {
    sugar
    coffee
    coffeeMate
    powderedMilk
    milk
    water
    rating
  }
}
```

to get initial cursor
```graphql
query {
  coffees{
    edges {
      cursor
      node {
        sugar
      }
    }
  }
}
```

to get even MORE coffee
```graphql
query {
  coffees(first:3, after:<the cursor you got>){
    edges{
      cursor
      node{
        id
        sugar
      }
    }
  }
}
```
