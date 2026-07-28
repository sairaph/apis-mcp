---
title: Retrieve an active entitlement
page_id: operation-get-v1-entitlements-active-entitlements-id-2a680459
path: operations/untagged
description: <p>Retrieve an active entitlement</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/entitlements/active_entitlements/{id}
operation_ids:
    - GetEntitlementsActiveEntitlementsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve an active entitlement

`GET /v1/entitlements/active_entitlements/{id}`

Operation ID: `GetEntitlementsActiveEntitlementsId`

<p>Retrieve an active entitlement</p>

## Definition

```yaml
{"summary": "Retrieve an active entitlement", "description": "<p>Retrieve an active entitlement</p>", "operationId": "GetEntitlementsActiveEntitlementsId", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "id", "in": "path", "description": "The ID of the entitlement.", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/entitlements.active_entitlement"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
