package sentencepiece

import (
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/proto"
)

// NewSentencepieceFromFile creates sentencepiece from file.
func NewSentencepieceFromFile(filename string, lowercase bool) (Sentencepiece, error) {
	s, objErr := NewEmptySentencepiece(filename, lowercase)
	if objErr != nil {
		return s, objErr
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return s, fmt.Errorf("Unable to read file : %s, err %v", filename, err)
	}
	var model ModelProto
	err = proto.Unmarshal(bytes, &model)
	if err != nil {
		return s, fmt.Errorf("Unable to read model file : %s, err %v", filename, err)
	}

	count := 0
	for i, piece := range model.GetPieces() {
		typ := piece.GetType()
		word := piece.GetPiece()
		switch typ {
		case ModelProto_SentencePiece_NORMAL, ModelProto_SentencePiece_USER_DEFINED:
			s.insert(word, piece.GetScore(), int32(i))
		case ModelProto_SentencePiece_UNKNOWN:
			s.SetUnknownIndex(int32(i))
		case ModelProto_SentencePiece_CONTROL:
			s.SetControlWord(word, int32(i))
		}
		count++
	}

	return s, nil
}
