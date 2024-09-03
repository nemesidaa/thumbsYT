@echo off

cd /d ..\proto\service
pause
protoc --go_out=..\gen\service --go_opt=paths=source_relative --go-grpc_out=..\gen\service --go-grpc_opt=paths=source_relative service.proto

pause
