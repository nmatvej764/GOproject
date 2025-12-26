@echo off
cd /d %~dp0..\services\api-service
go run cmd\api\main.go
pause
