---
title: Patch discovered operations
page_id: operation-patch-zones-zone-id-api-gateway-discovery-operations-48940df9
path: operations/api-shield-api-discovery
description: Update the `state` on one or more discovered operations
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/api_gateway/discovery/operations
operation_ids:
    - api-shield-api-patch-discovered-operations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch discovered operations

`PATCH /zones/{zone_id}/api_gateway/discovery/operations`

Operation ID: `api-shield-api-patch-discovered-operations`

Update the `state` on one or more discovered operations

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-api-patch-discovered-operations", "summary": "Patch discovered operations", "description": "Update the `state` on one or more discovered operations", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api_discovery_patch_multiple_request"}}}}, "responses": {"200": {"description": "Patch discovered operations response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_patch_discoveries_response"}}}}, "4XX": {"description": "Patch discovered operations response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield API Discovery"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.discovery.operations", "x-fern-sdk-method-name": "bulk-edit", "x-forge-hidden": true}
```
