# Simulador gateway para mensajes ISO 8583

_Aplicación Servidor para manejo y procesamiento de transferencias electrónicas de POS a PROSA con uso de ISO-8583_

_La presente aplicación actua como simulador para transacciones de tipo ISO 8583_

_El servidor recive tramas de servidores, procesa dicha información y responde con mensajes hardcodeados de tipo ISO-8583_

### Requisitos para correr el proyecto 📋

_El presente proyecto corre con la version go1.19.1 windows/amd64 de Golang_

### Instalación 🔧

_Para descargar Go puedes ir directamente a https://go.dev/dl/ y descargar la opción de Microsoft Windows._

Una vez finalizada la instalación, se creará la variable de entorno GOPATH apuntando a _%USERPROFILE%\go_ y allí mismo la carpeta pkg ambos de forma automática. También agregará el binario ejecutable de Go a la variable de entorno _Path_ (en esta variable de entorno se guardan muchos otros ejecutables).

_Para verificar que la instalación haya sido correcta abre una terminal en Windows PowerShell y ejecuta el siguiente comando:_

```shell
 go version
```

### En tu GOPATH (_%USERPROFILE%\go_) crear la carpeta _src_ para que dentro guardes el presente proyecto de Go.

_Puedes ejecutar este código desde la consola con:_

```shell
go run .
```

## Construido con 🛠️

_Herramientas usadas en el proyecto:_

- Golang

## Wiki 📖

Puedes encontrar mucho más de cómo se usa el protocolo ISO-8583 este proyecto en [Wiki](https://es.wikipedia.org/wiki/ISO_8583)
