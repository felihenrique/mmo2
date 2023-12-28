package shard

import (
	"mmo2/pkg/events"
	"mmo2/pkg/serialization"
)

type EventHandler = func(event events.Raw) serialization.ISerializable
type ResponseHandler = func(response serialization.ISerializable)

var SendEventsChan chan serialization.ISerializable = make(chan serialization.ISerializable, 1024)
