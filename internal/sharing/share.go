package sharing

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
	"github.com/pale-whale/share.me/internal/archive"
)

type Server struct {
	files    map[string]string
	listener net.Listener
}

func (hdlr Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pth := hdlr.files[r.URL.Path[1:]]
	if pth == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404")
		return
	}
	filename := path.Base(pth)

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	file, err := os.Open(pth)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	io.Copy(w, file)
}

func (hdlr *Server) ServeFile(path string, createId bool) string {
	strId := ""
	if createId {
		id := uuid.New()
		strId = id.String()[0:8]
	}
	fileinfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Fatal("File does not exist.")
	}
	if fileinfo.IsDir() {
		archive.Tar(path, "/tmp/")
	}
	hdlr.files[strId] = path
	return strId
}

func (s *Server) Serve() {
	srv := http.Server{
		Addr:    s.listener.Addr().String(),
		Handler: *s,
	}
	srv.Serve(s.listener)
}

func (s *Server) Addr() string {
	return fmt.Sprintf("http://%v", s.listener.Addr().String())
}

func findListenAddress(port string) net.Listener {
	addrs, _ := net.InterfaceAddrs()
	for _, a := range addrs {
		addr := strings.Split(a.String(), "/")[0]
		if addr != "127.0.0.1" {
			ln, err := net.Listen("tcp", fmt.Sprintf("%s:%v", addr, port))
			if err == nil {
				return ln
			}
		}
	}
	return nil
}

func CreateServer(port string) *Server {
	hdlr := &Server{
		files:    map[string]string{},
		listener: findListenAddress(port),
	}

	return hdlr
}
