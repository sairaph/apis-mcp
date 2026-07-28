---
title: Bulk remove label(s) on operation(s) in endpoint management
page_id: operation-delete-zones-zone-id-api-gateway-operations-labels-83d7d94b
path: operations/api-shield-labels
description: Bulk remove label(s) on operation(s) in endpoint management
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/api_gateway/operations/labels
operation_ids:
    - api-shield-operations-bulk-delete-labels-to-operations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Bulk remove label(s) on operation(s) in endpoint management

`DELETE /zones/{zone_id}/api_gateway/operations/labels`

Operation ID: `api-shield-operations-bulk-delete-labels-to-operations`

Bulk remove label(s) on operation(s) in endpoint management

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-operations-bulk-delete-labels-to-operations", "summary": "Bulk remove label(s) on operation(s) in endpoint management", "description": "Bulk remove label(s) on operation(s) in endpoint management", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_bulk_delete_labels_on_operation_request"}}}}, "responses": {"200": {"description": "Bulk remove label(s) on operation(s) in endpoint management response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_operation_with_labels_only"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Bulk remove label(s) on operation(s) in endpoint management response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Labels"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.operations.labels", "x-fern-sdk-method-name": "bulk-delete", "x-forge-hidden": true}
```
