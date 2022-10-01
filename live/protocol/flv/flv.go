package flv

import (
	"github.com/zhangshanwen/shard/live/protocol/rtmp"
)

type stream struct {
	Key string `json:"key"`
	Id  string `json:"id"`
}

type streams struct {
	Publishers []stream `json:"publishers"`
	Players    []stream `json:"players"`
}

func GetStreams(s *rtmp.RtmpStream) *streams {
	messages := new(streams)
	s.GetStreams().Range(func(key, val interface{}) bool {
		if k, ok := val.(*rtmp.Stream); ok {
			if k.GetReader() != nil {
				msg := stream{key.(string), k.GetReader().Info().UID}
				messages.Publishers = append(messages.Publishers, msg)
			}
		}
		return true
	})

	s.GetStreams().Range(func(key, val interface{}) bool {
		ws := val.(*rtmp.Stream).GetWs()

		ws.Range(func(k, v interface{}) bool {
			if pw, ok := v.(*rtmp.PackWriterCloser); ok {
				if pw.GetWriter() != nil {
					msg := stream{key.(string), pw.GetWriter().Info().UID}
					messages.Players = append(messages.Players, msg)
				}
			}
			return true
		})
		return true
	})

	return messages
}
