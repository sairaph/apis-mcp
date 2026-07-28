---
title: Get Zone Bot Management Config
page_id: operation-get-zones-zone-id-bot-management-ea2b9143
path: operations/bot-settings
description: Retrieve a zone's Bot Management Config
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/bot_management
operation_ids:
    - bot-management-for-a-zone-get-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Zone Bot Management Config

`GET /zones/{zone_id}/bot_management`

Operation ID: `bot-management-for-a-zone-get-config`

Retrieve a zone's Bot Management Config

## Definition

```yaml
{"operationId": "bot-management-for-a-zone-get-config", "summary": "Get Zone Bot Management Config", "description": "Retrieve a zone's Bot Management Config", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bot-management_identifier"}}], "responses": {"200": {"description": "Bot Management config response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bot-management_bot_management_response_body"}}}}, "4XX": {"description": "Bot Management config response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bot-management_bot_management_response_body"}, {"$ref": "#/components/schemas/bot-management_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Bot Settings"], "x-api-token-group": ["Bot Management Write", "Bot Management Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
