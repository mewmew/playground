package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kr/pretty"
	"github.com/pkg/errors"
)

func logger(handler http.Handler) http.Handler {
	logger := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(logger)
}

func main() {
	var (
		addr string
		recv bool
	)
	flag.StringVar(&addr, "http", ":8080", "HTTP service address")
	flag.BoolVar(&recv, "recv", false, "enable upload handler")
	flag.BoolVar(&recv, "max", false, "enable upload handler")
	flag.Parse()
	var handler http.Handler
	if recv {
		handler = logger(http.HandlerFunc(upload))
	} else {
		handler = logger(http.FileServer(http.Dir(".")))
	}
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}

// upload handles upload requests.
func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if _, err := fmt.Fprintf(w, uploadPage[1:]); err != nil {
			log.Fatalf("%+v", errors.WithStack(err))
		}
		return
	}
	if err := r.ParseMultipartForm(1024 * 1024 * 1024); err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}
	file, hdr, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	pretty.Println("hdr:", hdr.Header)
	f, err := ioutil.TempFile("/tmp", "upload_")
	if err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}
	defer f.Close()
	if _, err := io.Copy(f, file); err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}
}

const uploadPage = `
<!doctype html>
<html>
	<head>
		<title>upload file</title>
	</head>
	<body>
		<form enctype="multipart/form-data" action="/" method="POST">
			<input type="file" name="file">
			<input type="submit" value="upload">
		</form>
	</body>
</html>
`
