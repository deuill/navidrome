package domain

type BaseRepository interface {
	NewId(fields ...string) string
	CountAll() (int, error)
	Exists(id string) (bool, error)
}

type QueryOptions struct {
	SortBy string
	Alpha  bool
	Desc   bool
	Offset int
	Size   int
}