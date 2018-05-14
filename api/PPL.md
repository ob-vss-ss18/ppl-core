# Peak Power Leasing Prototyp

API und Systemdesign für einen Prototyp des Ski-Verleih-Systems "Peak Power Leasing".

## Komponenten

| Komponente | Heroku | Postgres-DB |
|---|:---:|:---:|
|Client|-|-|
|Auth|:ballot_box_with_check:|:ballot_box_with_check:|
|Stock|:ballot_box_with_check:|:ballot_box_with_check:|
|Customer|:ballot_box_with_check:|:ballot_box_with_check:|
|PriceCalculator|:ballot_box_with_check:|-|
|Reservation|:ballot_box_with_check:|:ballot_box_with_check:|
|Billing&Mailing|:ballot_box_with_check:|:ballot_box_with_check:|
|Leasing|:ballot_box_with_check:|:ballot_box_with_check:|

## Authentifizierung

Die Authentifizierung im System erfolgt über User-Objekte. Diese enthalten die E-Mail des Users, einen Api-Token sowie die Rolle des Users. Diese User-Objekte können beim Auth-Service auf Gültigkeit überprüft werden (siehe Auth).
