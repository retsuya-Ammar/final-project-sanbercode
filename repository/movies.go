package repository

import (
	"database/sql"
	"final-project-sanbercode/structs"
)

// MoviesRepo is the repository for Movies
type MoviesRepo struct {
	DB *sql.DB
}

// NewMoviesRepo is the constructor for Movies repository
func NewMoviesRepo(db *sql.DB) *MoviesRepo {
	return &MoviesRepo{DB: db}
}

// GetAll is the function to get all Movies
func (m *MoviesRepo) GetAll() ([]structs.Movies, error) {
	var movies []structs.Movies

	rows, err := m.DB.Query("SELECT * FROM movies")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var movie structs.Movies
		err = rows.Scan(&movie.ID, &movie.Name, &movie.Slug, &movie.Category, &movie.Video_url, &movie.Thumbnail_url, &movie.Rating, &movie.Is_featured, &movie.Created_at, &movie.Updated_at, &movie.Deleted_at)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

// GetByID is the function to get Movies by id
func (m *MoviesRepo) GetByID(id int) (structs.Movies, error) {
	var movie structs.Movies

	err := m.DB.QueryRow("SELECT * FROM movies WHERE id = $1", id).Scan(&movie.ID, &movie.Name, &movie.Slug, &movie.Category, &movie.Video_url, &movie.Thumbnail_url, &movie.Rating, &movie.Is_featured, &movie.Created_at, &movie.Updated_at, &movie.Deleted_at)
	if err != nil {
		return movie, err
	}

	return movie, nil
}

// Insert is the function to insert Movies
func (m *MoviesRepo) Insert(movie structs.Movies) (structs.Movies, error) {
	err := m.DB.QueryRow("INSERT INTO movies (name, slug, category, video_url, thumbnail_url, rating, is_featured) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", movie.Name, movie.Slug, movie.Category, movie.Video_url, movie.Thumbnail_url, movie.Rating, movie.Is_featured).Scan(&movie.ID)

	if err != nil {
		return movie, err
	}

	return movie, nil
}

// Update is the function to update Movies
func (m *MoviesRepo) Update(movie structs.Movies) (structs.Movies, error) {
	_, err := m.DB.Exec("UPDATE movies SET name = $1, slug = $2, category = $3, video_url = $4, thumbnail_url = $5, rating = $6, is_featured = $7 WHERE id = $8", movie.Name, movie.Slug, movie.Category, movie.Video_url, movie.Thumbnail_url, movie.Rating, movie.Is_featured, movie.ID)

	if err != nil {
		return movie, err
	}

	return movie, nil
}

// Delete is the function to delete Movies
func (m *MoviesRepo) Delete(id int) error {
	_, err := m.DB.Exec("DELETE FROM movies WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
