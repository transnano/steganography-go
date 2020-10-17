# steganography-go ![Releases](https://github.com/transnano/steganography-go/workflows/Releases/badge.svg) ![Publish Docker image](https://github.com/transnano/steganography-go/workflows/Publish%20Docker%20image/badge.svg) ![Vulnerability Scan](https://github.com/transnano/steganography-go/workflows/Vulnerability%20Scan/badge.svg) ![Code Scanning(CodeQL)](https://github.com/transnano/pagerduty-api/workflows/Code%20Scanning(CodeQL)/badge.svg)

![License](https://img.shields.io/github/license/transnano/steganography-go?style=flat)

![Container image version](https://img.shields.io/docker/v/transnano/steganography-go/latest?style=flat)
![Container image size](https://img.shields.io/docker/image-size/transnano/steganography-go/latest?style=flat)
![Container image pulls](https://img.shields.io/docker/pulls/transnano/steganography-go?style=flat)

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/transnano/steganography-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/transnano/steganography-go)](https://goreportcard.com/report/github.com/transnano/steganography-go)

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
