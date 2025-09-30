-- name: DeleteFeedFollowByUserAndFeed :exec
DELETE FROM feed_follows 
WHERE user = $1 AND feeds.id = $2;
