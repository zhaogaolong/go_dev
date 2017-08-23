/*
server 定义的所有的类型统一管理

*/

package server

import (
	"github.com/hpcloud/tail"
)

// app collect log config struct
type CollectConf struct {
	LogPath string
	Topic   string
}

type TailfObj struct {
	Tail *tail.Tail
	Conf CollectConf
}

type TextMsg struct {
	Msg   string
	Topic string
}

type TailObjMgr struct {
	TailObjs []*TailfObj
	MsgChan  chan *TextMsg
}
