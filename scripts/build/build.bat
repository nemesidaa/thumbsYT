@echo off
go build -o server.exe ..\..\cmd\main.go
pause
go build -o client.exe ..\..\CLI\cmd\main.go