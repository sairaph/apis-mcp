---
title: Get Gateway URL
page_id: operation-get-accounts-account-id-ai-gateway-gateways-gateway-id-url-provider-a18ae4c1
path: operations/ai-gateway-gateways
description: Retrieves the endpoint URL for an AI Gateway.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/url/{provider}
operation_ids:
    - aig-config-get-gateway-url
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Gateway URL

`GET /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/url/{provider}`

Operation ID: `aig-config-get-gateway-url`

Retrieves the endpoint URL for an AI Gateway.

## Definition

```yaml
{"operationId": "aig-config-get-gateway-url", "summary": "Get Gateway URL", "description": "Retrieves the endpoint URL for an AI Gateway.", "parameters": [{"name": "gateway_id", "in": "path", "required": true, "schema": {"description": "gateway id", "type": "string", "example": "my-gateway", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "0d37909e38d3e99c29fa2cd343ac421a"}}, {"name": "provider", "in": "path", "required": true, "schema": {"type": "string", "example": "workers-ai"}}], "responses": {"200": {"description": "Returns the log details", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "string"}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway Gateways"], "x-api-token-group": ["AI Gateway Write", "AI Gateway Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai-gateway.urls", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
