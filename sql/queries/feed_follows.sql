-- name: CreateFeedFollow :one
WITH inserted_feed_follows AS (
	INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING *
)
SELECT
	inserted_feed_follows.*,
	feeds.name AS feed_name,
	users.name AS user_name
FROM inserted_feed_follows
JOIN users ON users.id = inserted_feed_follows.user_id
JOIN feeds ON feeds.id = inserted_feed_follows.feed_id;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = $1;

-- name: GetFeedFollowsForUser :many
SELECT
	ff.*,
	u.name AS user_name,
	f.name AS feed_name
FROM feed_follows ff
JOIN users u ON u.id = ff.user_id
JOIN feeds f ON f.id = ff.feed_id
WHERE ff.user_id = $1
ORDER BY ff.created_at DESC;
