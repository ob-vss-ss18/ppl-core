# Ablauf

## Kundenlogin
* Kunde gibt Email in Client ein
* Client frägt bei ppl-auth an
* ppl-auth prüft ob User für Email bereits vorhanden
 * Falls ja, gibt OK zurück, erzeugt neuen Token und sendet Email an Adresse
 * Falls nein, prüfe bei ppl-customer ob Kunde vorhanden
  * Falls nein, gibt Fehler zurück
  * Falls ja, lege neuen User für Kunden an, gibt OK zurück, erzeugt neuen Token und sendet Email an Adresse

* Kunde gibt Token aus Email in Maske in Client ein
* Client sendet diese an ppl-auth, welcher diese prüft
 * Falls ok, sendet User mit Token zurück
 * Falls nicht, sendet Fehler zurück

## Mitarbeiterlogin
* Mitarbeiter gibt Email und Passwort in Client ein
* Client sendet diese an ppl-auth
* ppl-auth prüft ob diese Korrekt
 * Falls ja, gibt OK mit User zurück
 * Falls nein, gibt Error zurück

## Reservierung Kunde
* Kunde ist eingelogged
* Kunde kann über Client ein Listing verfügbarer Items anzeigen lassen
 * -> Anfrage an ppl-stock mit User-Objekt und entsprechendem Filter
  * ppl-stock überprüft User mit ppl-auth
 * Items werden in Client aufgelistet
* Nach erfolgreicher Auswahl der zu reservierenden Items: Anfrage an ppl-reservation mit User und Items
 * ppl-reservation prüft User und prüft Verfügbarkeit von Items, setzt diese auf Status "Reserviert" und legt eine Reservierung in der Datenbank an, gibt OK zurück und versendet Bestätigung per Mail

* Kunde kann über Client seiner bereits bestehenden Reservierungen anfragen
 * Anfrage an ppl-reservatino mit User, Listing als Rückgabe
 * Kunde kann Reservierung löschen mit Anfrage an ppl-reservation

## Ausleihvorgang
