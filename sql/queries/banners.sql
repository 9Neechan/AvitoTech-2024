-- name: CreateBanner :one
INSERT INTO banners (feature, tag, json_content, is_active, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUserBanner :one
SELECT * FROM banners WHERE feature = $1 AND tag = $2;

-- name: GetAllBannersByFeature :many
SELECT * FROM banners WHERE feature = $1;

-- name: GetAllBannersByTag :many
SELECT * FROM banners WHERE tag = $1;

-- name: UpdateBanner :one
UPDATE banners
SET json_content = $1, is_active = $2, updated_at = NOW()
WHERE feature = $3 AND tag = $3
RETURNING *;

-- name: DeleteBanner :one
DELETE FROM banners
WHERE feature = $1 AND tag = $2
RETURNING *;

