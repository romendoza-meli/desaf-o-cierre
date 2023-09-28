package internal

import (
	"database/sql"
	"encoding/json"
	"os"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

func LoadData(db *sql.DB) {
    customersPath := "./datos/customers.json"
    invoicesPath := "./datos/invoices.json"
    productsPath := "./datos/products.json"
    salesPath := "./datos/sales.json"
    var customers []domain.Customers
    var invoices []domain.Invoices
    var products []domain.Product
    var sales []domain.Sales
    customersData, err := os.ReadFile(customersPath)
    if err != nil {
        panic(err)
    }
    err = json.Unmarshal(customersData, &customers)
    if err != nil {
        panic(err)
    }
    for _, customer := range customers {
        _, err := db.Exec("INSERT INTO customers (id, first_name, last_name, `condition`) VALUES (?, ?, ?, ?)", customer.Id, customer.FirstName, customer.LastName, customer.Condition)
        if err != nil {
            panic(err)
        }
    }
    invoicesData, err := os.ReadFile(invoicesPath)
    if err != nil {
        panic(err)
    }
    err = json.Unmarshal(invoicesData, &invoices)
    if err != nil {
        panic(err)
    }
    for _, invoice := range invoices {
        _, err := db.Exec("INSERT INTO invoices (id, datetime, customer_id, total) VALUES (?, ?, ?, ?)", invoice.Id, invoice.Datetime, invoice.CustomerId, invoice.Total)
        if err != nil {
            panic(err)
        }
    }
    productsData, err := os.ReadFile(productsPath)
    if err != nil {
        panic(err)
    }
    err = json.Unmarshal(productsData, &products)
    if err != nil {
        panic(err)
    }
    for _, product := range products {
        _, err := db.Exec("INSERT INTO products (id, description, price) VALUES (?, ?, ?)", product.Id, product.Description, product.Price)
        if err != nil {
            panic(err)
        }
    }
    salesData, err := os.ReadFile(salesPath)
    if err != nil {
        panic(err)
    }
    err = json.Unmarshal(salesData, &sales)
    if err != nil {
        panic(err)
    }
    for _, sale := range sales {
        _, err := db.Exec("INSERT INTO sales (id, invoice_id, product_id, quantity) VALUES (?, ?, ?, ?)", sale.Id, sale.InvoicesId, sale.ProductId, sale.Quantity)
        if err != nil {
            panic(err)
        }
    }
}