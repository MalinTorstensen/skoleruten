package skoleruten

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type fridagDefinition struct {
	Date      string `json:"date,omitempty"`
	DateTo    string `json:"date_to,omitempty"`
	Month     int    `json:"month,omitempty"`
	Day       int    `json:"day,omitempty"`
	Name      string `json:"name"`
	Note      string `json:"note,omitempty"`
	KommuneID string `json:"kommune_id,omitempty"`
	Nasjonal  bool   `json:"nasjonal,omitempty"`
	Repeat    string `json:"repeat,omitempty"`
}

func LoadFridagerFromDefinitions(defs []fridagDefinition, startYear, endYear int) ([]Fridag, error) {
	if startYear > endYear {
		return nil, fmt.Errorf("startYear %d must not be after endYear %d", startYear, endYear)
	}

	var result []Fridag
	for _, def := range defs {
		if def.Date != "" {
			parsed, err := time.Parse("2006-01-02", def.Date)
			if err != nil {
				return nil, fmt.Errorf("invalid date %q: %w", def.Date, err)
			}

			parsedTo := parsed
			if def.DateTo != "" {
				parsedTo, err = time.Parse("2006-01-02", def.DateTo)
				if err != nil {
					return nil, fmt.Errorf("invalid date_to %q: %w", def.DateTo, err)
				}
				if parsedTo.Before(parsed) {
					return nil, fmt.Errorf("date_to %q is before date %q", def.DateTo, def.Date)
				}
			}

			if parsedTo.Year() < startYear || parsed.Year() > endYear {
				continue
			}

			result = append(result, Fridag{
				Dato:      parsed,
				DatoTil:   parsedTo,
				Navn:      def.Name,
				Notat:     def.Note,
				KommuneID: def.KommuneID,
				Nasjonal:  def.Nasjonal,
			})
			continue
		}

		if def.Repeat == "yearly" {
			if def.Month < 1 || def.Month > 12 || def.Day < 1 || def.Day > 31 {
				return nil, fmt.Errorf("invalid month/day for yearly repeat: %d/%d", def.Month, def.Day)
			}

			for year := startYear; year <= endYear; year++ {
				date := time.Date(year, time.Month(def.Month), def.Day, 0, 0, 0, 0, time.UTC)
				if int(date.Month()) != def.Month || date.Day() != def.Day {
					return nil, fmt.Errorf("invalid recurring date %d-%02d-%02d", year, def.Month, def.Day)
				}
				result = append(result, Fridag{
					Dato:      date,
					Navn:      def.Name,
					Notat:     def.Note,
					KommuneID: def.KommuneID,
					Nasjonal:  def.Nasjonal,
				})
			}
			continue
		}

		return nil, fmt.Errorf("definition must contain either date or repeat: %+v", def)
	}

	return result, nil
}

func LoadFridagerJSON(data []byte, startYear, endYear int) ([]Fridag, error) {
	var defs []fridagDefinition
	if err := json.Unmarshal(data, &defs); err != nil {
		return nil, err
	}

	return LoadFridagerFromDefinitions(defs, startYear, endYear)
}

func readSourceFile(name string) ([]byte, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("unable to determine current source path")
	}
	path := filepath.Join(filepath.Dir(filename), name)
	return os.ReadFile(path)
}

func DefaultFridager(startYear, endYear int) ([]Fridag, error) {
	data, err := readSourceFile("fridager.json")
	if err != nil {
		return nil, err
	}
	return LoadFridagerJSON(data, startYear, endYear)
}

func DefaultFridagerForYear(year int) ([]Fridag, error) {
	return DefaultFridager(year, year)
}
