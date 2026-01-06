package models

import (
	"encoding/json"
	"testing"

	"github.com/laiambryant/tcgdex/enums"
)

func TestCardResume_GetImageURL(t *testing.T) {
	var cr CardResume
	if got := cr.GetImageURL(enums.QualityLow, enums.ExtensionPng); got != nil {
		t.Fatalf("expected nil image URL, got %v", *got)
	}
	img := "http://img.example.com/cards/1"
	cr.Image = &img
	got := cr.GetImageURL(enums.QualityHigh, enums.ExtensionWebp)
	if got == nil {
		t.Fatalf("expected non-nil image URL")
	}
	want := "http://img.example.com/cards/1/high.webp"
	if *got != want {
		t.Fatalf("unexpected image url: want %s got %s", want, *got)
	}
}

func TestDamage_UnmarshalJSON(t *testing.T) {
	var d Damage
	if err := json.Unmarshal([]byte("10"), &d); err != nil {
		t.Fatalf("unexpected error unmarshaling number: %v", err)
	}
	if string(d) != "10" {
		t.Fatalf("unexpected damage value: want 10 got %s", d)
	}
	if err := json.Unmarshal([]byte("\"50+\""), &d); err != nil {
		t.Fatalf("unexpected error unmarshaling string: %v", err)
	}
	if string(d) != "50+" {
		t.Fatalf("unexpected damage value: want 50+ got %s", d)
	}
	if err := json.Unmarshal([]byte("true"), &d); err == nil {
		t.Fatalf("expected error unmarshaling invalid damage, got nil")
	}
}
