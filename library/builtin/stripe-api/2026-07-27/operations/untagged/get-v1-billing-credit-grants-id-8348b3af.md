---
title: Retrieve a credit grant
page_id: operation-get-v1-billing-credit-grants-id-8e8c0ff3
path: operations/untagged
description: <p>Retrieves a credit grant.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/billing/credit_grants/{id}
operation_ids:
    - GetBillingCreditGrantsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a credit grant

`GET /v1/billing/credit_grants/{id}`

Operation ID: `GetBillingCreditGrantsId`

<p>Retrieves a credit grant.</p>

## Definition

```yaml
{"summary": "Retrieve a credit grant", "description": "<p>Retrieves a credit grant.</p>", "operationId": "GetBillingCreditGrantsId", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "id", "in": "path", "description": "Unique identifier for the object.", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billing.credit_grant"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
