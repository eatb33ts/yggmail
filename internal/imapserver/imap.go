package imapserver

import (
	"log"

	idle "github.com/emersion/go-imap-idle"
	"github.com/emersion/go-imap/server"
	"github.com/emersion/go-sasl"
)

type IMAPServer struct {
	server  *server.Server
	backend *Backend
}

func NewIMAPServer(backend *Backend, addr string, insecure bool) (*IMAPServer, *IMAPNotify, error) {
	s := &IMAPServer{
		server:  server.New(backend),
		backend: backend,
	}
	notify := NewIMAPNotify(s.server, backend.Log)
	s.server.Addr = addr
	s.server.AllowInsecureAuth = insecure
	//s.server.Debug = os.Stdout
	s.server.Enable(idle.NewExtension())
	s.server.Enable(notify)
	s.server.EnableAuth(sasl.Login, func(conn server.Conn) sasl.Server {
		return sasl.NewLoginServer(func(username, password string) error {
			_, err := s.backend.Login(nil, username, password)
			return err
		})
	})
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	return s, notify, nil
}
