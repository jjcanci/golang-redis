package main

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

// ping tests connectivity for redis (PONG should be returned)
func ping(c redis.Conn) error {
	// Send PING command to Redis
	pong, err := c.Do("PING")
	if err != nil {
		return err
	}

	// PING command returns a Redis "Simple String"
	// Use redis.String to convert the interface type to string
	s, err := redis.String(pong, err)
	if err != nil {
		return err
	}

	log.Printf("PING Response = %s\n", s)
	// Output: PONG

	return nil
}

// set executes the redis SET command
func set(c redis.Conn) error {
	log.Println("Set values")
	_, err := c.Do("SET", "Ejemplo:Favorite Movie", "Repo Man")
	if err != nil {
		return err
	}
	_, err = c.Do("SET", "Ejemplo:Release Year", 1984)
	if err != nil {
		return err
	}
	return nil
}

// get executes the redis GET command
func get(c redis.Conn) error {

	// Simple GET example with String helper
	key := "Ejemplo:Favorite Movie"
	s, err := redis.String(c.Do("GET", key))
	if err != nil {
		return (err)
	}
	log.Printf("%s = %s\n", key, s)

	// Simple GET example with Int helper
	key = "Ejemplo:Release Year"
	i, err := redis.Int(c.Do("GET", key))
	if err != nil {
		return (err)
	}
	log.Printf("%s = %d\n", key, i)

	// Example where GET returns no results
	key = "Nonexistent Key"
	s, err = redis.String(c.Do("GET", key))
	if err == redis.ErrNil {
		log.Printf("%s does not exist\n", key)
	} else if err != nil {
		return err
	} else {
		log.Printf("%s = %s\n", key, s)
	}

	return nil
}

// set executes the redis SET command
func hset(c redis.Conn) error {
	log.Println("HSet values")
	key := "Ejemplo:bar"
	_, err := c.Do("HSET", key, "1", "First bar")
	if err != nil {
		return err
	}

	_, err = c.Do("HSET", key, "2", "Second bar")
	if err != nil {
		return err
	}

	_, err = c.Do("HSET", key, "3", "Third Value")
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// Connect
	log.Println("Connecting to redis...")
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	log.Println("Redis connected!!")

	// Send PING command to Redis
	err = ping(c)
	if err != nil {
		log.Fatal(err)
	}

	// Set values
	err = set(c)
	if err != nil {
		log.Fatal(err)
	}

	// Get values
	err = get(c)
	if err != nil {
		log.Fatal(err)
	}

	// HSet values
	err = hset(c)
	if err != nil {
		log.Fatal(err)
	}
}
