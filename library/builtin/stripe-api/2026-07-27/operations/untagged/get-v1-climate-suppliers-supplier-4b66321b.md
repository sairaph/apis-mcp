---
title: Retrieve a supplier
page_id: operation-get-v1-climate-suppliers-supplier-af9ac7f0
path: operations/untagged
description: <p>Retrieves a Climate supplier object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/climate/suppliers/{supplier}
operation_ids:
    - GetClimateSuppliersSupplier
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a supplier

`GET /v1/climate/suppliers/{supplier}`

Operation ID: `GetClimateSuppliersSupplier`

<p>Retrieves a Climate supplier object.</p>

## Definition

```yaml
{"summary": "Retrieve a supplier", "description": "<p>Retrieves a Climate supplier object.</p>", "operationId": "GetClimateSuppliersSupplier", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "supplier", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/climate.supplier"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
