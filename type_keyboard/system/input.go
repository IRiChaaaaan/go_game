package system

import (
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct {
	runes   []rune
	text    string
	counter int
}

func (i *Input) Update() {
	// ユーザーが入力したルーンを追加
	// AppendInputCharsの結果はフレームごとに変わるため、フレームごとに呼び出す必要があること
	i.runes = ebiten.AppendInputChars(i.runes[:0])
	i.text += string(i.runes)
	// 文字列が最大 10 行になるように調整
	ss := strings.Split(i.text, "\n")
	if len(ss) > 10 {
		i.text = strings.Join(ss[len(ss)-10:], "\n")
	}
	// エンターキーが押された場合は、改行を追加
	if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyNumpadEnter) {
		i.text += "\n"
	}
	// バックスペースキーが押された場合は、1文字削除
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(i.text) >= 1 {
			i.text = i.text[:len(i.text)-1]
		}
	}
	i.counter++
}

// 繰り返し状態を考慮してキーが押されたときにtrueを返す
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}
