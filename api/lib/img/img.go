package img

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"

	"github.com/nfnt/resize"
)

// max width of thumbnail = 450px
const Thumbnail uint = 450

type Img struct {
	data           []byte
	base64Encoding string
}

func FromBytes(src []byte) (*Img, error) {
	var buf bytes.Buffer
	w := base64.NewEncoder(base64.RawStdEncoding, &buf)
	if _, err := w.Write(src); err != nil {
		return nil, err
	}
	return &Img{src, buf.String()}, nil
}

func FromBase64(src string) (*Img, error) {
	decoded, err := base64.RawStdEncoding.DecodeString(src)
	if err != nil {
		return nil, err
	}
	return &Img{decoded, src}, nil
}

func (img *Img) Bytes() []byte {
	return img.data
}

func (img *Img) Base64() string {
	return img.base64Encoding
}

func (img *Img) Resize(size uint) error {
	// resize image
	image, err := jpeg.Decode(bytes.NewReader(img.data))
	if err != nil {
		return err
	}

	newImg := resize.Resize(size, 0, image, resize.Lanczos2)
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, newImg, nil); err != nil {
		return err
	}

	// get `img.data`
	img.data = buf.Bytes()

	// get `img.base64Encoding`
	w := base64.NewEncoder(base64.RawStdEncoding, &buf)
	defer w.Close()
	if _, err := w.Write(img.data); err != nil {
		return err
	}

	img.base64Encoding = buf.String()

	return nil
}
