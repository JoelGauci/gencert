package pdf

import (
	"fmt"
	"os"
	"path"

	"github.com/jung-kurt/gofpdf"
	"training.go/gencert/cert"
)

type PdfSaver struct {
	OutputDir string
}

func New(outputdir string) (*PdfSaver, error) {
	var p *PdfSaver
	err := os.MkdirAll(outputdir, os.ModePerm)
	if err != nil {
		return p, nil
	}

	p = &PdfSaver{
		OutputDir: outputdir,
	}

	return p, nil
}

func (p *PdfSaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")

	// Title
	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	// Background
	background(pdf)

	// --
	// Header
	Header(pdf, &cert)
	pdf.Ln(30)

	// --
	// Body
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelPresented, "C")
	pdf.Ln(30)

	// Body - Student name
	pdf.SetFont("Times", "B", 40)
	pdf.WriteAligned(0, 50, cert.Name, "C")
	pdf.Ln(30)

	// Body - Participation
	pdf.SetFont("Helvetica", "", 20)
	pdf.WriteAligned(0, 50, cert.LabelParticipation, "C")
	pdf.Ln(30)

	// Body - Date
	pdf.SetFont("Helvetica", "I", 15)
	date := fmt.Sprintf("Date: %v", cert.Date.Format("02/01/2006"))
	pdf.WriteAligned(0, 50, date, "C")

	// --
	// footer
	Footer(pdf)

	// save file
	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	fmt.Printf("Saved certificate to '%v'\n", path)
	return nil
}

func background(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	pageWidth, pageHeight := pdf.GetPageSize()
	pdf.ImageOptions("img/background.png",
		0, 0, // x, y
		pageWidth, pageHeight, // width, height
		false, opts, 0, "")
}

func Header(pdf *gofpdf.Fpdf, c *cert.Cert) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	margin := 30.0
	x := 0.0
	imageWidth := 30.0
	filenameGCloud := "img/google-cloud.png"
	pdf.ImageOptions(filenameGCloud,
		x+margin, 20, // x, y
		imageWidth, 0, // width, height
		false, opts, 0, "")

	pageWidth, _ := pdf.GetPageSize()
	imageWidthApigee := 35.0
	filenameApigee := "img/apigee.png"
	x = pageWidth - imageWidthApigee
	pdf.ImageOptions(filenameApigee,
		x-margin, 25, // x, y
		imageWidthApigee, 0, // width, height
		false, opts, 0, "")

	pdf.SetFont("Helvetica", "", 40)
	pdf.WriteAligned(0, 50, c.LabelCompletion, "C")
}

func Footer(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	imageWidth := 50.0
	filename := "img/certificate.png"

	pageWidth, pageHeight := pdf.GetPageSize()
	x := pageWidth - imageWidth - 20.0
	y := pageHeight - imageWidth - 10.0
	pdf.ImageOptions(filename,
		x, y,
		imageWidth, 0, // width, height
		false, opts, 0, "")

}
