package models

type Usage string
const  (
	Langlauf Usage = "Langlauf"
	Piste Usage = "Piste"
	Park  Usage = "Park"  // Original in Category but Example Data lists it here
	Race Usage = "Race"
	All_Mountain Usage = "All Mountain"
	Freeride Usage = "Freeride"
	Tour Usage = "Tour"
	Aplin Usage = "Alpin"
)

type Category string
const  (
	Skate Category = "Skate"
	Allround Category = "Allround"
	Twintip Category = "Twintip"
	Top Category = "Top"
	Sport Category = "Sport"
	Beginner Category = "Beginner"
	Classic Category = "Classic"
)

type Usertype string
const  (
	Kinder Usertype = "Kinder"
	KinderJugendliche Usertype = "Kinder / Jugend"
	Jugend Usertype = "Jugendliche"
	JugendErwachsen Usertype= "Jugend / Erw."
	Erwachsener Usertype = "Erwachsene"
)

type Gender string
const  (
	Male Gender = "Männlich"
	Female Gender = "Weiblich"
	Uni Gender = "Uni"
)

type Condition string
const  (
	New Condition = "Neu"
	Used Condition = "Benutzt"
	Defect Condition = "Defekt"
)

type Availability string
const  (
	Available Availability = "Verfügbar"
	Reserved Availability = "Reserviert"
	Leased Availability = "Ausgeliehen"
	Sold Availability = "Verkauft"
	Disposed Availability = "Entsorgt"
)