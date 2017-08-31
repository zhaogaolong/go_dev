package tailf

import (
	"context"
	"fmt"
	"go_dev/day11/project/logagent/server"
	"go_dev/day11/project/logagent/server/etcd"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/hpcloud/tail"
)

var (
	TailfObjsMgr *server.TailObjMgr
)

func InitTailf(chanSize int) (err error) {

	// init msg chan
	TailfObjsMgr = &server.TailObjMgr{}
	TailfObjsMgr.MsgChan = make(chan *server.TextMsg, chanSize)

	// 检查配置
	if len(etcd.Collect) == 0 {
		err = fmt.Errorf("invalid config for log collect, config:%v", chanSize)
		return
	}

	// TailfObjsMgr = &server.TailObjMgr{}

	for _, v := range etcd.Collect {
		logs.Debug("for etcd collect logPath:%s  topic:%s", v.LogPath, v.Topic)
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
		ctx, cancal := context.WithCancel(context.Background())
		obj.Cancel = cancal
		obj.Status = false

		TailfObjsMgr.TailObjs = append(TailfObjsMgr.TailObjs, obj)

		go readFromTail(ctx, obj)
		go watchConfig()
	}

	return
}

func readFromTail(ctx context.Context, tailfs *server.TailfObj) {
	logs.Debug("tailf goroutine start. path:%s", tailfs.Conf.LogPath)
	for true {
		select {
		case <-ctx.Done():
			logs.Warn("tailf goroutine exit, monitor path:%s", tailfs.Conf.LogPath)
			return
		}

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

func watchConfig() {
	for {
		_ = <-etcd.EtcdConfigUpdate
		updateConfig()
	}
}

func updateConfig() {

	if len(etcd.Collect) == 0 {
		logs.Warn("etcd config tailf null, stop all tailf goroutine")

		for _, tailfObj := range TailfObjsMgr.TailObjs {
			tailfObj.Cancel()
		}

		return
	}

	// TailfObjsMgr = &server.TailObjMgr{}
	var tailfsObj []*server.TailfObj

	for _, v := range etcd.Collect {
		logs.Debug("etcd new collect logPath:%s  topic:%s", v.LogPath, v.Topic)
		var have bool
		for _, tailfObj := range TailfObjsMgr.TailObjs {
			if tailfObj.Conf.LogPath == v.LogPath {
				tailfObj.Status = true
				tailfsObj = append(tailfsObj, tailfObj)
				have = true
			}
		}
		// 如果没有就立马启动一个goroutine
		if !have {
			tailfObj := monPath(v.LogPath, v.Topic)
			tailfsObj = append(tailfsObj, tailfObj)
		}
	}
	// 找出需要停止的goroutine
	for _, tailfObj := range TailfObjsMgr.TailObjs {
		if !tailfObj.Status {
			tailfObj.Cancel()
		}
	}

	// 恢复原状
	for _, tailfObj := range TailfObjsMgr.TailObjs {
		tailfObj.Status = false
	}
	TailfObjsMgr.TailObjs = tailfsObj
}

func monPath(path, topic string) *server.TailfObj {

	cc := server.CollectConf{
		Topic:   topic,
		LogPath: path,
	}

	obj := &server.TailfObj{
		Conf: cc,
	}

	tailfs, _ := tail.TailFile(cc.LogPath, tail.Config{
		ReOpen:    true,  // tail -F 重新打开
		Follow:    true,  // 更新后继续读
		MustExist: false, // 必须存在，true：如果不存在就退出，false：如果没有就等待检查。
		// Location: &tail.SeekInfo{Offset:0, Whence:2},
	})

	// if tailErr != nil {
	// 	// fmt.Println("tailf file err:", err)
	// 	return
	// }
	obj.Tail = tailfs
	ctx, cancal := context.WithCancel(context.Background())
	obj.Cancel = cancal
	obj.Status = false

	go readFromTail(ctx, obj)
	return obj

}
