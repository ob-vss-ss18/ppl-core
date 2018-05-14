```graphql
type Ski {
  id: Int!
  usage: [Usage]
  category: [Category]
  usertype: [Usertype]
  sex: [Sex]
  manufacturer: String!
  modell: String!
  length: Int
  bodyheight: Int
  bodyweight: Int
  color: String
  price_new: Float
  usage_state: State
  condition: Condition
  availability: Availability
}

enum Usage {
  RACE
  PISTE
  ALL_MOUNTAIN
  FREERIDE
  TOUR
  LANGLAUF
}

enum Category {
  TOP
  SPORT
  ALLROUND
  EINSTEIGER
  PARK_PIPE
  CLASSIC
  SKATE
}

enum Usertype {
  CHILD
  TEEN
  ADULT
}

enum Sex {
  UNI
  MALE
  FEMALE
  GIRL
  BOY
}

enum Condition {
  NEW
  USED
  DEFECT
}

enum Availability {
  AVAILABLE
  RESERVED
  LEASED
  SOLD
  DISPOSED
}

```
