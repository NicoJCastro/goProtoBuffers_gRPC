# Notas de Go

## ¿Qué es Protocol Buffers (Protobuf)?
Protocol Buffers, también conocidos como **protobuf**, son un formato binario que facilita el almacenamiento e intercambio de datos en aplicaciones. Fue desarrollado por Google Inc. y publicado parcialmente bajo una licencia BSD de 3 cláusulas.

- **Características principales**:
  - Permite serializar y deserializar datos de manera eficiente.
  - Utilizado comúnmente para la comunicación entre servicios distribuidos.
  - Compatible con varios lenguajes de programación.

## ¿Qué es gRPC?
gRPC es un sistema de llamada a procedimiento remoto (RPC) de código abierto desarrollado inicialmente por Google. Utiliza HTTP/2 como protocolo de transporte y Protocol Buffers como lenguaje de descripción de interfaz (IDL).

- **Características principales**:
  - Ofrece soporte para múltiples lenguajes.
  - Proporciona características avanzadas como streaming bidireccional y autenticación.
  - Utiliza HTTP/2, lo que permite mejor rendimiento en comparación con HTTP/1.1, ya que soporta multiplexación y compresión de cabeceras.


## Comando protoc 

-  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/student.proto