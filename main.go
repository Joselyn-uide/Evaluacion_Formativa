package main

import (
	"fmt"
)

type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate string
	shelfLife       int
}

type Tablet struct {
	Medicine
	dosePerTablet  float64
	isPrescription bool
}

type Syrup struct {
	Medicine
	volume float64
	flavor string
}

func NewMedicine(name string, manufacturer string, manufactureDate string, shelfLife int) Medicine {
	return Medicine{
		name:            name,
		manufacturer:    manufacturer,
		manufactureDate: manufactureDate,
		shelfLife:       shelfLife,
	}
}

func (m Medicine) GetName() string {
	return m.name
}

func (m Medicine) GetManufacturer() string {
	return m.manufacturer
}

func (m Medicine) GetManufactureDate() string {
	return m.manufactureDate
}

func (m *Medicine) GetShelfLife() int {
	return m.shelfLife
}

func (m *Medicine) SetName(name string) {
	if name != "" {
		m.name = name
	}
}

func (m *Medicine) SetManufacturer(manufacturer string) {
	if manufacturer != "" {
		m.manufacturer = manufacturer
	}
}

func (m *Medicine) SetManufactureDate(manufactureDate string) {
	if manufactureDate != "" {
		m.manufactureDate = manufactureDate
	}
}

func (m *Medicine) SetShelfLife(shelfLife int) {
	if shelfLife > 0 {
		m.shelfLife = shelfLife
	}
}

// CONSTRUCTORES
func NewTablet(name string, manufacturer string, manufactureDate string, shelfLife int, dosePerTablet float64, isPrescription bool) Tablet {
	return Tablet{
		Medicine:       NewMedicine(name, manufacturer, manufactureDate, shelfLife),
		dosePerTablet:  dosePerTablet,
		isPrescription: isPrescription,
	}
}

func NewSyrup(name string, manufacturer string, manufactureDate string, shelfLife int, volume float64, flavor string) Syrup {
	return Syrup{
		Medicine: NewMedicine(name, manufacturer, manufactureDate, shelfLife),
		volume:   volume,
		flavor:   flavor,
	}
}

// MÉTODOS
func (t Tablet) MostrarInfo() {
	fmt.Printf("[Tableta] Nombre: %s | Fabricante: %s | Dosis: %.2f mg | Requiere receta: %t | Vida útil: %d meses | Fabricación: %s\n",
		t.GetName(), t.GetManufacturer(), t.dosePerTablet, t.isPrescription, t.GetShelfLife(), t.GetManufactureDate())
}

func (s Syrup) MostrarInfo() {
	fmt.Printf("[Jarabe] Nombre: %s | Fabricante: %s | Volumen: %.2f ml | Sabor: %s | Vida útil: %d meses | Fabricación: %s\n",
		s.GetName(), s.GetManufacturer(), s.volume, s.flavor, s.GetShelfLife(), s.GetManufactureDate())
}

func main() {
	inventarioTablets := []Tablet{
		NewTablet("Paracetamol", "Farmacéutica A", "2023-01-15", 24, 500.0, false),
		NewTablet("Ibuprofeno", "Farmacéutica B", "2023-02-10", 36, 200.0, true),
	}

	inventarioJarabes := []Syrup{
		NewSyrup("Jarabe para la tos", "Farmacéutica C", "2023-03-05", 18, 100.0, "Miel"),
		NewSyrup("Jarabe para la fiebre", "Farmacéutica D", "2023-04-20", 12, 150.0, "Frutas"),
	}

}
