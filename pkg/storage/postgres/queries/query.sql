-- name: GetTotalSortedByConsole :many
SELECT console, sum("public"."games"."sorted") AS "sum"
    FROM "public"."games"
    GROUP BY console
    ORDER BY console DESC
    LIMIT 10;
    
-- name: GetTotalSortedByGenre :many
SELECT genre, sum("public"."games"."sorted") AS "sum"
    FROM "public"."games"
    GROUP BY genre
    ORDER BY genre DESC
    LIMIT 10;