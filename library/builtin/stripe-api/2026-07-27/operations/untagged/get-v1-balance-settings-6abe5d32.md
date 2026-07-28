---
title: Retrieve balance settings
page_id: operation-get-v1-balance-settings-e8cb4538
path: operations/untagged
description: |-
    <p>Retrieves balance settings for a given connected account.
     Related guide: <a href="/connect/authentication">Making API calls for connected accounts</a></p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/balance_settings
operation_ids:
    - GetBalanceSettings
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve balance settings

`GET /v1/balance_settings`

Operation ID: `GetBalanceSettings`

<p>Retrieves balance settings for a given connected account.
 Related guide: <a href="/connect/authentication">Making API calls for connected accounts</a></p>

## Definition

```yaml
{"summary": "Retrieve balance settings", "description": "<p>Retrieves balance settings for a given connected account.\n Related guide: <a href=\"/connect/authentication\">Making API calls for connected accounts</a></p>", "operationId": "GetBalanceSettings", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/balance_settings"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
