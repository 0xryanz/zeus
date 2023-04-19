module zeus

go 1.14

require (
	github.com/fsouza/go-dockerclient v1.6.5
	github.com/gorilla/websocket v1.4.2
	github.com/satori/go.uuid v1.2.0
	golang.org/x/net v0.0.0-20190501004415-9ce7a6920f09
	zcache v0.0.0
	zorm v0.0.0
	zweb v0.0.0
)

replace zweb => ./zweb

replace zorm => ./zorm

replace zcache => ./zcache
