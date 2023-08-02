package domain

import "fmt"

func getMailContent(summary Summary) string {

	subject := "Subject: Transaction Summary Information\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	message := fmt.Sprintf(`
<html>
	<head>
		<style>
			table {
				color: #384967;
			}
			td {
				padding: 5px;
			}
			tr.border-bottom > td {
				border-bottom: 1px solid #E7E7E7;
			}
		</style>
	</head>
	<body>
		<h2 style="color: #384967;">Transaction Summary Information</h2>
		<table>
			<tr class="border-bottom">
				<td><strong>Total balance is</strong></td>
				<td><strong style="font-size: 18px;">$%.2f</strong></td>
			</tr>
			<tr class="border-bottom">
				<td><strong>Number of transactions in July:</strong></td>
				<td>%d</td>
			</tr>
			<tr class="border-bottom">
				<td><strong>Number of transactions in August:</strong></td>
				<td>%d</td>
			</tr>
			<tr class="border-bottom">
				<td><strong>Average debit amount:</strong></td>
				<td>$%.2f</td>
			</tr>
			<tr class="border-bottom">
				<td><strong>Average credit amount:</strong></td>
				<td>$%.2f</td>
			</tr>
		</table>
		<br/>
		<img src="https://seeklogo.com/images/S/stori-logo-E64EA39237-seeklogo.com.png" alt="Logo" width="150" height="47">
		<br/>
		<p><i>Tarjeta de crédito sin Buró</i></p>
		<a href="https://www.storicard.com">https://www.storicard.com</a>
	</body>
</html>
`, summary.TotalBalance, summary.Months["July 2023"], summary.Months["August 2023"], summary.AvgDebitAmount, summary.AvgCreditAmount)

	return subject + mime + message
}
