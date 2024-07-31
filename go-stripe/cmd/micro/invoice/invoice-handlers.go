package main

import (
	"fmt"
	"github.com/go-pdf/fpdf"
	gofpdi2 "github.com/go-pdf/fpdf/contrib/gofpdi"
	"net/http"
	"time"
)

type Order struct {
	ID        int        `json:"id"`
	Quantity  int        `json:"quantity"`
	Amount    int        `json:"amount"`
	Product   string     `json:"product"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	Items     []Products `json:"products"`
}
type Products struct {
	Name     string `json:"name"`
	Amount   int    `json:"amount"`
	Quantity int    `json:"quantity"`
}

func (app *application) CreateAndSendInvoice(w http.ResponseWriter, r *http.Request) {
	// receive json
	var order Order
	err := app.readJSON(w, r, order)
	if err != nil {
		_ = app.badRequest(w, r, err)
		return
	}

	//order.ID = 100
	//order.Email = "me@here.com"
	//order.FirstName = "John"
	//order.LastName = "Smith"
	//order.Quantity = 1
	//order.Amount = 1000
	//order.Product = "Widget"
	//order.CreatedAt = time.Now()

	// generate a pdf invoice
	err = app.createInvoicePDF(order)
	if err != nil {
		_ = app.badRequest(w, r, err)
		return
	}

	// create mail attachments
	attachments := []string{
		fmt.Sprintf("./invoices/%d.pdf", order.ID),
	}

	// send mail with attachment
	err = app.SendMail("info@widgets.com", order.Email, "Your invoice", "invoice", attachments, nil)
	if err != nil {
		_ = app.badRequest(w, r, err)
		return
	}

	// send response
	var resp struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}
	resp.Error = false
	resp.Message = fmt.Sprintf("Order %d.pdf created and sent to %s", order.ID, order.Email)
	_ = app.writeJSON(w, http.StatusCreated, resp)
}

func (app *application) createInvoicePDF(order Order) error {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(10, 13, 10)
	pdf.SetAutoPageBreak(true, 0)

	importer := gofpdi2.NewImporter()
	t := importer.ImportPage(pdf, "./pdf-templates/Invoice.pdf", 1, "/MediaBox")
	pdf.AddPage()
	importer.UseImportedTemplate(pdf, t, 0, 0, 215.9, 0)

	//write info
	pdf.SetY(50)
	pdf.SetX(10)
	pdf.SetFont("Times", "", 11)
	pdf.CellFormat(97, 8, fmt.Sprintf("Attention: %s %s", order.FirstName, order.LastName), "", 0, "L", false, 0, "")
	pdf.Ln(5)
	pdf.CellFormat(97, 8, order.Email, "", 0, "L", false, 0, "")
	pdf.Ln(5)
	pdf.CellFormat(97, 8, order.CreatedAt.Format("2006-01-02"), "", 0, "L", false, 0, "")

	// range through products
	pdf.SetX(58)
	pdf.SetY(93)
	pdf.CellFormat(155, 8, order.Product, "", 0, "L", false, 0, "")
	pdf.SetX(166)
	pdf.CellFormat(20, 8, fmt.Sprintf("%d", order.Quantity), "", 0, "C", false, 0, "")
	pdf.SetX(185)
	pdf.CellFormat(20, 8, fmt.Sprintf("$%.2f", float32(order.Amount/100.0)), "", 0, "R", false, 0, "")

	invoicePath := fmt.Sprintf("./invoices/%d.pdf", order.ID)
	err := pdf.OutputFileAndClose(invoicePath)
	if err != nil {
		return err
	}
	return nil
}
