# Client

Der Client selbst stellt keine Api bereit. Die Funktionalität ist:

## Kunden
Kunden können sich über ihre Emailadresse mittels Passwordless Login anmelden.
* Kunde trägt Emailadresse in Maske ein
* Bekommt entweder einen Fehler in der Oberfläche falls Email ungültig bzw. eine Email mit einem 1-Time-Token
* Diesen kann er/sie in der Oberfläche eingeben und ist somit für eine bestimmte Zeit eingelogged.

### Funktionen
* Kunden können ein Listing der verfügbaren Ski o.Ae. anzeigen lassen
* Können einzelne Items reservieren (gibt Limit)
* Kunde kann sich aktuelle Reservierungen anzeigen lassen und diese Löschen


## Betreiber
Der Betreiber sowie die Mitarbeiter können sich über eine Email-Password Kombination anmelden. Die Hauptfunktionalität besteht in der Abwicklung von Verleih- und Leasingvorgängen.
* Listing aktueller Vorgänge für Kunden
* Suchen nach Kunden, Anlegen von Neukunden
* Abschliessen von Leasingvorgängen
* Listing von Inventar
* Anlegen neuer Items im Inventar
* ...
