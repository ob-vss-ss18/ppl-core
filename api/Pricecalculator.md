# PriceCalculator

Der PriceCalculator-Service ist für die Berechnung von Leasing- und Verleihvorgängen zuständig. Er erhält über API eine Liste von Items sowie die Art (Leasing/Verleih) und evtl. vorhandene (Familien-) Rabatte und Vergünstigungen und erstellt damit eine Auflisrung der Kosten und die entgültigen Kosten. Dieses Listung wird als Grundlage für die Rechnun genommen.

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
