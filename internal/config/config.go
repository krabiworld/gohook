package config

import "os"

type Config struct {
	Address      string
	Secret       string
	HappyEmoji   string
	SuccessEmoji string
	FailureEmoji string
}

var cfg *Config

func Init() {
	cfg = &Config{
		Address:      env("ADDR", ":8080"),
		Secret:       env("SECRET"),
		HappyEmoji:   env("HAPPY_EMOJI", "🥳"),
		SuccessEmoji: env("SUCCESS_EMOJI", "🥖"),
		FailureEmoji: env("FAILURE_EMOJI", "😭"),
	}
}

func env(key string, defaultValue ...string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}

func Get() *Config {
	return cfg
}
