package main

import (
	"fmt"

	"git.topfreegames.com/hackathon/leaderboards/api"
	"git.topfreegames.com/hackathon/leaderboards/input"
	"git.topfreegames.com/hackathon/leaderboards/leaderboards"
)

func main() {
	fmt.Println("initiating leaderboards")
	go input.WaitForInput()

	leaderboards.LoadLeaderboardData()
	api.InitializeAPIRoutes()
}
