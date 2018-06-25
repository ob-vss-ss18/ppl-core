package models

type Stick struct {
	Id           uint32
	Usage        Usage
	Usertype     Usertype
	Gender       Gender
	Manufacturer string
	Model        string
	Length       uint8
	Bodyheight   uint8
	GripKind     string
	Color        string
	Condition    Condition
	PriceNew     float32
	Status       Availability
}