package entertainment

import "errors"

type EntertainmentInterface interface {
	SetType(t string) *Entertainment
	SetMediaType(n string) *Entertainment
	SetGenre(g string) *Entertainment
	SetDescription(d string) *Entertainment
	SetCountry(c string) *Entertainment
	SetRating(r int) *Entertainment
	SetReleaseYear(r int) *Entertainment
	SetAvailable(a bool) *Entertainment
	Build() (*Entertainment, error)
}

func NewEntertainmentBuilder() EntertainmentInterface {
	return &Entertainment{}
}

func (e *Entertainment) SetType(t string) *Entertainment {
	e.Type = t
	return e
}

func (e *Entertainment) SetMediaType(n string) *Entertainment {
	e.MediaTitle = n
	return e
}

func (e *Entertainment) SetGenre(g string) *Entertainment {
	e.Genre = g
	return e
}

func (e *Entertainment) SetDescription(d string) *Entertainment {
	e.Description = d
	return e
}

func (e *Entertainment) SetCountry(c string) *Entertainment {
	e.Country = c
	return e
}

func (e *Entertainment) SetRating(r int) *Entertainment {
	e.Rating = r
	return e
}

func (e *Entertainment) SetReleaseYear(r int) *Entertainment {
	e.ReleaseYear = r
	return e
}

func (e *Entertainment) SetAvailable(a bool) *Entertainment {
	e.Available = a
	return e
}

func (e *Entertainment) Build() (*Entertainment, error) {
	if e.Rating < 0 || e.Rating > 10 {
		return nil, errors.New("rating must be between 0 and 10")
	}
	return e, nil
}
