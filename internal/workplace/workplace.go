package workplace

type Vacancy struct{ Room string }

type Period struct{}

type Workplace struct {
	v []Vacancy
}

func New() *Workplace {
	return &Workplace{}
}

func (wp *Workplace) AddRoom(s string) {
	wp.v = append(wp.v, Vacancy{Room: s})
}

func (wp *Workplace) AvailableRooms(p Period) []Vacancy {
	return wp.v
}
