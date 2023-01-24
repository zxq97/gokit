package generate

import (
	"time"

	"github.com/sony/sonyflake"
	"github.com/zxq97/gokit/pkg/constant"
)

type SonyIDGen struct {
	node *sonyflake.Sonyflake
}

func NewSonyIDDen(bt string) (*SonyIDGen, error) {
	st, err := time.Parse(constant.DateTimeFormat, bt)
	if err != nil {
		return nil, err
	}
	return &SonyIDGen{
		node: sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime: st,
		}),
	}, nil
}

func (g *SonyIDGen) GenID() int64 {
	id, _ := g.node.NextID()
	return int64(id)
}
