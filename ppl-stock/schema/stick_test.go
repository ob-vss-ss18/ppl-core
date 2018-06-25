package schema

import (
	"testing"
	"github.com/graphql-go/graphql"
	"github.com/ob-vss-ss18/ppl-stock/models"
)

func TestSchemaStick(t *testing.T) {
	var testStick = loadStickFromDatabase(5)
	var parameters = graphql.ResolveParams{}
	parameters.Source = testStick
	var field, ok = stickType.Fields()["id"].Resolve(parameters)
	if field != uint32(5) {
		t.Error("Expected 5, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["usage"].Resolve(parameters)
	if field != models.All_Mountain {
		t.Error("Expected All Mountain, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["usertype"].Resolve(parameters)
	if field != models.Kinder {
		t.Error("Expected Kinder, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["gender"].Resolve(parameters)
	if field != models.Uni {
		t.Error("Expected Uni, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["manufacturer"].Resolve(parameters)
	if field != "Müller" {
		t.Error("Expected Müller, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["modell"].Resolve(parameters)
	if field != "Super Stick 7000" {
		t.Error("Expected Super Stick 7000, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["length"].Resolve(parameters)
	if field != uint8(25) {
		t.Error("Expected 25, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["bodyheight"].Resolve(parameters)
	if field != uint8(13) {
		t.Error("Expected 13, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["grip_kind"].Resolve(parameters)
	if field != "sticky" {
		t.Error("Expected sticky, got ", field)
	}
	field, ok = stickType.Fields()["color"].Resolve(parameters)
	if field != "Blau" {
		t.Error("Expected Blau, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["price_new"].Resolve(parameters)
	if field != float32(14.99) {
		t.Error("Expected 14.99, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["condition"].Resolve(parameters)
	if field != models.Used {
		t.Error("Expected Benutzt, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["availability"].Resolve(parameters)
	if field != models.Available {
		t.Error("Expected Verfügbar, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
}
func TestSchemaStickNil(t *testing.T) {
	var parameters = graphql.ResolveParams{}
	var field, ok = stickType.Fields()["id"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["usage"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["usertype"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["gender"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["manufacturer"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["modell"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["length"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["bodyheight"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["grip_kind"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["color"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["price_new"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["condition"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}
	field, ok = stickType.Fields()["availability"].Resolve(parameters)
	if field != nil {
		t.Error("Expected nil, got ", field)
	}
	if ok != nil {
		t.Error("Expected nil, got ", ok)
	}

}
