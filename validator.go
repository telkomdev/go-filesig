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

	r.Seek(0, io.SeekStart)

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

func genericMultipleCompareBuffer(r io.ReadSeeker, t [][]byte) bool {
	buff, err := checkBuffer(r, t[0])
	if err != nil {
		return false
	}

	for _, v := range t {
		if bytes.Compare(v, buff) == 0 {
			return true
		}
	}

	return false
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
	return genericMultipleCompareBuffer(r, [][]byte{
		DIB_0,
		DIB_1,
	})
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
	return genericMultipleCompareBuffer(r, [][]byte{
		MPG_0,
		MPG_1,
	})
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
	return genericMultipleCompareBuffer(r, [][]byte{
		SWF_0,
		SWF_1,
	})
}

// Is3gp function will return true if File is a valid 3gp
func Is3gp(r io.ReadSeeker) bool {
	return genericMultipleCompareBuffer(r, [][]byte{
		THREE_GP_0,
		THREE_GP_1,
		THREE_GP_2,
	})
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
	return genericMultipleCompareBuffer(r, [][]byte{
		ZIP_0,
		ZIP_1,
		ZIP_2,
	})
}

// IsWebp function will return true if File is a valid Webp
func IsWebp(r io.ReadSeeker) bool {
	return genericCompareBuffer(r, WEBP)
}

func IsMp4(r io.ReadSeeker) bool {
	buff, err := checkBuffer(r, MP4_0)
	if len(buff) < 9 {
		return false
	}

	validMpFour0 := bytes.Compare(MP4_0[4:8], buff[4:8])
	if validMpFour0 == 0 {
		return true
	}

	buff, err = checkBuffer(r, MP4_1)
	if err != nil {
		return false
	}

	validMpFour1 := bytes.Compare(MP4_1[4:8], buff[4:8])
	validMpFour2 := bytes.Compare(MP4_2[4:8], buff[4:8])

	return validMpFour1 == 0 && validMpFour2 == 0
}

// IsSvg function will return true if File is a valid SVG
func IsSvg(r io.ReadSeeker) bool {
	buff, err := io.ReadAll(r)
	if err != nil {
		return false
	}
	r.Seek(0, io.SeekStart)
	buff = HtmlCommentRegex.ReplaceAll(buff, []byte{})
	return SvgRegex.Match(buff) && !ScriptRegex.Match(buff)
}
