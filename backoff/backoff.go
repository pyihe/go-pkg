package backoff

import (
	"time"

	"github.com/pyihe/go-pkg/rands"
)

var defaultConfig = Config{
	BaseDelay:  1 * time.Second,
	MaxDelay:   120 * time.Second,
	Multiplier: 1.6,
	Jitter:     0.2,
}

type Config struct {
	initial    bool
	BaseDelay  time.Duration
	MaxDelay   time.Duration
	Multiplier float64
	Jitter     float64
}

func NewConfig() *Config {
	return &Config{initial: true}
}

// Get 获取延时
func Get(config *Config, retries int) time.Duration {
	if config == nil {
		config = &defaultConfig
	} else {
		if !config.initial {
			panic("Config must be initialized by NewConfig()")
		}
	}
	if retries <= 0 {
		return config.BaseDelay
	}
	backoff, max := float64(config.BaseDelay), float64(config.MaxDelay)
	for backoff < max && retries > 0 {
		backoff *= config.Multiplier
		retries--
	}
	if backoff > max {
		backoff = max
	}

	backoff *= 1 + config.Jitter*(rands.Float64()*2-1)
	if backoff < 0 {
		return 0
	}
	return time.Duration(backoff)
}
