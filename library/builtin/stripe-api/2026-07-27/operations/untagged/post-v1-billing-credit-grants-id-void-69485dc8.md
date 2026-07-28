---
title: Void a credit grant
page_id: operation-post-v1-billing-credit-grants-id-void-6c610b19
path: operations/untagged
description: <p>Voids a credit grant.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/billing/credit_grants/{id}/void
operation_ids:
    - PostBillingCreditGrantsIdVoid
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Void a credit grant

`POST /v1/billing/credit_grants/{id}/void`

Operation ID: `PostBillingCreditGrantsIdVoid`

<p>Voids a credit grant.</p>

## Definition

```yaml
{"summary": "Void a credit grant", "description": "<p>Voids a credit grant.</p>", "operationId": "PostBillingCreditGrantsIdVoid", "parameters": [{"name": "id", "in": "path", "description": "Unique identifier for the object.", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billing.credit_grant"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
