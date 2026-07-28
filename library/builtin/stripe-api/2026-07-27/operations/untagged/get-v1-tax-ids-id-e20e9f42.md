---
title: Retrieve a tax ID
page_id: operation-get-v1-tax-ids-id-49b5e0ee
path: operations/untagged
description: <p>Retrieves an account or customer <code>tax_id</code> object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/tax_ids/{id}
operation_ids:
    - GetTaxIdsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a tax ID

`GET /v1/tax_ids/{id}`

Operation ID: `GetTaxIdsId`

<p>Retrieves an account or customer <code>tax_id</code> object.</p>

## Definition

```yaml
{"summary": "Retrieve a tax ID", "description": "<p>Retrieves an account or customer <code>tax_id</code> object.</p>", "operationId": "GetTaxIdsId", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tax_id"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
