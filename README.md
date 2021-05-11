# inventory api

This Api Serves an Inventory API to keep inventory of articles that Belogs to products.

It is possible to add or update inventory. 
Updating inventory change the name of the article and adds the given stock value to the current value.

It is possible to add or update products.
Updating a product changes the product articles.

It is possible to Lits product stocks and sell one or more products.

This application is a demonstor and intended to run in single instance.
All values kept in memory and useful for small datasets.

# Development

This application is written with go 1.15. Please install golang for development: [Golang Install](https://golang.org/doc/install)
Docker is needed for distribution packaging. Install docker for your platform: [Docker](https://docs.docker.com/get-docker/)

All commands are meant to run in the root of the project. There bash scripts for common commands:

## Run tests:

```bash
./test.sh
```

## Run build:

```bash
./build.sh
```

Application will be output into app file in the root of this project.
App can be run with optional part flag:

```bash
./app -port 8080
```

## Run build docker image:

```bash
./build_image.sh
```
Image name can be set in the script.
Docker image is built with distroless base, current size is `11.7MB`.
It is tested with Sync and no vulnerability found.


During development you can run the service with:

```bash
go run main.go
```

Default server runs on port 8080 and it is possible to override with `-port` parameter.

# API Driven Development

This project follows API Driven Design Approach. OpenAPI 3 spefication can be found as `./spec/api.yaml` file.

## Build API Source Code
```Bash
./buildspec.sh
```

## Documents generated with openapi-generator
```bash
openapi-generator generate -i ./spec/api.yaml -g markdown -o docs
```

This will generate source codes server, client, types and spec in `./api` folder.
These files includes default implementations and interfaces for server and types.
Also allows spec validation.

By following this practice, development follows in this order: 

- negotiate API changes with stakeholders

- update API Spec

- generate code and documents

- fix interface changes and tests


#### Example Curl methods to use when running on localhost

```bash
curl -v --header "Content-Type: application/json" \
  --request POST \
  --data '{"inventory":[{"art_id":"1","name":"leg","stock":"12"},{"art_id":"2","name":"screw","stock":"17"},{"art_id":"3","name":"seat","stock":"2"},{"art_id":"4","name":"table top","stock":"1"}]}' \
  http://localhost:8080/v0/inventory
```

```bash
curl -i -H "Accept: application/json" http://localhost:8080/v0/inventory
```

```bash
curl -v --header "Content-Type: application/json" \
  --request POST \
  --data '{"products":[{"name":"Dining Chair","contain_articles":[{"art_id":"1","amount_of":"4"},{"art_id":"2","amount_of":"8"},{"art_id":"3","amount_of":"1"}]},{"name":"Dinning Table","contain_articles":[{"art_id":"1","amount_of":"4"},{"art_id":"2","amount_of":"8"},{"art_id":"4","amount_of":"1"}]}]}' \
  http://localhost:8080/v0/products
```

```bash
curl -i -H "Accept: application/json"  http://localhost:8080/v0/products
```

```bash
curl -i -H "Accept: application/json"  http://localhost:8080/v0/productstock
```

```bash
curl -v --header "Content-Type: application/json" \
  --request POST \
  --data '{"orders":[{"product_name":"Dinning Table","number":1}]}' \
  http://localhost:8080/v0/sell
```

##### Input validation example:

This will return 400 and error message:

```bash
curl -v --header "Content-Type: application/json" \
  --request POST \
  --data '{"orders":[{"product_name":"Dinning Table","number":-1}]}' \
  http://localhost:8080/v0/sell

{"message":"request body has an error: doesn't match the schema: Error at \"/orders/0/number\": number must be at least 1"}
```


# TODO:

- Add more tests and edge cases.

- Add an actual database integration.

- Use unsigned integer to validate values instead of string in various places.

- For large dataset, paging will be needed in list endpoints.

- Write a kubernetes helm chart.







