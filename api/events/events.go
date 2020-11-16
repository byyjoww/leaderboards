package events

import (
	"encoding/json"
	"fmt"
	"net/http"

	"git.topfreegames.com/hackathon/leaderboards/leaderboards"
)

var eventsMap map[string]func(ctx Context) = map[string]func(ctx Context){
	"matchResult": handleMatchResult,
}

// Context is the full struct for the event parent.
type Context struct {
	Type     string `json:"type"`
	LevelID  string `json:"levelid"`
	PlayerID string `json:"playerid"`
	Score    int    `json:"score"`
}

// Listener listens for api events.
func Listener(w http.ResponseWriter, r *http.Request) {
	var ctx Context
	if err := json.NewDecoder(r.Body).Decode(&ctx); err != nil {
		fmt.Println(err)
	}

	// Print non-indented json
	// ctxString, _ := json.Marshal(ctx)
	// fmt.Println(string(ctxString))

	// Print indented json
	// ctxIndentedString, _ := json.MarshalIndent(ctx, "", "\t")
	// fmt.Println(string(ctxIndentedString))

	eventType := ctx.Type
	fmt.Println("context type: ", ctx.Type)
	value, isValid := eventsMap[eventType]
	if !isValid {
		fmt.Println("invalid event type")
		return
	}

	value(ctx)

	w.Header().Set("Content-Type", "application/json")
	response := "Request Succesfull"
	w.Write([]byte(response))
}

func handleMatchResult(ctx Context) {
	fmt.Println("Received the following score: ", ctx.Score)
	leaderboards.IncrementPlayerScore(ctx.LevelID, ctx.PlayerID, ctx.Score)
}
