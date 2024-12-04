package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gaschneider/blog_aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	state, err := NewState()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", state.config.DBUrl)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)
	state.SetDB(dbQueries)

	cmds := NewCommands()
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

	params, err := getParams()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	commandToRun := command{}
	commandToRun.name = params[0]
	commandToRun.params = params[1:]

	err = cmds.run(&state, commandToRun)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	os.Exit(0)
}
