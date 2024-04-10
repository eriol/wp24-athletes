package database // import "github.com/eriol/wp24-athletes/database"

type Athlete struct {
	Slug      string `json:"slug"`
	Name      string `json:"name"`
	Age       uint8  `json:"age"`
	Sport     string `json:"sport"`
	FamousFor string `json:"famous_for"`
}

func GetAthletes() ([]Athlete, error) {

	query := `
    SELECT
        athletes.slug,
        athletes.name,
        athletes.gender,
        sports.name,
        athletes.famous_for
    FROM
        athletes
    INNER JOIN
        sports
    ON
        athletes.sport_id = sports.id;`

	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}

	athletes := make([]Athlete, 0)
	for rows.Next() {
		athlete := Athlete{}

		err = rows.Scan(&athlete.Slug, &athlete.Name, &athlete.Age, &athlete.Sport, &athlete.FamousFor)
		if err != nil {
			return nil, err
		}

		athletes = append(athletes, athlete)
	}

	return athletes, nil
}
