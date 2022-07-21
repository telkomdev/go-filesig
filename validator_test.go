package filesig

import (
	"bytes"
	"testing"
)

func TestIsOneOf(t *testing.T) {
	buff := bytes.NewReader([]byte("%PDF-1.3"))
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
	buff := bytes.NewReader([]byte("%PDF-1.3"))
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
	buff := bytes.NewReader(GIF)
	valid := IsGif(buff)
	if !valid {
		t.Error("error: buffer not valid GIF file")
	}
}

func TestIsBmp(t *testing.T) {
	buff := bytes.NewReader(BMP)
	valid := IsBmp(buff)
	if !valid {
		t.Error("error: buffer not valid BMP file")
	}
}

func TestIsDib(t *testing.T) {
	buff := bytes.NewReader(DIB)
	valid := IsDib(buff)
	if !valid {
		t.Error("error: buffer not valid DIB file")
	}
}

func TestIsTiff(t *testing.T) {
	buff := bytes.NewReader(TIFF)
	valid := IsTiff(buff)
	if !valid {
		t.Error("error: buffer not valid TIFF file")
	}
}

func TestIsGzip(t *testing.T) {
	buff := bytes.NewReader(GZIP)
	valid := IsGzip(buff)
	if !valid {
		t.Error("error: buffer not valid GZIP file")
	}
}

func TestIsRar(t *testing.T) {
	buff := bytes.NewReader(RAR)
	valid := IsRar(buff)
	if !valid {
		t.Error("error: buffer not valid RAR file")
	}
}

func TestIsMkv(t *testing.T) {
	buff := bytes.NewReader(MKV)
	valid := IsMkv(buff)
	if !valid {
		t.Error("error: buffer not valid MKV file")
	}
}

func TestIsJar(t *testing.T) {
	buff := bytes.NewReader(JAR)
	valid := IsJar(buff)
	if !valid {
		t.Error("error: buffer not valid JAR file")
	}
}

func TestIsSwf(t *testing.T) {
	buff := bytes.NewReader(SWF)
	valid := IsSwf(buff)
	if !valid {
		t.Error("error: buffer not valid SWF file")
	}
}

func TestIsFlv(t *testing.T) {
	buff := bytes.NewReader(FLV)
	valid := IsFlv(buff)
	if !valid {
		t.Error("error: buffer not valid FLV file")
	}
}

func TestIsMp3(t *testing.T) {
	buff := bytes.NewReader(MP3)
	valid := IsMp3(buff)
	if !valid {
		t.Error("error: buffer not valid MP3 file")
	}
}

func TestIsAvif(t *testing.T) {
	buff := bytes.NewReader(AVIF)
	valid := IsAvif(buff)
	if !valid {
		t.Error("error: buffer not valid AVIF file")
	}
}

func TestIsMsOffice(t *testing.T) {
	buff := bytes.NewReader(MS_OFFICE)
	valid := IsMsOffice(buff)
	if !valid {
		t.Error("error: buffer not valid MS Office file")
	}
}

func TestIsWebp(t *testing.T) {
	buff := bytes.NewReader(WEBP)
	valid := IsWebp(buff)
	if !valid {
		t.Error("error: buffer not valid WEBP file")
	}
}

func TestIs3gp(t *testing.T) {
	buff := bytes.NewReader(THREE_GP_0)
	valid := Is3gp(buff)
	if !valid {
		t.Error("error: buffer not valid 3GP file")
	}
}

func TestIsZip(t *testing.T) {
	buff := bytes.NewReader(ZIP_0)
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
	buff = bytes.NewReader(MP4_0)
	valid = IsMp4(buff)
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
	buff := bytes.NewReader(APK)
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
