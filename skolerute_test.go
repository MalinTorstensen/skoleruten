package skoleruten

import (
	"testing"
	"time"
)

func TestHentForKommune(t *testing.T) {
	testData := []Fridag{
		{Dato: time.Now(), Navn: "Nasjonal fridag", Nasjonal: true},
		{Dato: time.Now(), Navn: "Bergen Spesifikk", KommuneID: "4601", Nasjonal: false},
		{Dato: time.Now(), Navn: "Oslo Spesifikk", KommuneID: "0301", Nasjonal: false},
	}

	service := NewService(testData)

	bergenResultat := service.HentForKommune("4601")
	if len(bergenResultat) != 2 {
		t.Errorf("Forventet 2 fridager for Bergen, men fikk %d", len(bergenResultat))
	}

	trondheimResultat := service.HentForKommune("5001")
	if len(trondheimResultat) != 1 {
		t.Errorf("Forventet 1 fridag for Trondheim, men fikk %d", len(trondheimResultat))
	}
}