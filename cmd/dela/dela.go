package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func logger(handler http.Handler) http.Handler {
	logger := func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
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
	fmt.Printf("listening on %q\n", addr)
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
	files := r.MultipartForm.File["files"]
	for i := range files {
		file, err := files[i].Open()
		if err != nil {
			log.Fatalf("%+v", errors.WithStack(err))
		}
		f, err := ioutil.TempFile("/tmp", "upload_")
		if err != nil {
			log.Fatalf("%+v", errors.WithStack(err))
		}
		if _, err := io.Copy(f, file); err != nil {
			log.Fatalf("%+v", errors.WithStack(err))
		}
		if err := f.Close(); err != nil {
			log.Fatalf("%+v", err)
		}
		if err := file.Close(); err != nil {
			log.Fatalf("%+v", err)
		}
		fmt.Printf("stored %q as %q\n", files[i].Filename, f.Name())
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
			<input type="file" name="files" multiple="multiple">
			<input type="submit" value="upload">
		</form>
	</body>
</html>
`
