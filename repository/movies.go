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
		return movies, err

	}

	defer rows.Close()

	for rows.Next() {
		var movie structs.Movies

		err := rows.Scan(&movie.ID, &movie.Name, &movie.Slug, &movie.Category, &movie.Video_url, &movie.Thumbnail_url, &movie.Rating, &movie.Is_featured, &movie.Created_at, &movie.Updated_at)

		if err != nil {
			return movies, err

		}

		movies = append(movies, movie)
	}

	return movies, nil
}

// GetByID is the function to get Movies by id
func (m *MoviesRepo) GetByID(id int) (structs.Movies, error) {
	var movie structs.Movies

	err := m.DB.QueryRow("SELECT * FROM movies WHERE id = $1", id).Scan(&movie.ID, &movie.Name, &movie.Slug, &movie.Category, &movie.Video_url, &movie.Thumbnail_url, &movie.Rating, &movie.Is_featured, &movie.Created_at, &movie.Updated_at)

	if err != nil {
		return movie, err

	}

	return movie, nil
}

// Insert is the function to insert Movies
func (m *MoviesRepo) Insert(movie structs.Movies, DbConnection *sql.DB) (structs.Movies, error) {
	err := DbConnection.QueryRow("INSERT INTO movies(name, slug, category, video_url, thumbnail, rating, is_featured, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id", movie.Name, movie.Slug, movie.Category, movie.Video_url, movie.Thumbnail_url, movie.Rating, movie.Is_featured, movie.Created_at, movie.Updated_at).Scan(&movie.ID)

	if err != nil {
		return movie, err
	}

	return movie, nil
}

// Update is the function to update Movies
func (m *MoviesRepo) Update(id int, movie structs.Movies) (structs.Movies, error) {
	_, err := m.DB.Exec("UPDATE movies SET name = $1, slug = $2, category = $3, video_url = $4, thumbnail = $5, rating = $6, is_featured = $7, updated_at = $8 WHERE id = $9", movie.Name, movie.Slug, movie.Category, movie.Video_url, movie.Thumbnail_url, movie.Rating, movie.Is_featured, movie.Updated_at, id)

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
