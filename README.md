Description
===========

This project demonstrates, how JSON-RPC can be used in a Service Oriented Architecure instead of XML-like WSDL.

![Architecture]
(http://i.imgur.com/GevMfU6.jpg)



TLS Handshake 
-------------

client.go

```
2015/07/09 14:26:46 client: connected to:  127.0.0.1:9999
<confidential data was removed>
2015/07/09 14:26:46 client: handshake:  true
2015/07/09 14:26:46 client: mutual:  true
2015/07/09 14:26:46 client: wrote "Hello\n" (6 bytes)
2015/07/09 14:26:46 client: read "Hello\n" (6 bytes)
2015/07/09 14:26:46 client: exiting



```


server.go

```
2015/07/09 14:26:42 server: listening
2015/07/09 14:26:46 server: accepted from 127.0.0.1:48613
2015/07/09 14:26:46 ok=true
2015/07/09 14:26:46 server: conn: waiting
2015/07/09 14:26:46 server: conn: echo "Hello\n"
2015/07/09 14:26:46 server: conn: wrote 6 bytes
2015/07/09 14:26:46 server: conn: waiting
2015/07/09 14:26:46 server: conn: read: EOF
2015/07/09 14:26:46 server: conn: closed

```
