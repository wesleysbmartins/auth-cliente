package cache

import (
	"context"
	"fmt"
)

func Set(key string, value any) error {

	ctx := context.Background()

	conn := RedisConnection()

	err := conn.Set(ctx, key, value, 0).Err()

	defer conn.Close()

	HandleError(err)

	return err
}

func Get(key string) (string, error) {
	ctx := context.Background()

	conn := RedisConnection()
	defer conn.Close()

	val, err := conn.Get(ctx, key).Result()

	return val, err
}

func GetAll() ([]string, error) {
	ctx := context.Background()

	var values []string

	conn := RedisConnection()
	defer conn.Close()

	keys, err := conn.Keys(ctx, "*").Result()

	for _, key := range keys {
		value, _ := conn.Get(ctx, key).Result()
		values = append(values, value)
	}

	return values, err
}

func Flush() {
	ctx := context.Background()
	conn := RedisConnection()
	defer conn.Close()
	err := conn.FlushAll(ctx).Err()
	if err != nil {
		fmt.Println("Erro ao limpar o Redis:", err)
		return
	}

	fmt.Println("O Redis foi limpo com sucesso.")
}
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
