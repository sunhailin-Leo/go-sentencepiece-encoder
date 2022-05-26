SentencePiece Encoder V1
======================

Combine two repos here:
* [go-sentencepiece-encoder](https://github.com/vikesh-raj/go-sentencepiece-encoder)
* [sentencepiecego](https://github.com/evan176/sentencepiecego)

---

Reason:
* 1、`go-sentencepiece-encoder` has different result with Python [sentencepiece](https://github.com/google/sentencepiece) library on same text and model, but `sentencepiecego` not.
* 2、`sentencepiecego` just have only one api to encode the text to ids, but `go-sentencepiece-encoder` not.

---

Example:

```go
package main

import (
	"fmt"
	
	"github.com/sunhailin-Leo/go-sentencepiece-encoder/sentencepiece"
)

func main() {
	text := "This is a sample text"
	spm, _ := sentencepiece.NewSentencepieceFromFile("<Your Model Here>", false)
	
	tokens := spm.Tokenize(text)
	fmt.Println(tokens)
	
	spm.Free()
}

```