// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package postgres

import ()

type Game struct {
	ID       int32
	Title    string
	Genre    string
	Console  string
	FileUrl  string
	ImageUrl string
	Sorted   int32
	Active   bool
}
