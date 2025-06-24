package backoff

import (
	"math/rand/v2"
	"time"
)

// CalculateExponentialBackoff вычисляет экспоненциальную задержку с jitter.
// Параметры:
//   - attempt: номер попытки (начинается с 1)
//   - minDelay: минимальная задержка
//   - maxDelay: максимальная задержка
//
// Возвращает вычисленную задержку с учетом экспоненциального роста и случайного jitter.
func CalculateExponentialBackoff(attempt int, minDelay, maxDelay time.Duration) time.Duration {
	if attempt < 1 {
		attempt = 1 // 1-я попытка имела базовый minDelay (2^0 = 1)
	}

	// Экспоненциальный рост: baseDelay * 2^(attempt-1)
	delay := min(minDelay*time.Duration(1<<(attempt-1)), maxDelay)

	// Добавляем случайный джиттер (от 0.5 до 1.5 от delay)
	jitter := 0.5 + rand.Float64() // [0.5, 1.5)
	delay = time.Duration(float64(delay) * jitter)

	return delay
}
