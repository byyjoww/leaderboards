package leaderboards

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
)

var (
	globalLeaderboardScores []Leaderboard
)

// PlayerData contains the player id and score for a particular level.
type PlayerData struct {
	PlayerID string
	Score    int
}

// Leaderboard contains the level ID and an array of player scores for that level.
type Leaderboard struct {
	LevelID string
	Scores  []PlayerData
}

// LoadLeaderboardData will load persistent scores from json file.
func LoadLeaderboardData() {
	data, err := ioutil.ReadFile("leaderboardsData.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Loading leaderboard data:  \n", string(data))
	err = json.Unmarshal(data, &globalLeaderboardScores)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// SaveLeaderboardData will load persistent scores from json file.
func SaveLeaderboardData() {
	data, err := json.MarshalIndent(globalLeaderboardScores, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = ioutil.WriteFile("leaderboardsData.json", data, 0644)

	fmt.Println("Leaderboards saved to json file.")
}

// IncrementPlayerScore will increase the player's persistent score.
func IncrementPlayerScore(_levelID string, _playerID string, _scoreAmount int) {

	level, err := getLeaderboardByID(_levelID)
	if err != nil {
		panic("invalid level")
	}

	player, err := getPlayerByID(level.Scores, _playerID)
	if err != nil {
		player = createNewPlayer(&level.Scores, _playerID)
	}

	addPlayerLevelScore(player, _scoreAmount)
	// replacePlayerLevelScore(player, _scoreAmount)

	sortLeaderboardByScore(level)
	SaveLeaderboardData()
}

func getPlayerByID(_leaderboard []PlayerData, _playerID string) (*PlayerData, error) {

	for i := 0; i < len(_leaderboard); i++ {
		if _leaderboard[i].PlayerID == _playerID {
			return &_leaderboard[i], nil
		}
	}

	return &PlayerData{}, errors.New("error: player not found")
}

func getLeaderboardByID(_levelID string) (*Leaderboard, error) {

	for i := 0; i < len(globalLeaderboardScores); i++ {
		if globalLeaderboardScores[i].LevelID == _levelID {
			return &globalLeaderboardScores[i], nil
		}
	}

	// create new leaderboard for id
	fmt.Println("Leaderboard doesnt exist, creating new leaderboard with id:", _levelID)
	createLeaderboard(_levelID)

	for i := 0; i < len(globalLeaderboardScores); i++ {
		if globalLeaderboardScores[i].LevelID == _levelID {
			return &globalLeaderboardScores[i], nil
		}
	}

	return &Leaderboard{}, errors.New("error: level not found")
}

func getLeaderboardIndex(level *Leaderboard) (int, error) {
	for k, v := range globalLeaderboardScores {
		if level.LevelID == v.LevelID {
			return k, nil
		}
	}
	return -1, errors.New("leaderboard not found")
}

func createNewPlayer(_level *[]PlayerData, _playerID string) *PlayerData {
	fmt.Println("Creating new player with id:", _playerID)

	newPlayer := PlayerData{
		PlayerID: _playerID,
		Score:    0,
	}

	*_level = append(*_level, newPlayer)

	for i := 0; i < len(*_level); i++ {
		if (*_level)[i].PlayerID == _playerID {
			return &(*_level)[i]
		}
	}

	panic("invalid player")
}

func addPlayerLevelScore(_player *PlayerData, _scoreAmount int) {
	_player.Score += _scoreAmount
	fmt.Println("Adding score to existing player: ", _player)
}

func replacePlayerLevelScore(_player *PlayerData, _scoreAmount int) {
	if _scoreAmount > _player.Score {
		_player.Score = _scoreAmount
	}

	fmt.Println("Substituting score for higher value: ", _player)
}

func sortLeaderboardByScore(level *Leaderboard) {
	sort.SliceStable(level.Scores, func(i, j int) bool {
		return level.Scores[i].Score > level.Scores[j].Score
	})
}

func createLeaderboard(_levelID string) {
	newLeaderboard := Leaderboard{
		LevelID: _levelID,
		Scores:  []PlayerData{},
	}
	globalLeaderboardScores = append(globalLeaderboardScores, newLeaderboard)
}

func resetLeaderboard(level *Leaderboard) {

	index, err := getLeaderboardIndex(level)
	if err != nil {
		fmt.Println(err)
		return
	}

	globalLeaderboardScores[index] = Leaderboard{
		LevelID: level.LevelID,
		Scores:  []PlayerData{},
	}
}
