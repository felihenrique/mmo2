package shard

import "mmo2/pkg/event_utils"

type EventHandler = func(player *Player, event event_utils.Raw)
