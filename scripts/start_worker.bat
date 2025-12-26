@echo off
cd /d %~dp0..\services\worker-service
go run cmd\worker\main.go
pause
