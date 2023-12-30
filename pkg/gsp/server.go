package gsp

import (
	"fmt"
	"mmo2/pkg/errors"
	"net"
)

type TcpServer struct {
	listener          net.Listener
	listening         bool
	newConnectionChan chan IPeer
	disconnectedChan  chan IPeer
	eventsChan        chan PeerEvent
}

func NewTcpServer() *TcpServer {
	server := TcpServer{}
	server.listening = false
	server.newConnectionChan = make(chan IPeer, 10)
	server.disconnectedChan = make(chan IPeer, 10)
	server.eventsChan = make(chan PeerEvent, 2048)
	return &server
}

func (s *TcpServer) Listening() bool {
	return s.listening
}

func (s *TcpServer) NewConnectionChan() <-chan IPeer {
	return s.newConnectionChan
}

func (s *TcpServer) DisconnectedChan() <-chan IPeer {
	return s.disconnectedChan
}

func (s *TcpServer) EventsChan() <-chan PeerEvent {
	return s.eventsChan
}

func (s *TcpServer) Listen(host string, port int) error {
	listener, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}
	s.listener = listener
	s.listening = true
	go s.connectionLoop()
	return nil
}

func (s *TcpServer) Close() error {
	err := s.listener.Close()
	if err != nil {
		return err
	}
	s.listening = false
	return nil
}

func (s *TcpServer) connectionLoop() {
	for s.listening {
		conn, err := s.listener.Accept()
		if err != nil {
			println(err.Error())
			continue
		}
		peer := NewPeer(conn)
		s.newConnectionChan <- peer
		go s.readEvents(peer)
	}
}

func (s *TcpServer) readEvents(peer *TcpPeer) {
	defer func() {
		s.disconnectedChan <- peer
		peer.Close()
	}()

	handleDisconnect := func(err error) bool {
		err = errors.Handle(err)
		if err == nil || errors.Is(err, errors.ErrTimeout) {
			return false
		}
		println(err.Error())
		return true
	}

	for s.listening && peer.connected {
		if handleDisconnect(peer.readEvents()) {
			break
		}
		for {
			event, err := peer.reader.Next()
			if err != nil {
				break
			}
			s.eventsChan <- PeerEvent{
				Peer:  peer,
				Event: event,
			}
			peer.reader.Pop()
			peer.updateRateLimit()
		}
		if handleDisconnect(peer.writeEvents()) {
			break
		}
	}
}
