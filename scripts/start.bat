@echo off

:: Diretório dos binários
set BIN_DIR=..\bin

echo Starting building process ...

.\build.bat && echo Starting API execution ... && ..\bin\api.exe