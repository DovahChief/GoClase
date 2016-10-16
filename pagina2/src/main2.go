package main

import ("net/http"
        "log"
        "strings"
        "os"
        "bufio")

type MiHandle struct {}

func (this *MiHandle) ServeHTTP(w http.ResponseWriter, r *http.Request){
    path := r.URL.Path[1:]
    log.Println(path)
    dato, err := os.Open(path)

    if err == nil {
        buffReader := bufio.NewReader(dato)
        var tipoCont string

        if strings.HasSuffix(path, ".css"){
            tipoCont = "text/css"
        }else if strings.HasSuffix(path, ".html"){
            tipoCont = "text/html"
        }else if strings.HasSuffix(path, ".woff"){
            tipoCont = "text/font"
        }else if strings.HasSuffix(path, ".ttf"){
            tipoCont = "text/font"
        }else if strings.HasSuffix(path, ".eot"){
            tipoCont = "text/font"
        }else if strings.HasSuffix(path, ".js"){
            tipoCont = "application/javascript"
        }else if strings.HasSuffix(path, ".png"){
            tipoCont = "image/png"
        }else if strings.HasSuffix(path, ".jpg"){
            tipoCont = "image/jpg"
        }else if strings.HasSuffix(path, ".scss"){
            tipoCont = "text/scss"
        }else if strings.HasSuffix(path, ".svg"){
            tipoCont = "image/svg+xml"
        }else if strings.HasSuffix(path, ".mp4"){
            tipoCont = "video/mp4"
        }else{
            tipoCont = "text/plain"
        }

        w.Header().Add("Content Type", tipoCont)
        buffReader.WriteTo(w)

    }else {
        w.WriteHeader(404)
        w.Write([]byte("404 bro - "+ http.StatusText(404)))
    }

}

func main()  {
    http.Handle("/", new(MiHandle))
    http.ListenAndServe(":8080", nil)
}


//   HANDLERS

//http.HandleFunc("/", handle)
//miMux := http.NewServeMux()
//http.ListenAndServe(":8080", nil)
//nil es el mux de servidor por defecto

// func handle(w http.ResponseWriter, req *http.Request)  {
//     w.Write([]byte("hola Clase"))
// }

// type persona struct{
//     fname string
// }
// pers1 := &persona{fname : "Pepe"}
// func (p *persona) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//     w.Write([]byte("Primer nombre: "+ p.fname))
// }
