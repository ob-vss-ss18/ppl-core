package schema

import (
	"testing"
	"github.com/graphql-go/graphql"
	"github.com/ob-vss-ss18/ppl-stock/models"
)

func TestSchemaSki(t *testing.T) {
	var testSki = loadSkiFromDatabase(10)
	var parameters = graphql.ResolveParams{}
	parameters.Source = testSki
	var field, ok = skiType.Fields()["id"].Resolve(parameters)
	if field != uint32(10) {
		t.Error("Expected 10, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["usage"].Resolve(parameters)
	if field != models.Langlauf {
		t.Error("Expected Langlauf, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["category"].Resolve(parameters)
	if field != models.Beginner {
		t.Error("Expected Beginner, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["usertype"].Resolve(parameters)
	if field != models.Erwachsener {
		t.Error("Expected Erwachsene, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["gender"].Resolve(parameters)
	if field != models.Male {
		t.Error("Expected Männlich, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["manufacturer"].Resolve(parameters)
	if field != "Fischer" {
		t.Error("Expected Fischer, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["modell"].Resolve(parameters)
	if field != "Super Ski 3000" {
		t.Error("Expected Super Ski 3000, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["length"].Resolve(parameters)
	if field != uint8(2) {
		t.Error("Expected 2, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["bodyheight"].Resolve(parameters)
	if field != uint8(3) {
		t.Error("Expected 3, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["bodyweight"].Resolve(parameters)
	if field != uint8(4) {
		t.Error("Expected 4, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["color"].Resolve(parameters)
	if field != "Rot" {
		t.Error("Expected Rot, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["price_new"].Resolve(parameters)
	if field != float32(34.99) {
		t.Error("Expected 34.99, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["condition"].Resolve(parameters)
	if field != models.New {
		t.Error("Expected Neu, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["availability"].Resolve(parameters)
	if field != models.Available {
		t.Error("Expected Verfügbar, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
}
func TestSchemaSkiNil(t *testing.T) {
	var parameters = graphql.ResolveParams{}
	var field, ok = skiType.Fields()["id"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["usage"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["category"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["usertype"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["gender"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["manufacturer"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["modell"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["length"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["bodyheight"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["bodyweight"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["color"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["price_new"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["condition"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = skiType.Fields()["availability"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}

}
