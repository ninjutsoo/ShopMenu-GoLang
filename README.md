# ShopMenu-GoLang

### Project overview
I have done something like this before with nodejs and in this product I converted the basis to Go.
the purpose of project is to show a list of products in "index page" on a local server and by clicking 
on them you can see a little discription and details about that product (fruit in this case).
since go is widely used for web/server programmers in this project we worked on this local server to use 
go properties.

### functions and structures
- type Product struct : here we make an struct for each product contaiting their properties (id, name, image, quantity, ...)
- http.HandleFunc : it is route handler for index page (where we have list of all fruits) and product page (where we have details of fruit). 
                    in first HandleFunc we inject each data from json file to template-card, then we inject all to template-overview.
                    in second HandleFunc we should make Pruduct page so based on the url that we are in, we extract the product id and based            
                    on that we can inject it in template-product and show it.
- replaceTemplate : main funciton used several time to inject data from json file to html file.
- getProductID : we use it to find the product in list of products based on id (used in http.HandleFunc).

### html files
3 html files each one to show some data, 
    template-card : to show each card's data.
    template-overview : to show list of products in index page.
    template-product : to show each product properties in product page.

### json files 
we basically here have all datas that we want to show on html pages, you can add and remove to the list easily.

### how to run
1- copy the file in your GOPATH directory
2- go to the directory useing cmd (or ide's terminal)
3- type: go run main.go
4- open your browser and type in: localhoset:4000
enjoy :)

also you can see images of the project in /img directory.
