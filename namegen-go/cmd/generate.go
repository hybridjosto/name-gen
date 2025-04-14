package cmd

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type NameCategory struct {
	Prefixes map[string]string `json:"prefixes"`
	Suffixes map[string]string `json:"suffixes"`
}

type NameData struct {
	MaleNames   NameCategory `json:"male_names"`
	FemaleNames NameCategory `json:"female_names"`
	LastNames   NameCategory `json:"last_names"`
}

var gender string

func generateName(gender string) string {
	data, err := os.ReadFile("name-options.json")
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	var nameData NameData
	if err := json.Unmarshal(data, &nameData); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomValueFromMap := func(m map[string]string) string {
		leftRoll := r.Intn(6) + 1
		rightRoll := r.Intn(6) + 1
		finalRoll := strconv.Itoa(leftRoll) + strconv.Itoa(rightRoll)
		return m[finalRoll]
	}

	var selected NameCategory
	switch gender {
	case "male":
		selected = nameData.MaleNames
	case "female":
		selected = nameData.FemaleNames
	default:
		log.Fatalf("Unsupported gender: %s", gender)
	}

	first := randomValueFromMap(selected.Prefixes) + randomValueFromMap(selected.Suffixes)
	last := randomValueFromMap(nameData.LastNames.Prefixes) + randomValueFromMap(nameData.LastNames.Suffixes)

	return first + " " + last
}
