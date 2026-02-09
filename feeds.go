package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func handleFollow(s *state, cmd command, user database.User) error {

	if len(cmd.args) < 1 {
		return fmt.Errorf("Usage: gator follow <feedurl>")
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	follow, err := s.db.CreateFeedFollow(context.Background(),
		database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return err
	}

	fmt.Printf("%s is now followed by %s\n", follow.FeedName, follow.UserName)

	return nil
}

func handleListFollowers(s *state, cmd command, user database.User) error {

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, follow := range follows {
		fmt.Printf("* %s (%s)\n", follow.Name, follow.Url)
	}

	return nil
}

func handleListFeeds(s *state, cmd command) error {

	feeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("* %s (%s) - %s\n", feed.Name, feed.Url, feed.UserName.String)
	}

	return nil
}

func handleAddFeed(s *state, cmd command, user database.User) error {

	if len(cmd.args) < 2 {
		return fmt.Errorf("Usage: gator addfeed <feedname> <feedurl>")
	}

	params := database.AddFeedParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.args[0], Url: cmd.args[1], UserID: user.ID}

	feed, err := s.db.AddFeed(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Println(feed)

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return err
	}

	fmt.Printf("%s is now followed by %s\n", follow.FeedName, follow.UserName)
	return nil
}

func handleUnfollow(s *state, cmd command, user database.User) error {

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.db.DeleteFollow(context.Background(), database.DeleteFollowParams{UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return err
	}

	return nil
}

func handleBrowse(s *state, cmd command) error {

	if len(cmd.args) < 1 {
		return fmt.Errorf("Usage: gator browse <limit>")
	}

	limit, err := strconv.Atoi(cmd.args[0])
	if err != nil {
		return err
	}

	posts, err := s.db.GetPosts(context.Background(), int32(limit))
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("%s (%s)\n", post.Title.String, post.Url)
	}

	return nil
}
