---
title: Delete a service token
page_id: operation-delete-zones-zone-id-access-service-tokens-service-token-id-02c3d900
path: operations/zone-level-access-service-tokens
description: Deletes a service token.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/access/service_tokens/{service_token_id}
operation_ids:
    - zone-level-access-service-tokens-delete-a-service-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a service token

`DELETE /zones/{zone_id}/access/service_tokens/{service_token_id}`

Operation ID: `zone-level-access-service-tokens-delete-a-service-token`

Deletes a service token.

## Definition

```yaml
{"operationId": "zone-level-access-service-tokens-delete-a-service-token", "summary": "Delete a service token", "description": "Deletes a service token.", "parameters": [{"name": "service_token_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Delete a service token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-18"}}}}, "4XX": {"description": "Delete a service token response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access service tokens"], "x-api-token-group": ["Access: Service Tokens Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.service-tokens", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
