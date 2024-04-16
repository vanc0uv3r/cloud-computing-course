package main

import (
	"fmt"
	"strconv"

	// "log"
	"context"
	"net/http"

	"github.com/go-redis/redis/v8"
)

var (
	rClient *redis.Client
	ctx     = context.Background()
)

func getCurrentState() (int, error) {
	val, err := rClient.Get(ctx, "var").Result()
	if err != nil {
		return 0, err
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	return intVal, nil
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	res, err := rClient.BgSave(ctx).Result()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}
	w.Write([]byte(res))
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	currentVal, err := getCurrentState()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}

	err = rClient.Set(ctx, "var", currentVal+1, 0).Err()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error on adding: %s", err.Error())))
		return
	}

	w.Write([]byte("Done!\n"))
}

func getCurrentHandler(w http.ResponseWriter, r *http.Request) {
	val, err := getCurrentState()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
	} else {
		w.Write([]byte(fmt.Sprintf("Current value is: %d", val)))
	}
}

func subHandler(w http.ResponseWriter, r *http.Request) {
	currentVal, err := getCurrentState()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}

	err = rClient.Set(ctx, "var", currentVal-1, 0).Err()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error on sub: %s", err.Error())))
		return
	}

	w.Write([]byte("Done!\n"))
}

func main() {
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/sub", subHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/get-current", getCurrentHandler)

	rClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	fmt.Println("Start Working")
	val, err := rClient.Exists(ctx, "var").Result()
	if err != nil {
		panic(err.Error())
	}
	if val == 0 {
		err = rClient.Set(ctx, "var", 0, 0).Err()
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Default value have set")
	}

	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err.Error())
	}

}
