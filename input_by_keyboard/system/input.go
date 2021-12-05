package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct {
	keys []ebiten.Key
}

func (i *Input) Update() {
	i.keys = inpututil.AppendPressedKeys(i.keys[:0])
}
