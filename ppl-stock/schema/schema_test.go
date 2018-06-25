package schema

import (
	"testing"
	"github.com/ob-vss-ss18/ppl-stock/models"
)

func TestSchemaExampleSki(t *testing.T) {
	var testSki = loadSkiFromDatabase(10)
	if testSki.Id != 10 {
		t.Error("Expected 10, got ", testSki.Id)
	}
	if testSki.Usage != models.Langlauf {
		t.Error("Expected 'Langlauf', got ", testSki.Usage)
	}
	if testSki.Category != models.Beginner {
		t.Error("Expected 'Beginner', got ", testSki.Category)
	}
	if testSki.Usertype != models.Erwachsener {
		t.Error("Expected 'Erwachsene', got ", testSki.Usertype)
	}
	if testSki.Gender != models.Male {
		t.Error("Expected 'Männlich', got ", testSki.Gender)
	}
	if testSki.Manufacturer != "Fischer" {
		t.Error("Expected 'Fischer', got ", testSki.Manufacturer)
	}
	if testSki.Model != "Super Ski 3000" {
		t.Error("Expected 'Super Ski 3000', got ", testSki.Model)
	}
	if testSki.Length != 2 {
		t.Error("Expected 2, got ", testSki.Length)
	}
	if testSki.Bodyheight != 3 {
		t.Error("Expected 3, got ", testSki.Bodyheight)
	}
	if testSki.Bodyweight != 4 {
		t.Error("Expected 4, got ", testSki.Bodyweight)
	}
	if testSki.Color != "Rot" {
		t.Error("Expected 'Rot', got ", testSki.Color)
	}
	if testSki.PriceNew != 34.99 {
		t.Error("Expected 34.99, got ", testSki.PriceNew)
	}
	if testSki.Condition != models.New {
		t.Error("Expected 'Neu', got ", testSki.Condition)
	}
	if testSki.Status != models.Available {
		t.Error("Expected 'Verfügbar', got ", testSki.Status)
	}
}

func TestSchemaExampleStick(t *testing.T) {
	var testStick = loadStickFromDatabase(5)

	if testStick.Id != 5 {
		t.Error("Expected 5, got ", testStick.Id)
	}
	if testStick.Usage != models.All_Mountain {
		t.Error("Expected 'All Mountain', got ", testStick.Usage)
	}
	if testStick.Usertype != models.Kinder {
		t.Error("Expected 'Kinder', got ", testStick.Usertype)
	}
	if testStick.Gender != models.Uni {
		t.Error("Expected 'Uni', got ", testStick.Gender)
	}
	if testStick.Manufacturer != "Müller" {
		t.Error("Expected 'Müller', got ", testStick.Manufacturer)
	}
	if testStick.Model != "Super Stick 7000" {
		t.Error("Expected 'Super Stick 7000', got ", testStick.Model)
	}
	if testStick.Length != 25 {
		t.Error("Expected 25, got ", testStick.Length)
	}
	if testStick.Bodyheight != 13 {
		t.Error("Expected 13, got ", testStick.Bodyheight)
	}
	if testStick.Color != "Blau" {
		t.Error("Expected 'Blau', got ", testStick.Color)
	}
	if testStick.PriceNew != 14.99 {
		t.Error("Expected 14.99, got ", testStick.PriceNew)
	}
	if testStick.Condition != models.Used {
		t.Error("Expected 'Benutzt', got ", testStick.Condition)
	}
	if testStick.Status != models.Available {
		t.Error("Expected 'Verfügbar', got ", testStick.Status)
	}
}
