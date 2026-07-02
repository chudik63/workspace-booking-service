package workspace

import "time"

type workspaceItem struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	Lat         float64   `db:"lat"`
	Lon         float64   `db:"lon"`
	FullAddress string    `db:"full_address"`
	Type        string    `db:"type"`
	Status      string    `db:"status"`
	Capacity    uint32    `db:"capacity"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	IsDeleted   bool      `db:"is_deleted"`
}
