package domain

type Playlist struct {
	Name string
}

type Id interface {
	Text() string
}

type AlbumId Id

type Album interface {
	Id() AlbumId
	Get() album
}

type Publisher struct {
	Name string
}

type TrackId Id

type Track interface {
	Id() TrackId
	Get() track
}

type Artist struct {
	Name string
}
