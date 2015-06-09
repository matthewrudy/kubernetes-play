package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
)

// Command-line flags.
var (
	httpAddr = flag.String("http", ":8080", "Listen address")
)

func main() {
	flag.Parse()
	http.Handle("/", NewServer())
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}

// Server implements the outyet server.
// It serves the user interface (it's an http.Handler)
// and polls the remote repository for changes.
type Server struct {
}

// NewServer returns an initialized outyet server.
func NewServer() *Server {
	return &Server{}
}

// ServeHTTP implements the HTTP user interface.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.render(w)
}

func (s *Server) render(w http.ResponseWriter) {
	data := struct {
		Words []string
	}{
		[]string{"Car", "Bar", "Jar"},
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Print(err)
	}
}

func getEnv() []string {
	env := os.Environ()
	sort.Strings(env)
	return env
}

// tmpl is the HTML template that drives the user interface.
var tmpl = template.Must(template.New("tmpl").Parse(`
<!DOCTYPE html><html><body>
	<h1>Words</h1>
	<ul>
  {{range $word := .Words}}
    <li>{{$word}}</li>
  {{end}}
	</ul>
</body></html>
`))
