package gsp

import (
	"fmt"
	"mmo2/pkg/events"
	"net"
	"os"
	"time"
)

type PeerEvent struct {
	Peer  *TcpPeer
	Event events.RawEvent
}

type TcpServer struct {
	listener         net.Listener
	listening        bool
	peerConnected    chan *TcpPeer
	peerDisconnected chan *TcpPeer
	newEventsChan    chan PeerEvent
}

func NewTcpServer() TcpServer {
	server := TcpServer{}
	server.listening = false
	server.peerConnected = make(chan *TcpPeer, 10)
	server.peerDisconnected = make(chan *TcpPeer, 10)
	server.newEventsChan = make(chan PeerEvent, 2048)
	return server
}

func (s *TcpServer) Listening() bool {
	return s.listening
}

func (s *TcpServer) PeerConnChan() <-chan *TcpPeer {
	return s.peerConnected
}

func (s *TcpServer) PeerDisChan() <-chan *TcpPeer {
	return s.peerDisconnected
}

func (s *TcpServer) NewEventsChan() <-chan PeerEvent {
	return s.newEventsChan
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
		s.peerConnected <- peer
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
	reader := events.NewReader()
	for s.listening {
		// READ
		err := handleError(reader.FillFrom(peer.conn))
		if err != nil && err != os.ErrDeadlineExceeded {
			println(err.Error())
			s.peerDisconnected <- peer
			break
		}
		for {
			rawEvent, err := reader.Next()
			if err != nil {
				break
			}
			s.newEventsChan <- PeerEvent{
				Peer:  peer,
				Event: rawEvent,
			}
			reader.Pop()
		}
		// WRITE
		peer.conn.SetWriteDeadline(time.Now().Add(time.Millisecond * 100))
		_, err = peer.writer.WriteTo(peer.conn)
		err = handleError(err)
		if err != nil && err != os.ErrDeadlineExceeded {
			println(err.Error())
			s.peerDisconnected <- peer
			break
		}
	}
}
