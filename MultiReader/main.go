package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

//TeeReaderとMultiReaderで省メモリでSeekが使えないReaderで巻き戻しをしてあげたい
//例としてheader bodyみたいな構造があるとして、headerを読んだら残りはbodyしか残らない
//headerだけ読みだした後に、header+bodyで読まなければいけないやつがあるとheaderの読み込み用とheader+bodyの読み込み用にBufferが二ついることになる
//なので巻き戻しをしてあげる(seekがあればそれでもいいけど)

func main() {
	f, _ := os.Open("test")
	defer f.Close()

	var b1 bytes.Buffer
	var b2 bytes.Buffer

	b1.ReadFrom(f)

	r := io.TeeReader(&b1, &b2)

	bb := make([]byte, 11)
	//header abcdだけ読みだす
	r.Read(bb)
	//header abcd
	fmt.Println(string(bb))

	mr := io.MultiReader(&b2, &b1)
	bs, _ := io.ReadAll(mr)

	// header abcd
	// body jjj
	fmt.Println(string(bs))

}
