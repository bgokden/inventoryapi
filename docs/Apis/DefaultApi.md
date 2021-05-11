# DefaultApi

All URIs are relative to *http://localhost/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getInventory**](DefaultApi.md#getInventory) | **GET** /inventory | Get current inventory
[**listProductStocks**](DefaultApi.md#listProductStocks) | **GET** /productstock | Lists products with stock
[**listProducts**](DefaultApi.md#listProducts) | **GET** /products | Lists products
[**sellFromInventory**](DefaultApi.md#sellFromInventory) | **POST** /sell | Sell specified products and update Inventory
[**upsertInventory**](DefaultApi.md#upsertInventory) | **POST** /inventory | Inserts or Updates stocks in Inventory
[**upsertProducts**](DefaultApi.md#upsertProducts) | **POST** /products | Insert or Update products


<a name="getInventory"></a>
# **getInventory**
> Inventory getInventory()

Get current inventory

    Get current inventory

### Parameters
This endpoint does not need any parameter.

### Return type

[**Inventory**](../Models/Inventory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listProductStocks"></a>
# **listProductStocks**
> ProductStocks listProductStocks()

Lists products with stock

    Lists products with stock

### Parameters
This endpoint does not need any parameter.

### Return type

[**ProductStocks**](../Models/ProductStocks.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listProducts"></a>
# **listProducts**
> Products listProducts()

Lists products

    List products

### Parameters
This endpoint does not need any parameter.

### Return type

[**Products**](../Models/Products.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="sellFromInventory"></a>
# **sellFromInventory**
> Message sellFromInventory(body)

Sell specified products and update Inventory

    Sell specified products and update Inventory

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SellOrder**](../Models/SellOrder.md)| Sell Order for products |

### Return type

[**Message**](../Models/Message.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="upsertInventory"></a>
# **upsertInventory**
> Message upsertInventory(body)

Inserts or Updates stocks in Inventory

    Inserts or Updates stocks in Inventory

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Inventory**](../Models/Inventory.md)| List of stocks |

### Return type

[**Message**](../Models/Message.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="upsertProducts"></a>
# **upsertProducts**
> Message upsertProducts(body)

Insert or Update products

    Inserts or Updates products

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Products**](../Models/Products.md)| List of products |

### Return type

[**Message**](../Models/Message.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

