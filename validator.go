package filesig

import (
	"bytes"
	"io"
)

type function func(io.ReadSeeker) bool

// IsOneOf function will validate File with multiple function
func IsOneOf(r io.ReadSeeker, functions ...function) bool {
	for _, f := range functions {
		valid := f(r)
		if valid {
			return true
		}
	}

	return false
}

func readUntil(l int, r io.ReadSeeker) ([]byte, error) {
	buff := make([]byte, l)

	_, err := r.Read(buff)
	if err != nil {
		return nil, err
	}

	_, err = r.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	return buff, nil
}

func checkBuffer(r io.ReadSeeker, t []byte) ([]byte, error) {
	l := len(t)

	buff, err := readUntil(l, r)
	if err != nil {
		return make([]byte, 0), err
	}

	return buff, nil
}

func genericCompareBuffer(r io.ReadSeeker, t []byte) bool {
	buff, err := checkBuffer(r, t)
	if err != nil {
		return false
	}

	valid := bytes.Compare(t, buff)
	return valid == 0
}

// IsPng function will return true if File is a valid PNG
func IsPng(r io.ReadSeeker) bool {
	return genericCompareBuffer(r, PNG)
}

// IsJpeg function will return true if File is a valid JPEG
func IsJpeg(r io.ReadSeeker) bool {
	return genericCompareBuffer(r, JPEG)
}

// IsPdf function will return true if File is a valid PDF
func IsPdf(r io.ReadSeeker) bool {
	return genericCompareBuffer(r, PDF)
}

// IsGif function will return true if File is a valid GIF
func IsGif(r io.ReadSeeker) bool {
	return genericCompareBuffer(r, GIF)
}

// IsAvif function will return true if File is a valid AVIF
func IsAvif(r io.ReadSeeker) bool {
	return genericCompareBuffer(r, AVIF)
}

// IsBmp function will return true if File is a valid BMP
func IsBmp(r io.ReadSeeker) bool {
	return genericCompareBuffer(r, BMP)
}

// IsDib function will return true if File is a valid DIB
func IsDib(r io.ReadSeeker) bool {
	return genericCompareBuffer(r, DIB)
}

// IsTiff function will return true if File is a valid TIFF
func IsTiff(r io.ReadSeeker) bool {
	return genericCompareBuffer(r, TIFF)
}

// IsMp3 function will return true if File is a valid MP3
func IsMp3(r io.ReadSeeker) bool {
	return genericCompareBuffer(r, MP3)
}

// IsMpg function will return true if File is a valid MPG
func IsMpg(r io.ReadSeeker) bool {
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
func IsFlv(r io.ReadSeeker) bool {
	return genericCompareBuffer(r, FLV)
}

// IsApk function will return true if File is a valid APK
func IsApk(r io.ReadSeeker) bool {
	return genericCompareBuffer(r, APK)
}

// IsMsOffice function will return true if File is a valid MS OFFICE Document (DOCX|PPTX|XLSX)
func IsMsOffice(r io.ReadSeeker) bool {
	return genericCompareBuffer(r, MS_OFFICE)
}

// IsJar function will return true if File is a valid JAR (Java Archive)
func IsJar(r io.ReadSeeker) bool {
	return genericCompareBuffer(r, JAR)
}

// IsSwf function will return true if File is a valid SWF
func IsSwf(r io.ReadSeeker) bool {
	return genericCompareBuffer(r, SWF)
}

// Is3gp function will return true if File is a valid 3gp
func Is3gp(r io.ReadSeeker) bool {
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
func IsMkv(r io.ReadSeeker) bool {
	return genericCompareBuffer(r, MKV)
}

// IsRar function will return true if File is a valid RAR
func IsRar(r io.ReadSeeker) bool {
	return genericCompareBuffer(r, RAR)
}

// IsGzip function will return true if File is a valid GZIP
func IsGzip(r io.ReadSeeker) bool {
	return genericCompareBuffer(r, GZIP)
}

// IsZip function will return true if File is a valid ZIP
func IsZip(r io.ReadSeeker) bool {
	l := len(ZIP_0)

	buff, err := readUntil(l, r)
	if err != nil {
		return false
	}

	valid_zip_0 := bytes.Compare(ZIP_0, buff)
	valid_zip_1 := bytes.Compare(ZIP_0, buff)
	valid_zip_2 := bytes.Compare(ZIP_2, buff)
	if valid_zip_0 != 0 && valid_zip_1 != 0 && valid_zip_2 != 0 {
		return false
	}

	return true
}

// IsWebp function will return true if File is a valid Webp
func IsWebp(r io.ReadSeeker) bool {
	return genericCompareBuffer(r, WEBP)
}

func IsMp4(r io.ReadSeeker) bool {
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
