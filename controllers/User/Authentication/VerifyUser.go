package authentication

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

var ctx = context.Background()

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func StoreOTP(userID string, otp string) error {
	key := "otp:" + userID

	err := redisClient.Set(ctx, key, otp, 5*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetOTP(userID string) (string, error) {
	key := "otp:" + userID

	otp, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("OTP expired or not found")
		}
		return "", err
	}

	return otp, nil
}

func DeleteOTP(userID string) error {
	key := "otp:" + userID
	return redisClient.Del(ctx, key).Err()
}

func VerifyOTP(userID, userOTP string) (bool, error) {
	storedOTP, err := GetOTP(userID)
	if err != nil {
		return false, err
	}

	if storedOTP != userOTP {
		return false, fmt.Errorf("invalid otp")
	}

	if err := DeleteOTP(userID); err != nil {
		return false, fmt.Errorf("failed to delete otp after verification")
	}

	return true, nil
}
