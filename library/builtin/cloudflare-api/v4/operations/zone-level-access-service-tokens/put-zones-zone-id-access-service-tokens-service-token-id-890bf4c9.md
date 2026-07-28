---
title: Update a service token
page_id: operation-put-zones-zone-id-access-service-tokens-service-token-id-5507d2c8
path: operations/zone-level-access-service-tokens
description: Updates a configured service token.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/access/service_tokens/{service_token_id}
operation_ids:
    - zone-level-access-service-tokens-update-a-service-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a service token

`PUT /zones/{zone_id}/access/service_tokens/{service_token_id}`

Operation ID: `zone-level-access-service-tokens-update-a-service-token`

Updates a configured service token.

## Definition

```yaml
{"operationId": "zone-level-access-service-tokens-update-a-service-token", "summary": "Update a service token", "description": "Updates a configured service token.", "parameters": [{"name": "service_token_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"client_secret_version": {"$ref": "#/components/schemas/access_client_secret_version"}, "duration": {"$ref": "#/components/schemas/access_duration-2"}, "name": {"$ref": "#/components/schemas/access_name-17"}, "previous_client_secret_expires_at": {"$ref": "#/components/schemas/access_previous_client_secret_expires_at"}}}}}}, "responses": {"200": {"description": "Update a service token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-18"}}}}, "4XX": {"description": "Update a service token response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access service tokens"], "x-api-token-group": ["Access: Service Tokens Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.service-tokens", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
