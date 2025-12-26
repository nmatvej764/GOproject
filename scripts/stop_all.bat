@echo off
echo Stopping infrastructure...
cd /d %~dp0..\deploy
docker compose down
echo DONE!
pause
