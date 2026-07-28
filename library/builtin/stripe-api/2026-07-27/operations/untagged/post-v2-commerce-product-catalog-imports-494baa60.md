---
title: Create a Product Catalog Import
page_id: operation-post-v2-commerce-product-catalog-imports-fb10455c
path: operations/untagged
description: Creates a ProductCatalogImport.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v2/commerce/product_catalog/imports
operation_ids:
    - PostV2CommerceProductCatalogImports
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a Product Catalog Import

`POST /v2/commerce/product_catalog/imports`

Operation ID: `PostV2CommerceProductCatalogImports`

Creates a ProductCatalogImport.

## Definition

```yaml
{"summary": "Create a Product Catalog Import", "description": "Creates a ProductCatalogImport.", "operationId": "PostV2CommerceProductCatalogImports", "requestBody": {"content": {"application/json": {"schema": {"required": ["feed_type", "metadata", "mode"], "type": "object", "properties": {"feed_type": {"type": "string", "description": "The type of catalog data to import.", "enum": ["inventory", "pricing", "product", "promotion"]}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Additional information about the import in a structured format."}, "mode": {"type": "string", "description": "The strategy for handling existing catalog data during import.", "enum": ["replace", "upsert"]}}}}}}, "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/v2.commerce.product_catalog_import"}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
