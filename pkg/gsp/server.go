package gsp

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
)

type EventHandler = func(peer *TcpPeer)

type TcpServer struct {
	listener           net.Listener
	listening          bool
	handlers           []EventHandler
	onPeerConnected    EventHandler
	onPeerDisconnected EventHandler
	maxClients         int32
	numClients         int32
}

func NewTcpServer(maxClients int32) TcpServer {
	server := TcpServer{}
	server.handlers = make([]EventHandler, 1000)
	if maxClients == 0 {
		server.maxClients = 1000
	} else {
		server.maxClients = maxClients
	}
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

func (s *TcpServer) OnPeerConnect(handler EventHandler) {
	s.onPeerConnected = handler
}

func (s *TcpServer) OnPeerDisconnect(handler EventHandler) {
	s.onPeerDisconnected = handler
}

func (s *TcpServer) OnEvent(evId uint16, handler EventHandler) {
	s.handlers[evId] = handler
}

func (s *TcpServer) connectionLoop() {
	var clientId int32
	for s.listening {
		clientId += 1
		conn, err := s.listener.Accept()
		if err != nil {
			println(err.Error())
			continue
		}
		println(s.maxClients, s.numClients)
		if int(s.numClients) == int(s.maxClients) {
			println("max clients excedeed, closing connection.")
			conn.Close()
			continue
		}
		peer := NewPeer(conn, clientId)
		s.onPeerConnected(peer)
		go s.readEvents(peer)
	}
}

func (s *TcpServer) readEvents(peer *TcpPeer) {
	for s.listening {
		var evId int16
		err := binary.Read(peer.conn, binary.BigEndian, &evId)
		if err != nil {
			err = handleError(err)
			if errors.Is(err, ErrDisconnected) {
				peer.Close()
				s.onPeerDisconnected(peer)
				break
			}
			println(err.Error())
			continue
		}
		handler := s.handlers[evId]
		if handler == nil {
			println("no handler found for id ", evId)
			continue
		}
		handler(peer)
		// if peer.writeBuffer.Len() > 0 {
		// 	// peer.conn.SetWriteDeadline(time.Now().Add(time.Millisecond * 10))
		// 	_, err := peer.writeBuffer.WriteTo(peer.conn)
		// 	if err != nil {
		// 		println(err.Error())
		// 	}
		// }
	}
	peer.Close()
	s.numClients -= 1
}
