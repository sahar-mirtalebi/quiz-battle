package domain

type Category string

const (
	FootballCategory   = "football"
	BasketballCategory = "basketball"
	CinemaCategory     = "cinema"
	MusicCategory      = "music"
	HistoryCategory    = "history"
	GeographyCategory  = "geography"
	ScienceCategory    = "science"
)

func AllCategories() []Category {
	return []Category{
		FootballCategory,
		BasketballCategory,
		CinemaCategory,
		MusicCategory,
		HistoryCategory,
		GeographyCategory,
		ScienceCategory,
	}
}

func (c Category) IsValid() bool {
	for _, category := range AllCategories() {
		if c == category {
			return true
		}
	}

	return false
}
