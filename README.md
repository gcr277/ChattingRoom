# ChattingRoom
### Terminal 1
```
mv ChattingRoom $GOPATH/src
cd $GOPATH
go build -o bin/ChattingRoom/server ChattingRoom/server/main
go build -o bin/ChattingRoom/client ChattingRoom/client/main
cd $GOPATH/bin/ChattingRoom
./server
```
### Terminal 2
```
cd $GOPATH/bin/ChattingRoom
./client
```
