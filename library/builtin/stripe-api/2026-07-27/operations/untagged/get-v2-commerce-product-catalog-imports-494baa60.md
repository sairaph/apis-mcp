---
title: List Product Catalog Imports
page_id: operation-get-v2-commerce-product-catalog-imports-590e8fc2
path: operations/untagged
description: Returns a list of ProductCatalogImport objects.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v2/commerce/product_catalog/imports
operation_ids:
    - GetV2CommerceProductCatalogImports
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# List Product Catalog Imports

`GET /v2/commerce/product_catalog/imports`

Operation ID: `GetV2CommerceProductCatalogImports`

Returns a list of ProductCatalogImport objects.

## Definition

```yaml
{"summary": "List Product Catalog Imports", "description": "Returns a list of ProductCatalogImport objects.", "operationId": "GetV2CommerceProductCatalogImports", "parameters": [{"name": "created", "in": "query", "description": "Filter for objects created at the specified timestamp.\nMust be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.", "required": false, "style": "form", "schema": {"type": "string", "format": "date-time"}}, {"name": "created_gt", "in": "query", "description": "Filter for objects created after the specified timestamp.\nMust be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.", "required": false, "style": "form", "schema": {"type": "string", "format": "date-time"}}, {"name": "created_gte", "in": "query", "description": "Filter for objects created on or after the specified timestamp.\nMust be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.", "required": false, "style": "form", "schema": {"type": "string", "format": "date-time"}}, {"name": "created_lt", "in": "query", "description": "Filter for objects created before the specified timestamp.\nMust be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.", "required": false, "style": "form", "schema": {"type": "string", "format": "date-time"}}, {"name": "created_lte", "in": "query", "description": "Filter for objects created on or before the specified timestamp.\nMust be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.", "required": false, "style": "form", "schema": {"type": "string", "format": "date-time"}}, {"name": "feed_type", "in": "query", "description": "Filter by the type of feed data being imported.", "required": false, "style": "form", "schema": {"type": "string", "enum": ["inventory", "pricing", "product", "promotion"]}}, {"name": "limit", "in": "query", "description": "The maximum number of results per page.", "required": false, "style": "form", "schema": {"type": "integer"}}, {"name": "page", "in": "query", "description": "The page token.", "required": false, "style": "form", "schema": {"type": "string"}}, {"name": "status", "in": "query", "description": "Filter by import status.", "required": false, "style": "form", "schema": {"type": "string", "enum": ["awaiting_upload", "failed", "processing", "succeeded", "succeeded_with_errors"]}}], "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"required": ["data", "next_page_url", "previous_page_url"], "type": "object", "properties": {"data": {"type": "array", "description": "List of ProductCatalogImport objects.", "items": {"$ref": "#/components/schemas/v2.commerce.product_catalog_import"}}, "next_page_url": {"type": "string", "description": "URL to fetch the next page of the list. If there are no more pages, the value is null.", "nullable": true}, "previous_page_url": {"type": "string", "description": "URL to fetch the previous page of the list. If there are no previous pages, the value is null.", "nullable": true}}}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
