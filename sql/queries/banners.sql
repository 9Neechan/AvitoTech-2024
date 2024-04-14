-- name: CreateBanner :one
with create_banner AS (
  INSERT into banner (feature_id, is_active, "content") VALUES ($2, $3, $4) RETURNING "id", feature_id
),
create_banner_relation as (
  INSERT into banner_relation (banner_id, feature_id, tag_id)
  SELECT cb.id as banner_id, cb.feature_id as feature_id, UNNEST($1::int[]) as tag_id 
  FROM create_banner AS cb
  RETURNING *
)

SELECT * FROM create_banner;

-- name: GetUserBanner :one
with find_banner as (
  SELECT banner_id, banner_relation.tag_id, banner_relation.feature_id 
  FROM banner_relation 
  WHERE banner_relation.feature_id=$2 AND banner_relation.tag_id=$1
)

SELECT
  b.id,
  b.feature_id,
  (SELECT ARRAY_AGG(tag_id) FROM banner_relation AS br WHERE br.banner_id = b.id) as tag_ids,
  b.content,
  b.is_active,
  b.created_at,
  b.updated_at
FROM banner as b JOIN find_banner as fb ON (b.id = fb.banner_id);

-- name: DeleteBanner :one
DELETE from banner 
WHERE "id" = $1
RETURNING *;

-- name: GetBannerByID :one
SELECT
		b.id,
		b.feature_id,
		(SELECT ARRAY_AGG(tag_id) FROM banner_relation AS br WHERE br.banner_id = b.id) as tag_ids,
		b.content,
		b.is_active,
		b.created_at,
		b.updated_at
FROM banner as b
WHERE b.id = $1;

-- name: InsertTags :exec
INSERT INTO banner_relation (banner_id, feature_id, tag_id) 
SELECT $1, $2, UNNEST($3::int[]);

-- name: DeleteTags :exec
DELETE from banner_relation 
WHERE banner_id=$1 AND feature_id=$2;

-- name: UpdateBanner :exec
UPDATE banner 
SET feature_id=$2, is_active=$3, content=$4, updated_at=CURRENT_TIMESTAMP
WHERE id=$1;

-- name: GetBannerListByFeatureId :many
SELECT
	b.id,
	b.feature_id,
	CAST (( 
    SELECT ARRAY_AGG(tag_id::INT) 
    FROM banner_relation AS br 
    WHERE br.banner_id = b.id
  ) AS INTEGER[]) as tag_ids,
	b.content,
	b.is_active,
	b.created_at,
	b.updated_at
FROM banner as b
WHERE b.feature_id = $1
ORDER BY b.created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetBannerListByTag :many
SELECT 	
  b.id,
	b.feature_id,
	CAST (( 
    SELECT ARRAY_AGG(tag_id::INT) 
    FROM banner_relation AS br 
    WHERE br.banner_id = b.id AND br.tag_id = $1
  ) AS INTEGER[]) as tag_ids,
	b.content,
	b.is_active,
	b.created_at,
	b.updated_at
FROM banner as b
ORDER BY b.created_at DESC
LIMIT $2 OFFSET $3;



