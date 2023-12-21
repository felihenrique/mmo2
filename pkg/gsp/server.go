package gsp

import (
	"fmt"
	"mmo2/pkg/errors"
	"mmo2/pkg/events"
	"net"
	"time"
)

type PeerEvent struct {
	Peer  IPeer
	Event events.Raw
}

type IServer interface {
	Listening() bool
	NewConnectionChan() <-chan IPeer
	DisconnectedChan() <-chan IPeer
	EventsChan() <-chan PeerEvent
	Listen(host string, port int) error
	Close() error
}

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
		if r := recover(); r != nil {
			fmt.Printf("error in peer %s: %s \n", peer.Addr(), r)
		}
		peer.Close()
	}()

	for s.listening {
		// READ
		peer.conn.SetReadDeadline(time.Now().Add(time.Millisecond * 200))
		err := errors.Handle(peer.reader.FillFrom(peer.conn))
		if err != nil && !errors.Is(err, errors.ErrTimeout) {
			println(err.Error())
			s.disconnectedChan <- peer
			break
		}
		for {
			rawEvent, err := peer.reader.Next()
			if err != nil {
				break
			}
			s.eventsChan <- PeerEvent{
				Peer:  peer,
				Event: rawEvent,
			}
			peer.reader.Pop()
		}
		// WRITE
		peer.conn.SetWriteDeadline(time.Now().Add(time.Millisecond * 200))
		_, err = peer.writer.Send(peer.conn)
		err = errors.Handle(err)
		if err != nil && !errors.Is(err, errors.ErrTimeout) {
			println(err.Error())
			s.disconnectedChan <- peer
			break
		}
	}
}
