package tmdb

import (
	"fmt"
)

// Genre struct
type Genre struct {
	Genres []struct {
		ID   int
		Name string
	}
}

// GetMovieGenres gets the list of movie genres
// http://docs.themoviedb.apiary.io/#reference/genres/genremovielist/get
func (tmdb *TMDb) GetMovieGenres(options map[string]string) (*Genre, error) {
	var availableOptions = map[string]struct{}{
		"language": {}}
	var movieGenres Genre
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/genre/movie/list?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, tmdb.client, &movieGenres)
	return result.(*Genre), err
}

// GetTvGenres gets the list of TV genres
// http://docs.themoviedb.apiary.io/#reference/genres/genretvlist/get
func (tmdb *TMDb) GetTvGenres(options map[string]string) (*Genre, error) {
	var availableOptions = map[string]struct{}{
		"language": {}}
	var tvGenres Genre
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/genre/tv/list?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, tmdb.client, &tvGenres)
	return result.(*Genre), err
}
