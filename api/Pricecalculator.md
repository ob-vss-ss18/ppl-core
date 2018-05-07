```graphql
input Calculation {
  type: CalcType
  items: [Item]
  family_discount: Int
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
