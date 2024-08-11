package models

func (m *testRepository) AllTelevision() ([]*Television, error) {
	return nil, nil
}

func (m *testRepository) GetTelevisionByID(id int) (*Television, error) {
	return nil, nil
}

func (m *testRepository) GetTelevisionByName(id string) (*Television, error) {
	return nil, nil
}

func (m *testRepository) RandomTelevision() (*Television, error) {
	return nil, nil
}

func (m *testRepository) AllTVShows() ([]*TVShow, error) {
	return nil, nil
}

func (m *testRepository) GetTVShowByName(name string) (*TVShow, error) {
	return nil, nil
}

//func (m *mysqlRepository) RandomTelevisionByGenre() ([]*Television, error) {
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//
//	query := `select ID, Name, MediaType, Genre, coalesce(Details, ''), Country, ReleaseDate
//	   from media_stream.Television where Genre = 'Comedy' order by rand() limit 1`
//
//	var television Television
//
//	err := m.DB.QueryRowContext(ctx, query).Scan(
//		&television.ID,
//		&television.Name,
//		&television.MediaType,
//		&television.Genre,
//		&television.Details,
//		&television.Country,
//		&television.ReleaseDate)
//	if err != nil {
//		log.Println("Internal Server Error", http.StatusInternalServerError)
//		return nil, err
//	}
//
//	return &[]Television{television}, nil
//}
