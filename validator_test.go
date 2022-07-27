package filesig

import (
	"bytes"
	"os"
	"testing"
)

func TestIsOneOf(t *testing.T) {
	buff, err := os.Open("./tmp/sample-0.pdf")

	if err != nil {
		t.Error("error: PDF file not found")
	}

	defer func() {
		buff.Close()
	}()

	valid := IsOneOf(buff, IsPng, IsJpeg, IsPdf, IsApk)
	if !valid {
		t.Error("error: buffer not valid PDF file")
	}

	valid = IsOneOf(buff, IsPng, IsJpeg, IsApk)
	if valid {
		t.Error("error: buffer not valid PDF file")
	}
}

func TestPdf(t *testing.T) {
	buff, err := os.Open("./tmp/sample-0.pdf")

	if err != nil {
		t.Error("error: PDF file not found")
	}

	defer func() {
		buff.Close()
	}()

	valid := IsPdf(buff)
	if !valid {
		t.Error("error: buffer not valid PDF file")
	}
}

func TestMpg(t *testing.T) {
	buff := bytes.NewReader(MPG_0)
	valid := IsMpg(buff)
	if !valid {
		t.Error("error: buffer not valid MPG file")
	}
	buff = bytes.NewReader(MPG_1)
	valid = IsMpg(buff)
	if !valid {
		t.Error("error: buffer not valid MPG file")
	}
}

func TestIsGif(t *testing.T) {
	buff, err := os.Open("./tmp/sample-0.gif")

	if err != nil {
		t.Error("error: GIF file not found")
	}

	defer func() {
		buff.Close()
	}()

	valid := IsGif(buff)
	if !valid {
		t.Error("error: buffer not valid GIF file")
	}
}

func TestIsBmp(t *testing.T) {
	buff, err := os.Open("./tmp/sample-0.bmp")

	if err != nil {
		t.Error("error: BMP file not found")
	}

	defer func() {
		buff.Close()
	}()

	valid := IsBmp(buff)
	if !valid {
		t.Error("error: buffer not valid BMP file")
	}
}

func TestIsDib(t *testing.T) {
	buff, err := os.Open("./tmp/sample-0.dib")

	if err != nil {
		t.Error("error: DIB file not found")
	}

	defer func() {
		buff.Close()
	}()

	valid := IsDib(buff)
	if !valid {
		t.Error("error: buffer not valid DIB file")
	}
}

func TestIsTiff(t *testing.T) {
	buff, err := os.Open("./tmp/sample-0.tiff")

	if err != nil {
		t.Error("error: TIFF file not found")
	}

	defer func() {
		buff.Close()
	}()

	valid := IsTiff(buff)
	if !valid {
		t.Error("error: buffer not valid TIFF file")
	}
}

func TestIsGzip(t *testing.T) {
	buff, err := os.Open("./tmp/sample-0.tgz")

	if err != nil {
		t.Error("error: GZIP file not found")
	}

	defer func() {
		buff.Close()
	}()

	valid := IsGzip(buff)
	if !valid {
		t.Error("error: buffer not valid GZIP file")
	}
}

func TestIsRar(t *testing.T) {
	buff, err := os.Open("./tmp/sample-0.rar")

	if err != nil {
		t.Error("error: RAR file not found")
	}

	defer func() {
		buff.Close()
	}()

	valid := IsRar(buff)
	if !valid {
		t.Error("error: buffer not valid RAR file")
	}
}

func TestIsMkv(t *testing.T) {
	buff, err := os.Open("./tmp/sample-0.mkv")

	if err != nil {
		t.Error("error: MKV file not found")
	}

	defer func() {
		buff.Close()
	}()

	valid := IsMkv(buff)
	if !valid {
		t.Error("error: buffer not valid MKV file")
	}
}

func TestIsJar(t *testing.T) {
	buff, err := os.Open("./tmp/sample-0.jar")

	if err != nil {
		t.Error("error: JAR file not found")
	}

	defer func() {
		buff.Close()
	}()

	valid := IsJar(buff)
	if !valid {
		t.Error("error: buffer not valid JAR file")
	}
}

func TestIsSwf(t *testing.T) {
	buff, err := os.Open("./tmp/sample-0.swf")

	if err != nil {
		t.Error("error: SWF file not found")
	}

	defer func() {
		buff.Close()
	}()

	valid := IsSwf(buff)
	if !valid {
		t.Error("error: buffer not valid SWF file")
	}
}

func TestIsFlv(t *testing.T) {
	buff, err := os.Open("./tmp/sample-0.flv")

	if err != nil {
		t.Error("error: FLV file not found")
	}

	defer func() {
		buff.Close()
	}()

	valid := IsFlv(buff)
	if !valid {
		t.Error("error: buffer not valid FLV file")
	}
}

func TestIsMp3(t *testing.T) {
	buff, err := os.Open("./tmp/sample-0.mp3")

	if err != nil {
		t.Error("error: MP3 file not found")
	}

	defer func() {
		buff.Close()
	}()

	valid := IsMp3(buff)
	if !valid {
		t.Error("error: buffer not valid MP3 file")
	}
}

func TestIsAvif(t *testing.T) {
	buff, err := os.Open("./tmp/sample-0.avif")

	if err != nil {
		t.Error("error: AVIF file not found")
	}

	defer func() {
		buff.Close()
	}()

	valid := IsAvif(buff)
	if !valid {
		t.Error("error: buffer not valid AVIF file")
	}
}

func TestIsMsOffice(t *testing.T) {
	buff, err := os.Open("./tmp/sample-0.docx")

	if err != nil {
		t.Error("error: MS Office file not found")
	}

	defer func() {
		buff.Close()
	}()

	valid := IsMsOffice(buff)
	if !valid {
		t.Error("error: buffer not valid MS Office file")
	}
}

func TestIsWebp(t *testing.T) {
	buff, err := os.Open("./tmp/sample-0.webp")

	if err != nil {
		t.Error("error: WEBP file not found")
	}

	defer func() {
		buff.Close()
	}()

	valid := IsWebp(buff)
	if !valid {
		t.Error("error: buffer not valid WEBP file")
	}
}

func TestIs3gp(t *testing.T) {
	buff, err := os.Open("./tmp/sample-0.3gp")

	if err != nil {
		t.Error("error: 3GP file not found")
	}

	defer func() {
		buff.Close()
	}()

	valid := Is3gp(buff)
	if !valid {
		t.Error("error: buffer not valid 3GP file")
	}
}

func TestIsZip(t *testing.T) {
	buff, err := os.Open("./tmp/sample-0.zip")

	if err != nil {
		t.Error("error: ZIP file not found")
	}

	defer func() {
		buff.Close()
	}()

	valid := IsZip(buff)
	if !valid {
		t.Error("error: buffer not valid ZIP file")
	}
}

func TestIsMp4(t *testing.T) {
	buff := bytes.NewReader(make([]byte, 0))
	valid := IsMp4(buff)
	if valid {
		t.Error("error: invalid MP4 buffer length")
	}

	realBuff, err := os.Open("./tmp/sample-0.mp4")

	if err != nil {
		t.Error("error: APK file not found")
	}

	defer func() {
		realBuff.Close()
	}()

	valid = IsMp4(realBuff)
	if !valid {
		t.Error("error: buffer not valid MP4 file")
	}
	buff = bytes.NewReader(make([]byte, 9))
	valid = IsMp4(buff)
	if valid {
		t.Error("error: buffer not valid MP4 file")
	}
}

func TestIsApk(t *testing.T) {
	buff, err := os.Open("./tmp/sample-0.apk")

	if err != nil {
		t.Error("error: APK file not found")
	}

	defer func() {
		buff.Close()
	}()

	valid := IsApk(buff)
	if !valid {
		t.Error("error: buffer not valid APK file")
	}
}

func TestGenericCompareBuffer(t *testing.T) {
	buff := bytes.NewReader([]byte{})
	valid := genericCompareBuffer(buff, PDF)
	if valid {
		t.Error("error: buffer is empty")
	}
}

func TestGenericMultipleCompareBuffer(t *testing.T) {
	buff := bytes.NewReader([]byte{})
	valid := genericMultipleCompareBuffer(buff, [][]byte{ZIP_0, ZIP_1, ZIP_2})
	if valid {
		t.Error("error: buffer is empty")
	}
	buff = bytes.NewReader(ZIP_0)
	valid = genericMultipleCompareBuffer(buff, [][]byte{{0x00}, ZIP_1, ZIP_2})
	if valid {
		t.Error("error: one of buffer type is invalid")
	}
}
