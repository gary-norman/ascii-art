package test

import (
	"github.com/gary-norman/ascii-art/api"
	"github.com/gary-norman/ascii-art/pkg"
)

func test() {
	test := []string{"hello",
		"HELLO",
		"HeLlo HuMaN",
		"1Hello 2There",
		"Hello\nThere",
		"Hello\n\nThere",
		"{Hello & There #}",
		"hello There 1 to 2!",
		"MaD3IrA&LiSboN",
		"1a\"#FdwHywR&/()=",
		"{|}~",
		"[\\]^_ 'a",
		"RGB",
		":;<=>?@",
		`'\\!" #$%&'"'"'()*+,-./'`,
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"abcdefghijklmnopqrstuvwxyz"}
	for k := 0; k < len(test); k++ {
		api.MakeArt(test[k], api.GetChars(pkg.PrepareBanner("")))
	}
}
