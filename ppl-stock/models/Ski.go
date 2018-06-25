package models

type Ski struct {
	Id           uint32
	Usage        Usage
	Category     Category
	Usertype     Usertype
	Gender       Gender
	Manufacturer string
	Model        string
	Length       uint8
	Bodyheight   uint8
	Bodyweight   uint8
	Color        string
	PriceNew     float32
	Condition    Condition
	Status       Availability
}