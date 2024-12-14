// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

type Album struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Artist      string   `json:"artist"`
	Genre       *string  `json:"genre,omitempty"`
	ReleaseDate string   `json:"releaseDate"`
	Year        int      `json:"year"`
	Tracks      []*Track `json:"tracks"`
	AddedAt     string   `json:"addedAt"`
}

type Artist struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Genre *string `json:"genre,omitempty"`
}

type Mutation struct {
}

type NewAlbum struct {
	Title       string        `json:"title"`
	Artist      string        `json:"artist"`
	Genre       *string       `json:"genre,omitempty"`
	ReleaseDate string        `json:"releaseDate"`
	Files       []*UploadFile `json:"files"`
}

type NewTrack struct {
	Title        string         `json:"title"`
	Artist       string         `json:"artist"`
	Album        string         `json:"album"`
	ArtistArtist string         `json:"artistArtist"`
	Genre        *string        `json:"genre,omitempty"`
	ReleaseDate  string         `json:"releaseDate"`
	TrackNumber  int            `json:"trackNumber"`
	File         graphql.Upload `json:"file"`
}

type Playlist struct {
	Name      string   `json:"name"`
	Tracks    []*Track `json:"tracks"`
	CreatedAt string   `json:"createdAt"`
}

type Query struct {
}

type Track struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Artist      string  `json:"artist"`
	Album       string  `json:"album"`
	AlbumArtist string  `json:"albumArtist"`
	Genre       *string `json:"genre,omitempty"`
	ReleaseDate string  `json:"releaseDate"`
	Year        int     `json:"year"`
	TrackNumber *int    `json:"trackNumber,omitempty"`
	AddedAt     string  `json:"addedAt"`
}

type UploadFile struct {
	Filename string         `json:"filename"`
	File     graphql.Upload `json:"file"`
}

type Order string

const (
	OrderAsc  Order = "ASC"
	OrderDesc Order = "DESC"
)

var AllOrder = []Order{
	OrderAsc,
	OrderDesc,
}

func (e Order) IsValid() bool {
	switch e {
	case OrderAsc, OrderDesc:
		return true
	}
	return false
}

func (e Order) String() string {
	return string(e)
}

func (e *Order) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Order(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Order", str)
	}
	return nil
}

func (e Order) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SortAlbumsBy string

const (
	SortAlbumsByUnspecified SortAlbumsBy = "UNSPECIFIED"
	SortAlbumsByTitle       SortAlbumsBy = "TITLE"
	SortAlbumsByArtist      SortAlbumsBy = "ARTIST"
	SortAlbumsByRelease     SortAlbumsBy = "RELEASE"
	SortAlbumsByAdded       SortAlbumsBy = "ADDED"
)

var AllSortAlbumsBy = []SortAlbumsBy{
	SortAlbumsByUnspecified,
	SortAlbumsByTitle,
	SortAlbumsByArtist,
	SortAlbumsByRelease,
	SortAlbumsByAdded,
}

func (e SortAlbumsBy) IsValid() bool {
	switch e {
	case SortAlbumsByUnspecified, SortAlbumsByTitle, SortAlbumsByArtist, SortAlbumsByRelease, SortAlbumsByAdded:
		return true
	}
	return false
}

func (e SortAlbumsBy) String() string {
	return string(e)
}

func (e *SortAlbumsBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortAlbumsBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortAlbumsBy", str)
	}
	return nil
}

func (e SortAlbumsBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SortArtistsBy string

const (
	SortArtistsByUnspecified SortArtistsBy = "UNSPECIFIED"
	SortArtistsByName        SortArtistsBy = "NAME"
)

var AllSortArtistsBy = []SortArtistsBy{
	SortArtistsByUnspecified,
	SortArtistsByName,
}

func (e SortArtistsBy) IsValid() bool {
	switch e {
	case SortArtistsByUnspecified, SortArtistsByName:
		return true
	}
	return false
}

func (e SortArtistsBy) String() string {
	return string(e)
}

func (e *SortArtistsBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortArtistsBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortArtistsBy", str)
	}
	return nil
}

func (e SortArtistsBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SortPlaylistsBy string

const (
	SortPlaylistsByUnspecified SortPlaylistsBy = "UNSPECIFIED"
	SortPlaylistsByCreated     SortPlaylistsBy = "CREATED"
	SortPlaylistsByName        SortPlaylistsBy = "NAME"
)

var AllSortPlaylistsBy = []SortPlaylistsBy{
	SortPlaylistsByUnspecified,
	SortPlaylistsByCreated,
	SortPlaylistsByName,
}

func (e SortPlaylistsBy) IsValid() bool {
	switch e {
	case SortPlaylistsByUnspecified, SortPlaylistsByCreated, SortPlaylistsByName:
		return true
	}
	return false
}

func (e SortPlaylistsBy) String() string {
	return string(e)
}

func (e *SortPlaylistsBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortPlaylistsBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortPlaylistsBy", str)
	}
	return nil
}

func (e SortPlaylistsBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SortTracksBy string

const (
	SortTracksByUnspecified SortTracksBy = "UNSPECIFIED"
	SortTracksByTitle       SortTracksBy = "TITLE"
	SortTracksByRelease     SortTracksBy = "RELEASE"
	SortTracksByAdded       SortTracksBy = "ADDED"
)

var AllSortTracksBy = []SortTracksBy{
	SortTracksByUnspecified,
	SortTracksByTitle,
	SortTracksByRelease,
	SortTracksByAdded,
}

func (e SortTracksBy) IsValid() bool {
	switch e {
	case SortTracksByUnspecified, SortTracksByTitle, SortTracksByRelease, SortTracksByAdded:
		return true
	}
	return false
}

func (e SortTracksBy) String() string {
	return string(e)
}

func (e *SortTracksBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortTracksBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortTracksBy", str)
	}
	return nil
}

func (e SortTracksBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
