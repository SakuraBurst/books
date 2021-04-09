package helpers

import (
	"io"

	"github.com/SakuraBurst/books.git/main/models"
)

func MakeNewInstanse(instanse models.InstanseMaker, body io.ReadCloser) models.InstanseMaker {
	newInstanse := instanse.NewInstanseFromJson(body)
	return newInstanse
}
