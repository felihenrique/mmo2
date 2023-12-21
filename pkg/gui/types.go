package gui

type Control struct {
	parent  *Control
	x       int32
	y       int32
	width   int32
	height  int32
	padding int32
}

type Panel struct {
	control Control
}
