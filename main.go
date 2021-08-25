package main

import (
	// for reading and parsing from json to struct
	"encoding/json"
	//print and sending response
	"fmt"
	// to show some errors
	"io/ioutil"
	// to make the server
	"net/http"
	// to work with "string" type, like replacing, converting, ...
	"strconv"
	"strings"
)

// Products struct which contains an array of products
type Products struct {
	Products []Product `json:"products"`
}

// Product struct which contains an id, a name, an image, and ...
type Product struct {
	ID          int    `json:"id"`
	ProductName string `json:"productName"`
	From        string `json:"from"`
	Image       string `json:"image"`
	Nutrients   string `json:"nutrients"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
	Organic     bool   `json:"organic"`
	Description string `json:"description"`
}

func main() {
	// Read html file templates
	overviewByte, err := ioutil.ReadFile("./templates/template-overview.html")
	if err != nil {
		fmt.Print(err)
	}
	cardByte, err := ioutil.ReadFile("./templates/template-card.html")
	if err != nil {
		fmt.Print(err)
	}
	productByte, err := ioutil.ReadFile("./templates/template-product.html")
	if err != nil {
		fmt.Print(err)
	}
	// convert content(array of bytes) to a 'string'
	tempOverview := string(overviewByte)
	tempCard := string(cardByte)
	tempProduct := string(productByte)

	// read our opened jsonFile as a byte array.
	file, _ := ioutil.ReadFile("./dev-data/data.json")

	// initialize Products array
	var products Products
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'products' which we defined above
	json.Unmarshal([]byte(file), &products)

	// Route handler for index page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var cards []string
		for i := 0; i < len(products.Products); i++ {
			cards = append(cards, replaceTemplate(tempCard, products.Products[i]))
		}
		cardsHtml := strings.Join(cards, "")
		output := strings.ReplaceAll(tempOverview, `{%PRODUCT_CARDS%}`, cardsHtml)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// to send response
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, output)
	})

	// Product page
	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		if productID, err := strconv.Atoi(r.URL.Query().Get("id")); err == nil {

			product := getProductByID(products.Products, productID)
			// inject data
			output := replaceTemplate(tempProduct, product)

			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			// to send response
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, output)
		}
	})

	http.ListenAndServe(":4000", nil)

}

// to replace products in temp
func replaceTemplate(temp string, product Product) string {
	temp = strings.ReplaceAll(temp, `{%PRODUCTNAME%}`, product.ProductName)
	temp = strings.ReplaceAll(temp, `{%IMAGE%}`, product.Image)
	temp = strings.ReplaceAll(temp, `{%PRICE%}`, product.Price)
	temp = strings.ReplaceAll(temp, `{%FROM%}`, product.From)
	temp = strings.ReplaceAll(temp, `{%NUTRIENTS%}`, product.Nutrients)
	temp = strings.ReplaceAll(temp, `{%QUANTITY%}`, product.Quantity)
	temp = strings.ReplaceAll(temp, `{%DESCRIPTION%}`, product.Description)
	temp = strings.ReplaceAll(temp, `{%ID%}`, strconv.Itoa(product.ID))

	if !product.Organic {
		temp = strings.ReplaceAll(temp, `{%NOT_ORGANIC%}`, "not-organic")
	}

	return temp
}

// to find the prodoct with "id" in list of products
func getProductByID(products []Product, id int) Product {
	for i := range products {
		if products[i].ID == id {
			return products[i]
		}
	}
	return products[0]
}
