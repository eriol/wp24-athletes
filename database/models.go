package database // import "github.com/eriol/wp24-athletes/database"

type Athlete struct {
	Slug      string `json:"slug"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Age       uint8  `json:"age"`
	Sport     string `json:"sport"`
	FamousFor string `json:"famous_for"`
}

type SearchType int

const (
	SearchByName SearchType = iota
	SearchBySport
	SearchByFamousFor
)

func GetAthletes() ([]Athlete, error) {

	query := `
    SELECT
        athletes.slug,
        athletes.name,
        athletes.gender,
        athletes.age,
        sports.name,
        athletes.famous_for
    FROM
        athletes
    INNER JOIN
        sports
    ON
        athletes.sport_id = sports.id
    ORDER BY
        athletes.name;`

	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}

	athletes := make([]Athlete, 0)
	for rows.Next() {
		athlete := Athlete{}

		err = rows.Scan(
			&athlete.Slug,
			&athlete.Name,
			&athlete.Gender,
			&athlete.Age,
			&athlete.Sport,
			&athlete.FamousFor,
		)
		if err != nil {
			return nil, err
		}

		athletes = append(athletes, athlete)
	}

	return athletes, nil
}

func GetAthlete(slug string) (Athlete, error) {

	query := `
    SELECT
        athletes.slug,
        athletes.name,
        athletes.gender,
        athletes.age,
        sports.name,
        athletes.famous_for
    FROM
        athletes
    INNER JOIN
        sports
    ON
        athletes.sport_id = sports.id
    WHERE
        athletes.slug = ?;`

	athlete := Athlete{}

	if err := database.QueryRow(query, slug).Scan(
		&athlete.Slug,
		&athlete.Name,
		&athlete.Gender,
		&athlete.Age,
		&athlete.Sport,
		&athlete.FamousFor,
	); err != nil {
		return athlete, err
	}

	return athlete, nil
}

func GetRandomAthlete() (Athlete, error) {

	query := `
    SELECT
        athletes.slug,
        athletes.name,
        athletes.gender,
        athletes.age,
        sports.name,
        athletes.famous_for
    FROM
        athletes
    INNER JOIN
        sports
    ON
        athletes.sport_id = sports.id
    ORDER BY RANDOM()
    LIMIT 1;`

	athlete := Athlete{}

	if err := database.QueryRow(query).Scan(
		&athlete.Slug,
		&athlete.Name,
		&athlete.Gender,
		&athlete.Age,
		&athlete.Sport,
		&athlete.FamousFor,
	); err != nil {
		return athlete, err
	}

	return athlete, nil
}

func Search(t SearchType, q string) ([]Athlete, error) {

	var query string
	if t == SearchByName {
		query = `
    SELECT
        athletes.slug,
        athletes.name,
        athletes.gender,
        athletes.age,
        sports.name,
        athletes.famous_for
    FROM
        athletes
    INNER JOIN
        sports
    ON
        athletes.sport_id = sports.id
    WHERE
        athletes.name LIKE '%'||?||'%'
    ORDER BY
        athletes.name;`
	} else if t == SearchBySport {
		query = `
    SELECT
        athletes.slug,
        athletes.name,
        athletes.gender,
        athletes.age,
        sports.name,
        athletes.famous_for
    FROM
        athletes
    INNER JOIN
        sports
    ON
        athletes.sport_id = sports.id
    WHERE
        sports.name LIKE '%'||?||'%';`
	} else if t == SearchByFamousFor {
		query = `
    SELECT
        athletes.slug,
        athletes.name,
        athletes.gender,
        athletes.age,
        sports.name,
        athletes.famous_for
    FROM
        athletes
    INNER JOIN
        sports
    ON
        athletes.sport_id = sports.id
    WHERE
        athletes.famous_for LIKE '%'||?||'%';`
	}

	rows, err := database.Query(query, q)
	if err != nil {
		return nil, err
	}

	athletes := make([]Athlete, 0)
	for rows.Next() {
		athlete := Athlete{}

		err = rows.Scan(
			&athlete.Slug,
			&athlete.Name,
			&athlete.Gender,
			&athlete.Age,
			&athlete.Sport,
			&athlete.FamousFor,
		)
		if err != nil {
			return nil, err
		}

		athletes = append(athletes, athlete)
	}

	return athletes, nil
}
