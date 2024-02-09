package util

import (
	"bytes"
	"image"
	"image/jpeg"
	"os"

	"github.com/blackflagsoftware/agenda/config"
	"github.com/skip2/go-qrcode"
)

type (
	QrImage struct {
		Name  string
		Image []byte
		Error error
	}
)

func NewQrImage(link string) *QrImage {
	q := &QrImage{}
	q.Create(link)
	return q
}
func (q *QrImage) Create(link string) {
	q.Image, q.Error = qrcode.Encode(link, qrcode.Medium, 32)
	if q.Error == nil {
		name := GenerateRandomString(8)
		q.Name = config.DocumentDir + "/assets/" + name + ".jpg"
		img, _, err := image.Decode(bytes.NewReader(q.Image))
		if err != nil {
			q.Error = err
		}
		out, _ := os.Create(q.Name)
		defer out.Close()
		var opts jpeg.Options
		opts.Quality = 100
		err = jpeg.Encode(out, img, &opts)
		if err != nil {
			q.Error = err
		}
		// q.Error = os.WriteFile(q.Name, q.Image, 0644)
	}
}

func (q *QrImage) Close() {
	// remove
	if _, err := os.Stat(q.Name); os.IsExist(err) {
		os.Remove(q.Name)
	}
}
