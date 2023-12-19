package shard

import "mmo2/pkg/events"

type EventHandler = func(player *Player, event events.Raw)
