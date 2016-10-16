package main

//ejemplo simple y lento maneja 500 qps
import (
    "bufio"
	"log"
	"net"
    "net/http" //ventaja de GO a que la stl hace de todo
)

func main()  {
    //escucha conexion
    if ln , err := net.Listen("tcp" ,  ":8080"); err == nil {
        for{
                //Accepta conexion
                if conn, err := ln.Accept(); err == nil { //si aceptamos la conn
                    reader := bufio.NewReader(conn)
                    //lee el Request del cliente
                    if req , err := http.ReadRequest(reader);  err == nil{
                        //conecta a backend
                        if be , err := net.Dial("tcp" , "127.0.0.1:8081"); err == nil{ // be -> back-end
                            be_reader := bufio.NewReader(be)
                            //manda al be
                            if err :=  req.Write(be) ; err == nil{
                                //leemos la respuesta del be
                                if resp , err := http.ReadResponse( be_reader , req); err == nil{
                                    //mandar respuesta al cliente y cerrar
                                    resp.Close = true
                                    if err := resp.Write(conn); err == nil {
									   log.Printf("proxied %s: got %d", req.URL.Path, resp.StatusCode)
								       }
								conn.Close()
                                }
                            }
                        }
                    }

                }
        }//ciclo infinito que acepta solicitudes
    }//no tratamos los errores solo vemos que no se arroje uno

}
