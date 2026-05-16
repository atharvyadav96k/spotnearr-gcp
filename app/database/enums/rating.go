package enums

type Stars int

const (
	One   Stars = 1
	Two   Stars = 2
	Three Stars = 3
	Four  Stars = 4
	Five  Stars = 5
)

func (s Stars) IsValid() bool {
	return s == One || s == Two || s == Three || s == Four || s == Five
}
