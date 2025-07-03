# go bits

Реализации структур данных и конкурентных паттернов. Репозиторий не несет никакой практической пользы и создан в рамках изучения языка.

### concurrency

* __fan_in__
* __fan_out__
* __filter__
* __semaphore__
* __tee__
* __transformer__

### keyvalue

Хранилище keyvalue с сегментацией на бакеты.

### msa

* __limiter__ - простой rate limiter на бакетах с пополнением раз в интервал
* __throttle__ - троттлер из книги "Облачный Go" с добавленными дженериками

### structures

* __list__ - двухсвязный список
* __stack_lifo_slice__ - стэк LIFO на слайсе
* __stack_fifo_slice__ - стэк FIFO на слайсе


