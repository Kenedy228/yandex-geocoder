package geocoder

import (
	"fmt"
)

type Address struct {
	Country  string
	District string
	City     string
	Street   string
	House    string
}

func (a *Address) toGeocode() string {
	return fmt.Sprintf("%s, %s, %s, %s, %s",
		a.Country,
		a.District,
		a.City,
		a.Street,
		a.House)
}
