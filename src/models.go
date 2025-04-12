package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jerysun/NewsRefreshed/internal/database"
)

type User struct {
  ID        uuid.UUID `json:"id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  Name      string    `json:"name"`
  ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(user database.User) User {
  return User {
    ID:        user.ID,
    CreatedAt: user.CreatedAt,
    UpdatedAt: user.UpdatedAt,
    Name:      user.Name,
    ApiKey:    user.ApiKey,
  }
}

type Feed struct {
  ID            uuid.UUID   `json:"id"`
  CreatedAt     time.Time   `json:"created_at"`
  UpdatedAt     time.Time   `json:"updated_at"`
  Name          string      `json:"name"`
  Url           string      `json:"url"`
  UserID        uuid.UUID   `json:"user_id"`
  LastFetchedAt *time.Time `json:"last_fetched_at"`
}

func databaseFeedToFeed(feed database.Feed) Feed {
  return Feed {
    ID:             feed.ID,
    CreatedAt:      feed.CreatedAt,
    UpdatedAt:      feed.UpdatedAt,
    Name:           feed.Name,
    Url:            feed.Url,
    UserID:         feed.UserID,
    LastFetchedAt:  nullTimeToTimePtr(feed.LastFetchedAt),
  }
}

func databaseFeedsToFeeds(feeds []database.Feed) []Feed {
  goFeeds := []Feed{}
  for _, feed := range feeds {
    goFeeds = append(goFeeds, databaseFeedToFeed(feed))
  }
  
  return goFeeds
}

type FeedFollow struct {
  ID        uuid.UUID `json:"id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  UserID    uuid.UUID `json:"user_id"`
  FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(feedFollow database.FeedFollow) FeedFollow {
  return FeedFollow {
    ID:        feedFollow.ID,
    CreatedAt: feedFollow.CreatedAt,
    UpdatedAt: feedFollow.UpdatedAt,
    UserID:    feedFollow.UserID,
    FeedID:    feedFollow.FeedID,
  }
}

func databaseFeedFollowsToFeedFollows(feedFollows []database.FeedFollow) []FeedFollow {
  goFeedFollows := make([]FeedFollow, len(feedFollows))
  for i, feedFollow := range feedFollows {
    goFeedFollows[i] = databaseFeedFollowToFeedFollow(feedFollow)
  }
  return goFeedFollows
}

func nullTimeToTimePtr(t sql.NullTime) *time.Time {
  if t.Valid {
    return &t.Time
  }
  return nil
}