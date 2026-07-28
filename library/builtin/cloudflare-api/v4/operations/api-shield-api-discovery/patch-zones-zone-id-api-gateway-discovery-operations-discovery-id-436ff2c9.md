---
title: Patch discovered operation
page_id: operation-patch-zones-zone-id-api-gateway-discovery-operations-discovery-id-13d9a720
path: operations/api-shield-api-discovery
description: Update the `state` on a discovered operation
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/api_gateway/discovery/operations/{discovery_id}
operation_ids:
    - api-shield-api-patch-discovered-operation
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch discovered operation

`PATCH /zones/{zone_id}/api_gateway/discovery/operations/{discovery_id}`

Operation ID: `api-shield-api-patch-discovered-operation`

Update the `state` on a discovered operation

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_discovery_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-api-patch-discovered-operation", "summary": "Patch discovered operation", "description": "Update the `state` on a discovered operation", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"state": {"allOf": [{"$ref": "#/components/schemas/api-shield_api_discovery_state_patch"}]}}}}}}, "responses": {"200": {"description": "Patch discovered operation response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_patch_discovery_response"}}}}, "4XX": {"description": "Patch discovered operation response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield API Discovery"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"]}
```
