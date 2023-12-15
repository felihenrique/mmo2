package gsp

import (
	"errors"
	"fmt"
	"mmo2/pkg/events"
	"net"
)

type EventHandler = func(peer *TcpPeer, rawEvent events.RawEvent)
type PeerHandler = func(peer *TcpPeer)

type TcpServer struct {
	listener           net.Listener
	listening          bool
	handlers           []EventHandler
	onPeerConnected    PeerHandler
	onPeerDisconnected PeerHandler
}

func NewTcpServer() TcpServer {
	server := TcpServer{}
	server.handlers = make([]EventHandler, 1000)
	server.listening = false
	return server
}

func (s *TcpServer) Listen(host string, port int) error {
	listener, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}
	s.listener = listener
	s.listening = true
	if s.onPeerConnected == nil {
		return errors.New("on peer connected is required")
	}
	if s.onPeerDisconnected == nil {
		return errors.New("on peer disconnected is required")
	}
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

func (s *TcpServer) OnPeerConnect(handler PeerHandler) {
	s.onPeerConnected = handler
}

func (s *TcpServer) OnPeerDisconnect(handler PeerHandler) {
	s.onPeerDisconnected = handler
}

func (s *TcpServer) OnEvent(evId int16, handler EventHandler) {
	s.handlers[evId] = handler
}

func (s *TcpServer) connectionLoop() {
	for s.listening {
		conn, err := s.listener.Accept()
		if err != nil {
			println(err.Error())
			continue
		}
		peer := NewPeer(conn)
		s.onPeerConnected(peer)
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
main:
	for s.listening {
		err := reader.FillFrom(peer.conn)
		if err != nil {
			err = handleError(err)
			if errors.Is(err, ErrDisconnected) {
				s.onPeerDisconnected(peer)
				break
			}
			println(err.Error())
			break
		}
		for {
			rawEvent, err := reader.Next()
			if err != nil {
				if errors.Is(err, events.ErrNotEnoughBytes) {
					break
				}
				println(err.Error())
				break main
			}
			evType := events.GetType(rawEvent)
			handler := s.handlers[evType]
			if handler == nil {
				println("no handler found for id ", evType)
				continue
			}
			handler(peer, rawEvent)
			reader.Pop()
			_, err = peer.writer.WriteTo(peer.conn)
			if err != nil {
				println(err.Error())
			}
		}
	}
}
