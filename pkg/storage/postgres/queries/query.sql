-- name: GetTotalSortedByConsole :many
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
LIMIT 10;

-- name: GetTotalSortedByGenre :many
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
LIMIT 10;

-- name: GetTop10Games :many
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
LIMIT 10;

-- name: GetTotalGames :many
SELECT
    count(*) AS "count"
FROM
    "public"."games";

-- name: GetTotalGamesByConsole :many
SELECT
    console,
    count(*) AS "sum"
FROM
    "public"."games"
GROUP BY
    "console"
ORDER BY
    "sum" DESC;

-- name: GetThreeRandomGames :many
SELECT ss.* FROM (SELECT DISTINCT id,title,genre,file_url,image_url FROM games) ss ORDER BY random() LIMIT 3;

-- name: UpdateSortedGame :exec
UPDATE games 
      SET 
      sorted = sorted + 1,
      lastupdatedat = now()
   WHERE id = $1;