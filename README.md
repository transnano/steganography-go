# steganography-go ![Releases](https://github.com/transnano/steganography-go/workflows/Releases/badge.svg) ![Publish Docker image](https://github.com/transnano/steganography-go/workflows/Publish%20Docker%20image/badge.svg) ![Vulnerability Scan](https://github.com/transnano/steganography-go/workflows/Vulnerability%20Scan/badge.svg)

Embed confidential information in png image.

## Usage

作成

```sh
$ steganography (infile) (outfile) (input-text)

# ex)
$ steganography tea.png out.png 捜すのに時あり、諦めるのに時がある
```

確認

```sh
$ steganography (infile)

# ex)
$ steganography out.png
[TEXT] 捜すのに時あり、諦めるのに時がある
```

## Ref

- [ゼロからはじめるGo言語(13) PNG画像に暗号文を埋め込むプログラムを作ってみよう | マイナビニュース](https://news.mynavi.jp/article/gogogo-13/)
