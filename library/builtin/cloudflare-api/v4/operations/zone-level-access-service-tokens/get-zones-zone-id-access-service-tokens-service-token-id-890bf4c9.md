---
title: Get a service token
page_id: operation-get-zones-zone-id-access-service-tokens-service-token-id-b3304637
path: operations/zone-level-access-service-tokens
description: Fetches a single service token.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/access/service_tokens/{service_token_id}
operation_ids:
    - zone-level-access-service-tokens-get-a-service-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a service token

`GET /zones/{zone_id}/access/service_tokens/{service_token_id}`

Operation ID: `zone-level-access-service-tokens-get-a-service-token`

Fetches a single service token.

## Definition

```yaml
{"operationId": "zone-level-access-service-tokens-get-a-service-token", "summary": "Get a service token", "description": "Fetches a single service token.", "parameters": [{"name": "service_token_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Get a service token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-18"}}}}, "4XX": {"description": "Get a service token response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access service tokens"], "x-api-token-group": ["Access: Service Tokens Write", "Access: Service Tokens Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.service-tokens", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
