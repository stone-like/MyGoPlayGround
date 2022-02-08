package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//lara.tarの中身
// 3de551b010625fe9f56d95b3d0d32dde305aded551937318b6ef6cc4e2ff67c8.json
// a1fbf9d45134247cd64343e28a88e83d637cae82ea0d9b96dab44bdb7cc03d66/
// a1fbf9d45134247cd64343e28a88e83d637cae82ea0d9b96dab44bdb7cc03d66/VERSION
// a1fbf9d45134247cd64343e28a88e83d637cae82ea0d9b96dab44bdb7cc03d66/json
// a1fbf9d45134247cd64343e28a88e83d637cae82ea0d9b96dab44bdb7cc03d66/layer.tar
// manifest.json
// repositories

func main() {

	file, err := os.Open("lara.tar")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// gr, _ := gzip.NewReader(file)
	// defer gr.Close()
	tr := tar.NewReader(file)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}

		if hdr.Name == "a1fbf9d45134247cd64343e28a88e83d637cae82ea0d9b96dab44bdb7cc03d66/layer.tar" {
			break
		}

	}

	//今trはa1fbf9d45134247cd64343e28a88e83d637cae82ea0d9b96dab44bdb7cc03d66/layer.tarを指しているところで止まっているはず
	// ここで新しくtarReaderを作りtarの中のtarであるlayer.tarを読む

	newTr := tar.NewReader(tr)

	for {
		hdr, err := newTr.Next()
		if err == io.EOF {
			break
		}

		if strings.Contains(hdr.Name, "composer.lock") {
			fmt.Println(hdr.Name)
		}

	}

	//layer.tarを読んでもとのtrが指す位置がおかしくなっていないか確認(Nextをしてmanifest.jsonであればよい)<-ok

	hdr, _ := tr.Next()
	//manifest.json
	fmt.Println(hdr.Name)

}
