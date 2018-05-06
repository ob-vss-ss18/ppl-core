```graphql
input Calculation {
  type: CalcType
}

enum CalcType {
  RENTAL
  LEASING
}

input Item {
  id: Int
  price_new: Float
  discout_perc: Float
  discout_abs: Float
  itemType: ItemType
}

enum ItemType {
  SKI
  SHOE
}

```
