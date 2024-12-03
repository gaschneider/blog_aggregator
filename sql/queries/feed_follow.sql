-- name: CreateFeedFollow :one
WITH insert_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)

SELECT i.*, u.name AS user_name, f.name AS feed_name
    FROM insert_feed_follow i
    INNER JOIN users u
    ON u.id = i.user_id
    INNER JOIN feeds f
    ON f.id = i.feed_id
    LIMIT 1;

-- name: GetFeedFollowsForUser :many
SELECT ff.*, u.name AS user_name, f.name AS feed_name
    FROM feed_follows ff
    INNER JOIN users u
    ON u.id = ff.user_id
    INNER JOIN feeds f
    ON f.id = ff.feed_id
    WHERE ff.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE user_id = $1 AND feed_id = $2;