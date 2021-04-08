package helpers

import (
	"io"

	"github.com/SakuraBurst/books.git/main/models"
)

func MakeNewInstanse(str models.InstanseMaker, body io.ReadCloser) models.InstanseMaker {
	str.NewInstanseFromJson(body)
	return str
}
