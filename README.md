# ChattingRoom
### Compile
```
mv ChattingRoom $GOPATH/src
cd $GOPATH
[server]
go build -o bin/ChattingRoom/server ChattingRoom/server/main

[client-linux]
go build -o bin/ChattingRoom/client ChattingRoom/client/main

[client-windows]
env GOOS=windows go build -o bin/chattingRoom/client.exe ChattingRoom/client/main
