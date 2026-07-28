---
title: Retrieve a Product Catalog Import
page_id: operation-get-v2-commerce-product-catalog-imports-id-0abc844e
path: operations/untagged
description: Retrieves a ProductCatalogImport by ID.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v2/commerce/product_catalog/imports/{id}
operation_ids:
    - GetV2CommerceProductCatalogImportsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a Product Catalog Import

`GET /v2/commerce/product_catalog/imports/{id}`

Operation ID: `GetV2CommerceProductCatalogImportsId`

Retrieves a ProductCatalogImport by ID.

## Definition

```yaml
{"summary": "Retrieve a Product Catalog Import", "description": "Retrieves a ProductCatalogImport by ID.", "operationId": "GetV2CommerceProductCatalogImportsId", "parameters": [{"name": "id", "in": "path", "description": "The ID of the ProductCatalogImport to retrieve.", "required": true, "style": "simple", "schema": {"type": "string"}}], "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/v2.commerce.product_catalog_import"}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.resource_missing"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
