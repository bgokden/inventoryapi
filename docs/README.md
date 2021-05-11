# Documentation for Inventory API

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *http://localhost/v0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**getInventory**](Apis/DefaultApi.md#getinventory) | **GET** /inventory | Get current inventory
*DefaultApi* | [**listProductStocks**](Apis/DefaultApi.md#listproductstocks) | **GET** /productstock | Lists products with stock
*DefaultApi* | [**listProducts**](Apis/DefaultApi.md#listproducts) | **GET** /products | Lists products
*DefaultApi* | [**sellFromInventory**](Apis/DefaultApi.md#sellfrominventory) | **POST** /sell | Sell specified products and update Inventory
*DefaultApi* | [**upsertInventory**](Apis/DefaultApi.md#upsertinventory) | **POST** /inventory | Inserts or Updates stocks in Inventory
*DefaultApi* | [**upsertProducts**](Apis/DefaultApi.md#upsertproducts) | **POST** /products | Insert or Update products


<a name="documentation-for-models"></a>
## Documentation for Models

 - [Article](./Models/Article.md)
 - [Inventory](./Models/Inventory.md)
 - [Message](./Models/Message.md)
 - [Order](./Models/Order.md)
 - [Product](./Models/Product.md)
 - [ProductStock](./Models/ProductStock.md)
 - [ProductStocks](./Models/ProductStocks.md)
 - [Products](./Models/Products.md)
 - [SellOrder](./Models/SellOrder.md)
 - [Stock](./Models/Stock.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

All endpoints do not require authorization.
