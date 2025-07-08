# base62

Конвертация `int` в `base62`. Применяется для получения сокращенного токена на основе уникальной числовой последовательности.

```go
v := 9223372036854775807 // Max int
b := base62.EncodeInt(v)

fmt.Println(b) // K9VIxAiFIwH

fmt.Println(base62.DecodeInt(b)) // 9223372036854775807
```