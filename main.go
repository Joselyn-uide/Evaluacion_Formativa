package main

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

func NuevoMedicine(name string, manufacturer string, manufactureDate string, shelfLife int) Medicine {
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

func NewTablet(name string, manufacturer string, manufactureDate string, shelfLife int, dosePerTablet float64, isPrescription bool) Tablet {
	return Tablet{
		Medicine:       NuevoMedicine(name, manufacturer, manufactureDate, shelfLife),
		dosePerTablet:  dosePerTablet,
		isPrescription: isPrescription,
	}
}

func NewSyrup(name string, manufacturer string, manufactureDate string, shelfLife int, volume float64, flavor string) Syrup {
	return Syrup{
		Medicine: Medicine(name, manufacturer, manufactureDate, shelfLife),
		volume:   volume,
		flavor:   flavor,
	}
}

func main() {
}
