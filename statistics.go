package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Statistics struct {
	Wins             int `json:"wins"`
	Losses           int `json:"losses"`
	Ties             int `json:"ties"`
	MoneyLost        int `json:"money_lost"`
	HandsPlayed      int `json:"hands_played"`
	CorrectCounts    int `json:"correct_counts"`
	IncorrectCounts  int `json:"incorrect_counts"`
	CardCountStreak  int `json:"card_count_streak"`
	MaxCardCountStreak int `json:"max_card_count_streak"`
}

func LoadStatistics() (Statistics, error) {
	var stats Statistics
	file, err := ioutil.ReadFile("stats.json")
	if err != nil {
		if os.IsNotExist(err) {
			return stats, nil
		}
		return stats, err
	}
	err = json.Unmarshal(file, &stats)
	if err != nil {
		return stats, err
	}
	return stats, nil
}

func SaveStatistics(stats Statistics) error {
	file, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("stats.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}
