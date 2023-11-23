package observer

import (
	"context"
	"fmt"

	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisState struct {
	redis *redis.Client
}

var playersIdIncrement = 1

func NewRedisState() *RedisState {

	return &RedisState{
		redis.NewClient(&redis.Options{
			// Addr: "redis:6379",
			Addr: "localhost:6380",
			DB:   0,
		})}
}

func (s *RedisState) RedisDeleteAllPlayers() {
	s.redis.Del(context.Background(), "players")
}

func (s *RedisState) RedisSavePlayers(players []*Player) {
	fmt.Println("Saving players to Redis")
	for _, p := range players {
		p.ID = playersIdIncrement

		playerString, err := json.Marshal(p)
		if err != nil {
			log.Println(err)
			continue
		}

		err = s.redis.RPush(context.TODO(), "players", string(playerString)).Err()
		if err != nil {
			log.Println(err)
			continue
		}

		playersIdIncrement++
		log.Println("Saved players to Redis")
	}
}


func (s *RedisState) RedisListPlayers() []Player {
	playerStrings := s.redis.LRange(context.TODO(), "players", 0, -1).Val()

	players := []Player{}

	for _, p := range playerStrings {
		player := Player{}
		err := json.Unmarshal([]byte(p), &player)

		if err != nil {
			log.Println(err)
			return players
		}

		players = append(players, player)
	}

	return players
}
