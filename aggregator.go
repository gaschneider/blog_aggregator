package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.params) != 1 {
		return fmt.Errorf("aggregator handler expects single parameter for time between reqs")
	}

	timeBetweenRequests := cmd.params[0]

	timeDuration, err := time.ParseDuration(timeBetweenRequests)
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)
	fmt.Println()

	ticker := time.NewTicker(timeDuration)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
		fmt.Println()
		fmt.Println()
	}
}
