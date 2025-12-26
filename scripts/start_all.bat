@echo off
echo ===============================
echo Starting ALL services...
echo ===============================

REM 1) Start infrastructure (Kafka, Redis, UI)
echo [1/3] Starting infrastructure...
cd /d %~dp0..\deploy
docker compose up -d

REM Small wait so Kafka/Redis have time to start
echo Waiting 5 seconds...
timeout /t 5 /nobreak >nul

REM 2) Start API service in new window
echo [2/3] Starting api-service...
start "api-service" cmd /k "cd /d %~dp0..\services\api-service && go run cmd\api\main.go"

REM 3) Start Worker service in new window
echo [3/3] Starting worker-service...
start "worker-service" cmd /k "cd /d %~dp0..\services\worker-service && go run cmd\worker\main.go"

ec

echo Kafka UI: http://localhost:8085
echo RedisInsight: http://localhost:5540
echo API: http://localhost:8081

pause
