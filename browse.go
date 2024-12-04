package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gaschneider/blog_aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.params) > 1 {
		return fmt.Errorf("browse handler expects single optional parameter for limit of posts")
	}

	var limit int32
	limit = 2

	if len(cmd.params) == 1 {
		num, err := strconv.ParseInt(cmd.params[0], 10, 32)
		if err != nil {
			return fmt.Errorf("browse handler expects a number as parameter")

		}

		limit = int32(num)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})

	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println()
		fmt.Println()
		fmt.Printf("%v - %v\n", post.Title, post.PublishedAt)
		fmt.Printf("%v\n\n", post.Url)
		fmt.Printf("%v\n", post.Description)
		fmt.Println("-------------------------------------")
	}

	return nil
}
