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
  radius: Float
  color: String
  price_new: Float
  usage_state: State
}

type Binding {

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

enum State {
  NEW
  USED
  DEFECT
  DISPOSED
  SOLD
}

```
