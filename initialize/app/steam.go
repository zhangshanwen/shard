package app

import (
	"github.com/zhangshanwen/shard/initialize/conf"
	"net"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/zhangshanwen/shard/live/protocol/rtmp"
)

var (
	S *rtmp.RtmpStream
)

func initStream() {
	S = rtmp.NewRtmpStream()
}

func startRtmp(stream *rtmp.RtmpStream) {
	rtmpAddr := conf.C.Rtmp

	rtmpListen, err := net.Listen("tcp", rtmpAddr)
	if err != nil {
		log.Fatal(err)
	}

	var rtmpServer *rtmp.Server

	rtmpServer = rtmp.NewRtmpServer(stream, nil)

	defer func() {
		if r := recover(); r != nil {
			log.Error("RTMP server panic: ", r)
		}
	}()
	log.Info("RTMP Listen On ", rtmpAddr)
	rtmpServer.Serve(rtmpListen)
}

func InitRtmp() {
	defer func() {
		if r := recover(); r != nil {
			log.Error("rtmp panic: ", r)
			time.Sleep(1 * time.Second)
		}
	}()
	initStream()
	startRtmp(S)
}
