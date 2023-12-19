package shard

import "mmo2/pkg/events"

type EventHandler = func(event events.Raw)
