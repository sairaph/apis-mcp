---
title: Update a credit grant
page_id: operation-post-v1-billing-credit-grants-id-77923ff4
path: operations/untagged
description: <p>Updates a credit grant.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/billing/credit_grants/{id}
operation_ids:
    - PostBillingCreditGrantsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update a credit grant

`POST /v1/billing/credit_grants/{id}`

Operation ID: `PostBillingCreditGrantsId`

<p>Updates a credit grant.</p>

## Definition

```yaml
{"summary": "Update a credit grant", "description": "<p>Updates a credit grant.</p>", "operationId": "PostBillingCreditGrantsId", "parameters": [{"name": "id", "in": "path", "description": "Unique identifier for the object.", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "expires_at": {"description": "The time when the billing credits created by this credit grant expire. If set to empty, the billing credits never expire.", "anyOf": [{"type": "integer", "format": "unix-time"}, {"type": "string", "enum": [""]}]}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of key-value pairs you can attach to an object. You can use this to store additional information about the object (for example, cost basis) in a structured format."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "expires_at": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billing.credit_grant"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
