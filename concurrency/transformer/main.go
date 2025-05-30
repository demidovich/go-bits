// Преобразователь (декоратор) данных канала
// Прикидывается таким же каналом, с данными которого были осуществлены некоторые действия
//
// Может понадобиться если есть канал, на поток данных которого нельзя повлиять
// и есть часть приложения, принимающая этот канал, на которую так же нельзя повлиять
//
// channel -> transformation -> channel

package main

import "fmt"

func main() {
	jobs := fakeJobsGenerator(10)

	// Обработка данных канала без преобразователя

	// for v := range jobs {
	// 	fmt.Println(v)
	// }

	// Обработка данных канала с преобразованием

	var method = func(v string) string {
		return fmt.Sprintf("%s, transformed", v)
	}

	for v := range Transformer(jobs, method) {
		fmt.Println(v)
	}
}

// Не реализована проверка типов
func Transformer[T any](in <-chan T, method func(T) T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for v := range in {
			out <- method(v)
		}
	}()

	return out
}

func fakeJobsGenerator(count int) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for i := range count {
			out <- fmt.Sprintf("job %d", i)
		}
	}()

	return out
}
