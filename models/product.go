package models

type Stock struct {
	ID     int
	Title  string
	Descr  string
	Price  int
	Bought int
	Amount int
}

// func (s *Store) InsertStock(stock Stock) (error) {
// 	var ID int
// 	query := sq.Insert("pets").
// 		Columns("pet_name", "pet_kind", "pet_age").
// 		Values(name, kind, age).
// 		Suffix("RETURNING \"id\"").
// 		RunWith(s.DB).
// 		PlaceholderFormat(sq.Dollar)

// 	err := query.QueryRow().Scan(&ID)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return ID, nil
// }