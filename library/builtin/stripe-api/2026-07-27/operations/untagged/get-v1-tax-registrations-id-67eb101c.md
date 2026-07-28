---
title: Retrieve a registration
page_id: operation-get-v1-tax-registrations-id-8c3bd87b
path: operations/untagged
description: <p>Returns a Tax <code>Registration</code> object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/tax/registrations/{id}
operation_ids:
    - GetTaxRegistrationsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a registration

`GET /v1/tax/registrations/{id}`

Operation ID: `GetTaxRegistrationsId`

<p>Returns a Tax <code>Registration</code> object.</p>

## Definition

```yaml
{"summary": "Retrieve a registration", "description": "<p>Returns a Tax <code>Registration</code> object.</p>", "operationId": "GetTaxRegistrationsId", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tax.registration"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
