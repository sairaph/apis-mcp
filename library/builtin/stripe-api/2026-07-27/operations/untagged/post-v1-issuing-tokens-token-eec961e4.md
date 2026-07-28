---
title: Update a token status
page_id: operation-post-v1-issuing-tokens-token-61405aa2
path: operations/untagged
description: <p>Attempts to update the specified Issuing <code>Token</code> object to the status specified.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/issuing/tokens/{token}
operation_ids:
    - PostIssuingTokensToken
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update a token status

`POST /v1/issuing/tokens/{token}`

Operation ID: `PostIssuingTokensToken`

<p>Attempts to update the specified Issuing <code>Token</code> object to the status specified.</p>

## Definition

```yaml
{"summary": "Update a token status", "description": "<p>Attempts to update the specified Issuing <code>Token</code> object to the status specified.</p>", "operationId": "PostIssuingTokensToken", "parameters": [{"name": "token", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["status"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "status": {"type": "string", "description": "Specifies which status the token should be updated to.", "enum": ["active", "deleted", "suspended"]}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.token"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
