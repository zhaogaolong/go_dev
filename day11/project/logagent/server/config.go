/*
server 定义的所有的类型统一管理

*/

package server

import (
	"context"

	"github.com/hpcloud/tail"
)

// app collect log config struct
type CollectConf struct {
	LogPath string `json:"log_path"`
	Topic   string `josn:"topic"`
}

type TailfObj struct {
	Tail   *tail.Tail
	Conf   CollectConf
	Cancel context.CancelFunc
	Status bool
	// Cancel interface{}
}

type TextMsg struct {
	Msg   string
	Topic string
}

type TailObjMgr struct {
	TailObjs []*TailfObj
	MsgChan  chan *TextMsg
}
