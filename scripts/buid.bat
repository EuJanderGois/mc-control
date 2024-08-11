@echo off

:: Diretório de build
set BUILD_DIR=..\bin

:: Criar o diretório de build se não existir
if not exist %BUILD_DIR% mkdir %BUILD_DIR%

:: Compila o inicializador no destino de binarios
go build -o %BUILD_DIR%\start.exe ..\cmd\start.go

:: Compila o inicializador na raiz do servidor
go build -o ..\..\start.exe ..\cmd\start.go

echo Build completed successfully.

:: Pausar o script e esperar o usuário pressionar qualquer tecla
pause
