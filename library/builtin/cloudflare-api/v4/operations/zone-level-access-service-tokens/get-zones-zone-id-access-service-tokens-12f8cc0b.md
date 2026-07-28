---
title: List service tokens
page_id: operation-get-zones-zone-id-access-service-tokens-2da400a8
path: operations/zone-level-access-service-tokens
description: Lists all service tokens.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/access/service_tokens
operation_ids:
    - zone-level-access-service-tokens-list-service-tokens
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List service tokens

`GET /zones/{zone_id}/access/service_tokens`

Operation ID: `zone-level-access-service-tokens-list-service-tokens`

Lists all service tokens.

## Definition

```yaml
{"operationId": "zone-level-access-service-tokens-list-service-tokens", "summary": "List service tokens", "description": "Lists all service tokens.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "List service tokens response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-19"}}}}, "4XX": {"description": "List service tokens response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access service tokens"], "x-api-token-group": ["Access: Service Tokens Write", "Access: Service Tokens Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.service-tokens", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
