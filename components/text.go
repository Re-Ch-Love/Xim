package components

import (
	"github.com/LouisCheng-CN/xim/types"
)

// Text 文本
type Text struct {
	Content string
	events  chan *types.Event
	types.BaseComponent
}

func (t Text) Compose(_ *types.Context) {

}