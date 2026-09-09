package skoleruten

import "time"

type Fridag struct {
	Dato		time.Time	`json:"dato"`
	DatoTil		time.Time	`json:"dato_til,omitempty"`
	Navn		string		`json:"navn"`
	Notat		string		`json:"notat"`
	KommuneID	string		`json:"kommune_id"`
	Nasjonal	bool		`json:"nasjonal"`
}

type SkoleruteService struct {
	data []Fridag
}

func NewService(data []Fridag) *SkoleruteService {
	return &SkoleruteService{data: data}
}

func (s *SkoleruteService) HentForKommune(KommuneID string) []Fridag {
	var resultat []Fridag
	for _, dag := range s.data {
		if dag.Nasjonal || dag.KommuneID == KommuneID {
			resultat = append(resultat, dag)
		}
	}
	return resultat
}