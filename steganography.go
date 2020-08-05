package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Chunk : チャンクを管理する構造体を定義
type Chunk struct {
	Size uint32 // チャンクサイズ
	Type []byte // チャンクタイプ
	Data []byte // データ本体
	Crc  uint32 // CRC32の値
}

// PNGのシグネチャ
const pngSignature = "\x89PNG\r\n\x1a\n"

// パスワード
const password = "wLCCc9bNFavKkS7x3UVmIppKlGqVxKB4"

// 手軽にnバイト読み込む関数を定義
func readN(r io.Reader, n int) []byte {
	buf := make([]byte, n)
	cnt, err := r.Read(buf)
	if err != nil || n != cnt {
		return []byte{}
	}
	return buf
}

// 4byte整数を読む関数を定義
func readInt32(r io.Reader) uint32 {
	n := readN(r, 4)
	if len(n) != 4 {
		return 0
	}
	return binary.BigEndian.Uint32(n)
}

// ReadPNGFile : PNGファイルを読む
func ReadPNGFile(fname string) ([]Chunk, error) {
	// ファイルを全部メモリに読み込む
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, fmt.Errorf("ファイル読み込みエラー")
	}
	// 先頭のPNGシグネチャを取り込む
	r := bytes.NewReader(buf)
	if string(readN(r, 8)) != pngSignature {
		return nil, fmt.Errorf("PNGファイルではありません")
	}
	// 複数のチャンクを繰り返し読む
	res := []Chunk{}
	for {
		// サイズ、タイプ、データ、CRCを順に読む
		chunk := Chunk{}
		chunk.Size = readInt32(r)
		chunk.Type = readN(r, 4)
		chunk.Data = readN(r, int(chunk.Size))
		chunk.Crc = readInt32(r) // CRC32
		if len(chunk.Type) == 0 {
			break
		}
		// println("[CHUNK]", string(chunk.Type))
		if string(chunk.Type) == "xANG" {
			text := xorText(string(chunk.Data), password)
			println("[TEXT]", text)
		}
		res = append(res, chunk)
	}
	return res, nil
}

// uint32の値をバッファに書き込む
func writeUInt32(buf *bytes.Buffer, v uint32) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, v)
	buf.Write(b)
}

// WritePNGFile : PNGファイルを書く
func WritePNGFile(fname string, chunks []Chunk) error {
	// 書き込み用のバッファを作る
	buf := bytes.NewBuffer([]byte{})
	// シグネチャを書き込む
	buf.Write([]byte(pngSignature))
	// 繰り返しチャンクを書き込む
	for _, chunk := range chunks {
		if string(chunk.Type) == "IEND" {
			continue
		}
		writeUInt32(buf, chunk.Size)
		buf.Write(chunk.Type)
		buf.Write(chunk.Data)
		writeUInt32(buf, chunk.Crc)
		println("write:", string(chunk.Type))
	}
	// IENDを最後に書き込む
	writeUInt32(buf, 0)
	buf.Write([]byte("IEND"))
	buf.Write([]byte{0xAE, 0x42, 0x60, 0x82})
	return ioutil.WriteFile(fname, buf.Bytes(), 0644)
}

// 簡単な暗号化
func xorText(text string, password string) string {
	dat := []byte(text)
	key := []byte(password)
	for i := range []byte(text) {
		dat[i] ^= key[i%len(key)]
	}
	return string(dat)
}

// 独自チャンクを追加
func appendAngouChunk(chunks []Chunk, text string) []Chunk {
	chunk := Chunk{
		Size: uint32(len(text)),
		Type: []byte("xANG"),
		Data: []byte(xorText(text, password)),
	}
	checkData := append(chunk.Type, chunk.Data...)
	chunk.Crc = crc32.ChecksumIEEE(checkData)
	return append(chunks, chunk)
}

func main() {
	if len(os.Args) <= 1 {
		println("作成: steganography (infile) (outfile) (text)")
		println("確認: steganography (infile)")
	}

	if len(os.Args) == 3+1 {
		infile := os.Args[1]
		outfile := os.Args[2]
		text := os.Args[3]
		chunks, err := ReadPNGFile(infile)
		if err != nil {
			log.Fatal(err)
		}
		chunks = appendAngouChunk(chunks, text)
		WritePNGFile(outfile, chunks)
	}
	if len(os.Args) == 1+1 {
		ReadPNGFile(os.Args[1])
	}
}
