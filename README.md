# guestbook.go

This just a another guestbook application to learn myself go. This is written as
a single page application since the exercise of this project is to focus on go,
and server side rendering is totally ok for this type of application.

In the future I want to extend this with other features, using it as a
playground for myself.

The following open source technologies are used:

* [redis/redis](https://github.com/redis/redis)
* [gin-gonic/gin](https://github.com/gin-gonic/gin)
* [Masterminds/sprig](https://github.com/Masterminds/sprig)
* [fomantic/Fomantic-UI](https://github.com/fomantic/Fomantic-UI)

## Development

### Run the tests:

```
go test ./src/...
```

If you want to continuously run the test while developing concider using
[`gow`](https://github.com/mitranim/gow):

```
gow test ./src/...
```

### Start the server:

```
go run main.go
```

If you wangt to continuously restart the server while developing consider using
[`air`](https://github.com/cosmtrek/air):

```
air
```
