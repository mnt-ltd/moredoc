package converter

import (
	"os/exec"
	"strings"
	"testing"
	"time"

	"go.uber.org/zap"
)

var (
	converter *Converter
	pdfFiles  = []string{
		"../../cache/one.pdf",
		"../../cache/two.pdf",
		"../../cache/three.pdf",
	}
)

func init() {
	logger, _ := zap.NewDevelopment()
	converter = NewConverter(logger)
	converter.SetCachePath("../../cache/convert")
}

func TestConvertToPDF(t *testing.T) {
	type args struct {
		src     string
		timeout []time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantDst string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDst, err := converter.ConvertToPDF(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToPDF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDst != tt.wantDst {
				t.Errorf("ConvertToPDF() = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}

func TestConvertOfficeToPDF(t *testing.T) {
	files := []string{
		"../../cache/xlsx.xlsx",
		"../../cache/ppt.ppt",
		"../../cache/docx.docx",
	}
	for _, file := range files {
		dst, err := converter.ConvertOfficeToPDF(file)
		if err != nil {
			t.Errorf("ConvertOfficeToPDF() error = %v， source file = %s", err, file)
			return
		}
		t.Log(dst)
	}
}

func TestConvertEPUBToPDF(t *testing.T) {
	type args struct {
		src     string
		timeout []time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantDst string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDst, err := converter.ConvertEPUBToPDF(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertEPUBToPDF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDst != tt.wantDst {
				t.Errorf("ConvertEPUBToPDF() = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}

func TestConvertUMDToPDF(t *testing.T) {
	type args struct {
		src     string
		timeout []time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantDst string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDst, err := converter.ConvertUMDToPDF(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertUMDToPDF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDst != tt.wantDst {
				t.Errorf("ConvertUMDToPDF() = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}

func TestConvertTXTToPDF(t *testing.T) {
	txtFile := "../../cache/txt.txt"
	dst, err := converter.ConvertTXTToPDF(txtFile)
	if err != nil {
		t.Errorf("ConvertTXTToPDF() error = %v， source file = %s", err, txtFile)
		return
	}

	t.Log(dst)
}

func TestConvertMOBIToPDF(t *testing.T) {
	type args struct {
		src     string
		timeout []time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantDst string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDst, err := converter.ConvertMOBIToPDF(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertMOBIToPDF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDst != tt.wantDst {
				t.Errorf("ConvertMOBIToPDF() = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}

func TestConvertPDFToTxt(t *testing.T) {
	for _, file := range pdfFiles {
		dst, err := converter.ConvertPDFToTxt(file)
		if err != nil {
			t.Errorf("ConvertPDFToTxt() error = %v， source file = %s", err, file)
			return
		}
		t.Log(dst)
	}
}

func TestConvertCHMToPDF(t *testing.T) {
	type args struct {
		src     string
		timeout []time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantDst string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDst, err := converter.ConvertCHMToPDF(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertCHMToPDF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDst != tt.wantDst {
				t.Errorf("ConvertCHMToPDF() = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}

func TestConvertPDFToSVG(t *testing.T) {
	for _, file := range pdfFiles {
		pages, err := converter.ConvertPDFToSVG(file, 0, 10000, false, true)
		if err != nil {
			t.Errorf("ConvertPDFToTxt() error = %v， source file = %s", err, file)
			return
		}
		t.Logf("file = %s", file)
		for page, pagePath := range pages {
			t.Log(file, page, pagePath)
		}
	}
}

func TestConvertPDFToPNG(t *testing.T) {
	for _, file := range pdfFiles {
		pages, err := converter.ConvertPDFToPNG(file, 0, 10000)
		if err != nil {
			t.Errorf("ConvertPDFToTxt() error = %v， source file = %s", err, file)
			return
		}
		t.Logf("file = %s", file)
		for page, pagePath := range pages {
			t.Log(file, page, pagePath)
		}
	}
}

func TestCountPDFPages(t *testing.T) {
	for _, file := range pdfFiles {
		now := time.Now()
		pages, err := converter.CountPDFPages(file)
		t.Log(time.Since(now))
		if err != nil {
			t.Errorf("CountPDFPages() error = %v， source file = %s", err, file)
			return
		}

		t.Logf("file = %s，pages = %d", file, pages)
	}
}

func TestExistCommand(t *testing.T) {
	s := "我是中国人"
	t.Log(strings.Count(s, "") - 1)
	t.Logf("calibre= %v", converter.ExistCalibre())
	t.Logf("svgo= %v", converter.ExistSVGO())
	t.Logf("mupdf= %v", converter.ExistMupdf())
	t.Logf("soffice= %v", converter.ExistSoffice())
	t.Log(exec.LookPath("soffice1"))
}
