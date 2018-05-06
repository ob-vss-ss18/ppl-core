```graphql
type Customer {
  id: Int!
  name: String!
  surname: String!
  address: Address!
  skill: Skill
  email: String!
  telephone: String!
  birthday: Date
}

type Address {
  street: String!
  number: Int!
  zipcode: Int!
  city: String!
}

enum Skill {
  BEGINNER
  ADVANCED
  PRO
}
```
