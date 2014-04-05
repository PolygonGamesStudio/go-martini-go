go-martini-go
=============

**To run server type:**

**Setting $GOPATH**
```
$ export GOPATH=$HOME/go

$ export PATH=$PATH:$GOPATH/bin
```

**Download dependens**
```
$ go get github.com/codegangsta/martini
$ go get github.com/coopernurse/gorp
$ go get github.com/lib/pq
```

**run**
```
$ su postgres
$ pg_ctl start
```
отдельно:
```
$ go run src/server.go src/view.go src/model.go
```

Настройка БД
=============

*Все команды выполнять от суперпользователя:*
```
$ sudo su
```
Полный текст смотри в http://hexvolt.blogspot.ru/2012/11/postgresql-91-ubuntu-1204.html

Для работы с постгресом используем pgadmin3
```
$ sudo apt-get install pgadmin3
```

не забываем добавить поле isfavorite в userPlacesMTM

**Docs**
https://github.com/go-martini/martini
https://github.com/coopernurse/gorp
https://github.com/PuerkitoBio/martini-api-example/blob/master/api.go
https://github.com/PuerkitoBio/martini-api-example/blob/master/server.go

======
FAQ:
при изменении структуры таблиц надо удалять все данные, потому что у них в новом поле стоит nil, а наш gorp на этом валится нах

вставлять удобно можно во все таблицы кроме MTM - там делаем скрипт инсерта
в остальных можно сделать выборку и менять прям в таблице
