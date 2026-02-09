package main

import (
	"database/sql"
	"gator/internal/database"

	_ "github.com/lib/pq"
)
import (
	"fmt"
	"gator/internal/config"
	"os"
)

func main() {

	args := os.Args[1:]

	cfg, err := config.Read()
	if err != nil {
		handleError(true, err)
		return
	}
	cmd, err := getCommand(args)
	if err != nil {
		handleError(true, err)
		return
	}

	s := state{cfg: &cfg}
	commandMap, err := buildCmdMap()

	cmdFnc, ok := commandMap.handlers[cmd.name]
	if !ok {
		handleError(true, fmt.Errorf("unknown command: %s", args[0]))
		return
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		handleError(false, err)
		return
	}
	defer db.Close()
	s.db = database.New(db)

	err = cmdFnc(&s, cmd)
	if err != nil {
		handleError(false, err)
	}

	return
}

func handleError(usage bool, err error) {
	if err != nil {
		fmt.Println(err)
		if usage {
			printUsage()
		}
		os.Exit(1)
		return
	}
}

func printUsage() {
	fmt.Println("Usage: gator [command]")
}

func getCommand(args []string) (command, error) {
	if len(args) < 1 {
		return command{}, fmt.Errorf("no command specified")
	}

	return command{name: args[0], args: args[1:]}, nil
}

func buildCmdMap() (commands, error) {
	cmdMap := newCommands()

	err := cmdMap.register("login", handlerLogin)
	if err != nil {
		return commands{}, err
	}

	err = cmdMap.register("register", handlerRegister)
	if err != nil {
		return commands{}, err
	}

	err = cmdMap.register("reset", handlerReset)
	if err != nil {
		return commands{}, err
	}

	err = cmdMap.register("users", handlerUsers)
	if err != nil {
		return commands{}, err
	}

	err = cmdMap.register("agg", handleAgg)
	if err != nil {
		return commands{}, err
	}

	err = cmdMap.register("addfeed", middlewareLoggedIn(handleAddFeed))
	if err != nil {
		return commands{}, err
	}

	err = cmdMap.register("feeds", handleListFeeds)
	if err != nil {
		return commands{}, err
	}

	err = cmdMap.register("follow", middlewareLoggedIn(handleFollow))
	if err != nil {
		return commands{}, err
	}

	err = cmdMap.register("following", middlewareLoggedIn(handleListFollowers))
	if err != nil {
		return commands{}, err
	}

	err = cmdMap.register("unfollow", middlewareLoggedIn(handleUnfollow))
	if err != nil {
		return commands{}, err
	}

	err = cmdMap.register("browse", handleBrowse)
	if err != nil {
		return commands{}, err
	}

	return cmdMap, nil
}
