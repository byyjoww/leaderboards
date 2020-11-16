package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"git.topfreegames.com/hackathon/leaderboards/leaderboards"
)

// WaitForInput listens to input to the console. Use as goroutine.
func WaitForInput() {
	// Create scanner to read input in real time
	scanner := bufio.NewScanner(os.Stdin)

	// Initial message
	fmt.Println("Please type in your input:")

	// Start loop to listen for input from client
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println("Input received from server:", input)

		checkForInput(input)
	}
}

func checkForInput(input string) {
	// Create new leaderboard with levelID
	if strings.HasPrefix(input, "leaderboard create ") {
		indexString := strings.TrimPrefix(input, "leaderboard create ")
		leaderboards.CreateLeaderboard(indexString)
		return
	}

	// Resets all leaderboards. Use with caution!
	if strings.HasPrefix(input, "leaderboard reset all") {
		leaderboards.ResetAllLeaderboards()
		return
	}

	// Reset a specific leaderboard based on levelID
	if strings.HasPrefix(input, "leaderboard reset ") {
		indexString := strings.TrimPrefix(input, "leaderboard reset ")
		leaderboards.ResetLeaderboard(indexString)
		return
	}

	// Removes all leaderboards. Use with caution!
	if strings.HasPrefix(input, "leaderboard remove all") {
		leaderboards.RemoveAllLeaderboards()
		return
	}

	// Remove a specific leaderboard based on levelID
	if strings.HasPrefix(input, "leaderboard remove ") {
		indexString := strings.TrimPrefix(input, "leaderboard remove ")
		leaderboards.RemoveLeaderboard(indexString)
		return
	}

	fmt.Println("input not recognized.")
}
