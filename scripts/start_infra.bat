@echo off
cd /d %~dp0..\deploy
docker compose up -d
echo Infrastructure started!
echo Kafka UI: http://localhost:8085
echo RedisInsight: http://localhost:5540
pause
