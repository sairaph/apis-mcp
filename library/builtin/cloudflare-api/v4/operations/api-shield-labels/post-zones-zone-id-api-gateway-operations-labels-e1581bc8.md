---
title: Bulk attach label(s) on operation(s) in endpoint management
page_id: operation-post-zones-zone-id-api-gateway-operations-labels-1ac9ca8c
path: operations/api-shield-labels
description: Bulk attach label(s) on operation(s) in endpoint management
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/api_gateway/operations/labels
operation_ids:
    - api-shield-operations-bulk-post-labels-to-operations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Bulk attach label(s) on operation(s) in endpoint management

`POST /zones/{zone_id}/api_gateway/operations/labels`

Operation ID: `api-shield-operations-bulk-post-labels-to-operations`

Bulk attach label(s) on operation(s) in endpoint management

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-operations-bulk-post-labels-to-operations", "summary": "Bulk attach label(s) on operation(s) in endpoint management", "description": "Bulk attach label(s) on operation(s) in endpoint management", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_bulk_post_labels_on_operation_request"}}}}, "responses": {"200": {"description": "Bulk attach label(s) on operation(s) in endpoint management response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_operation_with_labels_only"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Bulk attach label(s) on operation(s) in endpoint management response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Labels"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.operations.labels", "x-fern-sdk-method-name": "bulk-create", "x-forge-hidden": true}
```
