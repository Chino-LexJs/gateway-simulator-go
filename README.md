# Simulador gateway para mensajes ISO 8583

_Aplicaci贸n Servidor para manejo y procesamiento de transferencias electr贸nicas de POS a PROSA con uso de ISO-8583_

_La presente aplicaci贸n actua como simulador para transacciones de tipo ISO 8583_

_El servidor recive tramas de servidores, procesa dicha informaci贸n y responde con mensajes hardcodeados de tipo ISO-8583_

### Requisitos para correr el proyecto 馃搵

_El presente proyecto corre con la version go1.19.1 windows/amd64 de Golang_

### Instalaci贸n 馃敡

_Para descargar Go puedes ir directamente a https://go.dev/dl/ y descargar la opci贸n de Microsoft Windows._

Una vez finalizada la instalaci贸n, se crear谩 la variable de entorno GOPATH apuntando a _%USERPROFILE%\go_ y all铆 mismo la carpeta pkg ambos de forma autom谩tica. Tambi茅n agregar谩 el binario ejecutable de Go a la variable de entorno _Path_ (en esta variable de entorno se guardan muchos otros ejecutables).

_Para verificar que la instalaci贸n haya sido correcta abre una terminal en Windows PowerShell y ejecuta el siguiente comando:_

```shell
 go version
```

### En tu GOPATH (_%USERPROFILE%\go_) crear la carpeta _src_ para que dentro guardes el presente proyecto de Go.

_Puedes ejecutar este c贸digo desde la consola con:_

```shell
go run .
```

## Construido con 馃洜锔?

_Herramientas usadas en el proyecto:_

- Golang

## Wiki 馃摉

Puedes encontrar mucho m谩s de c贸mo se usa el protocolo ISO-8583 este proyecto en [Wiki](https://es.wikipedia.org/wiki/ISO_8583)
