# Поиск минимального MSU

Запускать с помощью докера. 

Пример запуска:

Сначала выполняем команду: docker build . -t dock_mtu
```bash
$ docker build . -t dock_mtu
Sending build context to Docker daemon  2.155MB
Step 1/12 : FROM golang:1.18 as builder
 ---> c6740b5c6ec7
Step 2/12 : RUN apt-get update
 ---> Using cache
 ---> c7adf3359eb4
Step 3/12 : WORKDIR /project
 ---> Using cache
 ---> 84bb1b031033
Step 4/12 : COPY Makefile /project
 ---> Using cache
 ---> 44a69ea55fd1
Step 5/12 : COPY . /project
 ---> Using cache
 ---> 6b44e6107108
Step 6/12 : RUN make bin/server
 ---> Using cache
 ---> 99dff95003d7
Step 7/12 : FROM ubuntu:20.04
 ---> 817578334b4d
Step 8/12 : RUN apt update && yes | apt upgrade
 ---> Using cache
 ---> 0b279abbe3ce
Step 9/12 : RUN yes | apt install python3 iputils-ping
 ---> Using cache
 ---> 84e8dc452e62
Step 10/12 : WORKDIR /
 ---> Using cache
 ---> b4c8b75770d3
Step 11/12 : COPY --from=builder /project/bin/server /bin/server
 ---> Using cache
 ---> b9cad049e278
Step 12/12 : ENTRYPOINT ["/bin/server"]
 ---> Using cache
 ---> c2cdd0d23ead
Successfully built c2cdd0d23ead
Successfully tagged dock_mtu:latest
```
Затем команду: run dock_mtu --host HOST

```bash
$ docker run dock_mtu --host google.com
MSU is 96
```

Так же можно запустить такой командой: 
```
go run prog.go --host ya.ru
```

Хост по дефолту стоит ya.ru.

# Решение
Сначала бин. подъемами ищем верхнюю границу для нашего MTU (пингуем начиная с размера пакета равного единицы, пока пинг не начнет выдавать ошибку). Далее бинарным поиском ищем самый большой размер пакета на котором пинг не возвращает ошибки. Устанавливаем флаг запрещающий фрагментацию.

