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
