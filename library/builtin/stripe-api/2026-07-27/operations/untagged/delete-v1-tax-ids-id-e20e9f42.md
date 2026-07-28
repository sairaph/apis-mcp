---
title: Delete a tax ID
page_id: operation-delete-v1-tax-ids-id-43f037b2
path: operations/untagged
description: <p>Deletes an existing account or customer <code>tax_id</code> object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v1/tax_ids/{id}
operation_ids:
    - DeleteTaxIdsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Delete a tax ID

`DELETE /v1/tax_ids/{id}`

Operation ID: `DeleteTaxIdsId`

<p>Deletes an existing account or customer <code>tax_id</code> object.</p>

## Definition

```yaml
{"summary": "Delete a tax ID", "description": "<p>Deletes an existing account or customer <code>tax_id</code> object.</p>", "operationId": "DeleteTaxIdsId", "parameters": [{"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/deleted_tax_id"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
