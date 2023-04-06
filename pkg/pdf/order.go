package pdfmaker

import (
	"fmt"
	"os"
	"time"

	"github.com/Invan2/invan_order_service/config"
	"github.com/Invan2/invan_order_service/models"
	"github.com/pkg/errors"
)

func (p *PdFmaker) MakeOrderHTML(order *models.CreatedOrderPDFRequest) (string, error) {

	var (
		lang   = "en"
		fields = map[string]string{
			"shop name": order.Shop.Title,
			"datetime":  time.Now().Format(config.DateTimeFormat),
			"seller":    fmt.Sprintf("%s %s", order.CreatedBy.FirstName, order.CreatedBy.LastName),
			"cashier":   fmt.Sprintf("%s %s", order.CreatedBy.FirstName, order.CreatedBy.LastName),
			"customer":  fmt.Sprintf("%s %s", order.Client.FirstName, order.Client.LastName),
			"contacts":  "+998935559562",
		}

		footerFields map[string]interface{} = map[string]interface{}{
			"SubTotal": fmt.Sprintf("%.2f %s", order.TotalPirce, "UZS"),
			"Discount": fmt.Sprintf("%.2f %s", order.TotalDiscountPrice, "UZS"),
			"Total":    fmt.Sprintf("%.2f %s", order.TotalPirce, "UZS"),
		}
	)

	html := `
		<!DOCTYPE html>
		<html lang="en">

		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Document</title>
			<style>
				* {
					margin: 0;
					padding: 0;
					box-sizing: border-box;
					font-family: sans-serif;
				}

				body {
					background: rgba(0,0,0,0.1);
				}

				.cheque {
					width: 80mm;
					background: #fff;
					padding: 18px;
				}

				.logo {
					display: block;
					margin: 0 auto;
					width: 80%;
				}

				.content {
					margin: 20px 0;
					padding-top: 20px;
					padding-bottom: 5px;
					border-top: 1px solid rgba(0,0,0,0.1);
					border-bottom: 1px solid rgba(0,0,0,0.1);
				}

				.flex {
					display: flex;
					justify-content: space-between;
					font-size: 14px;
					margin-bottom: 12px;
				}

				.label {
					font-weight: 600;
				}

				.barcode {
					width: 100%;
				}

				.products {
					display: grid;
					gap: 8px;
					margin-bottom: 15px;
					padding-bottom: 15px;
					border-bottom: 1px solid rgba(0,0,0,0.1)
				}

				.normal-text {
					display: block;
					font-weight: 400;
					margin-top: 5px;
				}

				.prices {
					display: grid;
					gap: 0px;
					padding-bottom: 10px;
				}
			</style>
			
			<script src="https://cdn.jsdelivr.net/gh/davidshimjs/qrcodejs/qrcode.min.js"></script>
		</head>
		<body>
	`

	html += `		<div class="cheque">`

	// Add Logo
	if order.Cheque.Logo != nil {
		html += fmt.Sprintf(
			`
        	<div alt="logo" class="logo" style="padding: %.fpx %.fpx %.fpx %.fpx;">
        	    <img width="100%%" onerror='this.style.display="none"' src="%s">
        	</div>
 		`,
			order.Cheque.Logo.Top,
			order.Cheque.Logo.Right,
			order.Cheque.Logo.Bottom,
			order.Cheque.Logo.Left,
			fmt.Sprintf("https://%s/%s/%s", p.cfg.MinioEndpoint, "file", order.Cheque.Logo.Image),
		)
	}
	// Add hozirontal rule
	html += `<hr />`

	// Add Headers
	headers := ""
	for _, field := range order.Cheque.Blocks[0].Fields {
		headers += fmt.Sprintf(
			`
				<div class="content">
					<div class="flex">
						<p class="label">%s</p>
						<p>%s</p>
					</div>
				</div>
			`,
			fields[field.NameTranslation[lang]],
			field.NameTranslation[lang],
		)
	}

	// Add hozirontal rule
	html += headers + `<hr />`

	// Add items
	items := ""
	for i, item := range order.Items {

		items += fmt.Sprintf(
			`
			<div class="products">
				<div class="product">
					<p class="label">%d. %s</p>
			 		<span class="normal-text">
						%.2f %s %.2f (%.2f %s)
					</span>
				</div>
			</div>
			`,
			i+1,
			item.ProductName,
			item.Value,
			item.MeasurementUnit.ShortName,
			item.Price,
			item.Value*item.Price,
			"UZS",
		)
	}

	// Add hozirontal rule
	html += items + `<hr />`

	// Add footer
	footers := ""
	for key, value := range footerFields {
		footers += fmt.Sprintf(
			`
			<div class="prices">
				<div class="flex">
					<p class="label">%s:</p>
					<p>%v</p>
				</div>
			</div>
			`,
			key,
			value,
		)
	}

	// Add hozirontal rule
	html += footers + `<hr />`

	// Add QR code
	if order.QRCode != "" {
		html += `<div class="barcode" id="barcode"></div>`

		html += fmt.Sprintf(
			`
			<script>
				const qrcode = new QRCode(document.getElementById('barcode'), {
					text: '%s',
					width: 128,
					height: 128,
					colorDark: '#000',
					colorLight: '#fff',
					correctLevel: QRCode.CorrectLevel.H
					});
			</script>
			
				`,
			order.QRCode,
		)
	}

	html += `
		</div>
		</body>
		</html>
	
	`

	outputPath := fmt.Sprintf("./%s.html", order.Id)

	f, err := os.Create(outputPath)
	if err != nil {
		return "", errors.Wrap(err, "error while create html file")
	}
	defer f.Close()

	_, err = f.WriteString(html)
	if err != nil {
		return "", errors.Wrap(err, "error while write html")
	}

	return outputPath, nil
}
