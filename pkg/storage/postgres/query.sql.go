// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: query.sql

package postgres

import (
	"context"
)

const getTop10Games = `-- name: GetTop10Games :many
SELECT
    "title",
    "image_url",
    sum("public"."games"."sorted") AS "sum"
FROM
    "public"."games"
GROUP BY
    "title",
    "image_url"
ORDER BY
    "sum" DESC,
    "title" ASC
LIMIT 10
`

type GetTop10GamesRow struct {
	Title    string
	ImageUrl string
	Sum      int64
}

func (q *Queries) GetTop10Games(ctx context.Context) ([]GetTop10GamesRow, error) {
	rows, err := q.db.Query(ctx, getTop10Games)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTop10GamesRow
	for rows.Next() {
		var i GetTop10GamesRow
		if err := rows.Scan(&i.Title, &i.ImageUrl, &i.Sum); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTotalGames = `-- name: GetTotalGames :many
SELECT
    count(*) AS "count"
FROM
    "public"."games"
`

func (q *Queries) GetTotalGames(ctx context.Context) ([]int64, error) {
	rows, err := q.db.Query(ctx, getTotalGames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var count int64
		if err := rows.Scan(&count); err != nil {
			return nil, err
		}
		items = append(items, count)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTotalGamesByConsole = `-- name: GetTotalGamesByConsole :many
SELECT
    console,
    count(*) AS "sum"
FROM
    "public"."games"
GROUP BY
    "console"
ORDER BY
    "sum" DESC
`

type GetTotalGamesByConsoleRow struct {
	Console string
	Sum     int64
}

func (q *Queries) GetTotalGamesByConsole(ctx context.Context) ([]GetTotalGamesByConsoleRow, error) {
	rows, err := q.db.Query(ctx, getTotalGamesByConsole)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTotalGamesByConsoleRow
	for rows.Next() {
		var i GetTotalGamesByConsoleRow
		if err := rows.Scan(&i.Console, &i.Sum); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTotalSortedByConsole = `-- name: GetTotalSortedByConsole :many
SELECT
    "console",
    sum("public"."games"."sorted") AS "sum"
FROM
    "public"."games"
GROUP BY
    "console"
ORDER BY
    "sum" DESC,
    "console" ASC
LIMIT 10
`

type GetTotalSortedByConsoleRow struct {
	Console string
	Sum     int64
}

func (q *Queries) GetTotalSortedByConsole(ctx context.Context) ([]GetTotalSortedByConsoleRow, error) {
	rows, err := q.db.Query(ctx, getTotalSortedByConsole)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTotalSortedByConsoleRow
	for rows.Next() {
		var i GetTotalSortedByConsoleRow
		if err := rows.Scan(&i.Console, &i.Sum); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTotalSortedByGenre = `-- name: GetTotalSortedByGenre :many
SELECT
    "genre",
    sum("public"."games"."sorted") AS "sum"
FROM
    "public"."games"
GROUP BY
    "genre"
ORDER BY
    "sum" DESC,
    "genre" ASC
LIMIT 10
`

type GetTotalSortedByGenreRow struct {
	Genre string
	Sum   int64
}

func (q *Queries) GetTotalSortedByGenre(ctx context.Context) ([]GetTotalSortedByGenreRow, error) {
	rows, err := q.db.Query(ctx, getTotalSortedByGenre)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTotalSortedByGenreRow
	for rows.Next() {
		var i GetTotalSortedByGenreRow
		if err := rows.Scan(&i.Genre, &i.Sum); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
