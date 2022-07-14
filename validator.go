package filesig

import (
	"bytes"
	"io"
)

type function func(io.Reader) bool

// IsOneOf function will validate File with multiple function
func IsOneOf(r io.Reader, functions ...function) bool {
	for _, f := range functions {
		bufRead := &bytes.Buffer{}
		tee := io.TeeReader(r, bufRead)

		valid := f(tee)

		if valid {
			return true
		}

		r = bufRead
	}

	return false
}

func readUntil(l int, r io.Reader) ([]byte, error) {
	buff := make([]byte, l)

	_, err := r.Read(buff)
	if err != nil {
		return nil, err
	}

	return buff, nil
}

// IsPng function will return true if File is a valid PNG
func IsPng(r io.Reader) bool {
	l := len(PNG)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid := bytes.Compare(PNG, buff)
	if valid != 0 {
		return false
	}

	return true
}

// IsJpeg function will return true if File is a valid JPEG
func IsJpeg(r io.Reader) bool {
	l := len(JPEG)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid := bytes.Compare(JPEG, buff)
	if valid != 0 {
		return false
	}

	return true
}

// IsPdf function will return true if File is a valid PDF
func IsPdf(r io.Reader) bool {
	l := len(PDF)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid := bytes.Compare(PDF, buff)
	if valid != 0 {
		return false
	}

	return true
}

// IsGif function will return true if File is a valid GIF
func IsGif(r io.Reader) bool {
	l := len(GIF)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid := bytes.Compare(GIF, buff)
	if valid != 0 {
		return false
	}

	return true
}

// IsAvif function will return true if File is a valid AVIF
func IsAvif(r io.Reader) bool {
	l := len(AVIF)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid := bytes.Compare(AVIF, buff)
	if valid != 0 {
		return false
	}

	return true
}

// IsBmp function will return true if File is a valid BMP
func IsBmp(r io.Reader) bool {
	l := len(BMP)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid := bytes.Compare(BMP, buff)
	if valid != 0 {
		return false
	}

	return true
}

// IsDib function will return true if File is a valid DIB
func IsDib(r io.Reader) bool {
	l := len(DIB)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid := bytes.Compare(DIB, buff)
	if valid != 0 {
		return false
	}

	return true
}

// IsTiff function will return true if File is a valid TIFF
func IsTiff(r io.Reader) bool {
	l := len(TIFF)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid := bytes.Compare(TIFF, buff)
	if valid != 0 {
		return false
	}

	return true
}

// IsMp3 function will return true if File is a valid MP3
func IsMp3(r io.Reader) bool {
	l := len(MP3)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid := bytes.Compare(MP3, buff)
	if valid != 0 {
		return false
	}

	return true
}

// IsMpg function will return true if File is a valid MPG
func IsMpg(r io.Reader) bool {
	l1 := len(MPG_0)

	buff, err := readUntil(l1, r)
	if err != nil {
		return false
	}

	valid_mpg_0 := bytes.Compare(MPG_0, buff)
	valid_mpg_1 := bytes.Compare(MPG_1, buff)
	if valid_mpg_0 != 0 && valid_mpg_1 != 0 {
		return false
	}

	return true
}

// IsFlv function will return true if File is a valid FLV
func IsFlv(r io.Reader) bool {
	l := len(FLV)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid := bytes.Compare(FLV, buff)
	if valid != 0 {
		return false
	}

	return true
}

// IsApk function will return true if File is a valid APK
func IsApk(r io.Reader) bool {
	l := len(APK)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid := bytes.Compare(APK, buff)
	if valid != 0 {
		return false
	}

	return true
}

// IsMsOffice function will return true if File is a valid MS OFFICE Document (DOCX|PPTX|XLSX)
func IsMsOffice(r io.Reader) bool {
	l := len(MS_OFFICE)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid := bytes.Compare(MS_OFFICE, buff)
	if valid != 0 {
		return false
	}

	return true
}

// IsJar function will return true if File is a valid JAR (Java Archive)
func IsJar(r io.Reader) bool {
	l := len(JAR)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid := bytes.Compare(JAR, buff)
	if valid != 0 {
		return false
	}

	return true
}

// IsSwf function will return true if File is a valid SWF
func IsSwf(r io.Reader) bool {
	l := len(SWF)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid := bytes.Compare(SWF, buff)
	if valid != 0 {
		return false
	}

	return true
}

// Is3gp function will return true if File is a valid 3gp
func Is3gp(r io.Reader) bool {
	l1 := len(THREE_GP_0)

	buff, err := readUntil(l1, r)
	if err != nil {
		return false
	}

	valid_3gp_0 := bytes.Compare(THREE_GP_0, buff)
	valid_3gp_1 := bytes.Compare(THREE_GP_1, buff)
	valid_3gp_2 := bytes.Compare(THREE_GP_2, buff)
	if valid_3gp_0 != 0 && valid_3gp_1 != 0 && valid_3gp_2 != 0 {
		return false
	}

	return true
}

// IsMkv function will return true if File is a valid MKV
func IsMkv(r io.Reader) bool {
	l := len(MKV)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid := bytes.Compare(MKV, buff)
	if valid != 0 {
		return false
	}

	return true
}

// IsRar function will return true if File is a valid RAR
func IsRar(r io.Reader) bool {
	l := len(RAR)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid := bytes.Compare(RAR, buff)
	if valid != 0 {
		return false
	}

	return true
}

// IsGzip function will return true if File is a valid GZIP
func IsGzip(r io.Reader) bool {
	l := len(GZIP)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid := bytes.Compare(GZIP, buff)
	if valid != 0 {
		return false
	}

	return true
}

// IsZip function will return true if File is a valid ZIP
func IsZip(r io.Reader) bool {
	l := len(ZIP)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid := bytes.Compare(ZIP, buff)
	if valid != 0 {
		return false
	}

	return true
}

// IsWebp function will return true if File is a valid Webp
func IsWebp(r io.Reader) bool {
	l := len(WEBP)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid := bytes.Compare(WEBP, buff)
	if valid != 0 {
		return false
	}

	return true
}

func IsMp4(r io.Reader) bool {
	l := len(MP4_0)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	if len(buff) < 9 {
		return false
	}

	valid_mp4_0 := bytes.Compare(MP4_0[4:8], buff[4:8])

	if valid_mp4_0 == 0 {
		return true
	}

	l = len(MP4_1)

	buff, err = readUntil(l, r)
	if err != nil {
		return false
	}

	valid_mp4_1 := bytes.Compare(MP4_1[4:8], buff[4:8])
	valid_mp4_2 := bytes.Compare(MP4_2[4:8], buff[4:8])
	if valid_mp4_1 != 0 && valid_mp4_2 != 0 {
		return false
	}

	return true
}
