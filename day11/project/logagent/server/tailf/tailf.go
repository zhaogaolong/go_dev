package tailf

import (
	"fmt"
	"go_dev/day11/project/logagent/server"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/hpcloud/tail"
)

var (
	TailfObjsMgr *server.TailObjMgr
)

func InitTailf(cellectConf []server.CollectConf, chanSize int) (err error) {
	if len(cellectConf) == 0 {
		err = fmt.Errorf("invalid config for log collect, config:%v", chanSize)
		return
	}

	TailfObjsMgr = &server.TailObjMgr{}
	TailfObjsMgr.MsgChan = make(chan *server.TextMsg, chanSize)
	// TailfObjsMgr = &server.TailObjMgr{}

	for _, v := range cellectConf {
		obj := &server.TailfObj{
			Conf: v,
		}

		tailfs, tailErr := tail.TailFile(v.LogPath, tail.Config{
			ReOpen:    true,  // tail -F 重新打开
			Follow:    true,  // 更新后继续读
			MustExist: false, // 必须存在，true：如果不存在就退出，false：如果没有就等待检查。
			// Location: &tail.SeekInfo{Offset:0, Whence:2},
		})

		if tailErr != nil {
			// fmt.Println("tailf file err:", err)
			err = tailErr
			return
		}
		obj.Tail = tailfs

		TailfObjsMgr.TailObjs = append(TailfObjsMgr.TailObjs, obj)

		go readFromTail(obj)
	}

	return
}

func readFromTail(tailfs *server.TailfObj) {
	for true {
		line, ok := <-tailfs.Tail.Lines
		if !ok {
			logs.Warn("tailf chan closed reopen, filename:%v", tailfs.Tail.Filename)
			time.Sleep(time.Second)
			continue
		}

		textMsg := server.TextMsg{
			Msg:   line.Text,
			Topic: tailfs.Conf.Topic,
		}
		TailfObjsMgr.MsgChan <- &textMsg

	}

}
