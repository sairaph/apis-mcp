---
title: Delete multiple operations
page_id: operation-delete-zones-zone-id-api-gateway-operations-5b1a9d06
path: operations/api-shield-endpoint-management
description: Bulk removes multiple API operations from API Shield endpoint management in a single request. Efficient for cleaning up unused endpoints.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/api_gateway/operations
operation_ids:
    - api-shield-endpoint-management-delete-multiple-operations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete multiple operations

`DELETE /zones/{zone_id}/api_gateway/operations`

Operation ID: `api-shield-endpoint-management-delete-multiple-operations`

Bulk removes multiple API operations from API Shield endpoint management in a single request. Efficient for cleaning up unused endpoints.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-endpoint-management-delete-multiple-operations", "summary": "Delete multiple operations", "description": "Bulk removes multiple API operations from API Shield endpoint management in a single request. Efficient for cleaning up unused endpoints.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_object-with-operation-id"}, "example": [{"operation_id": "b17c8043-99a0-4202-b7d9-8f7cdbee02cd"}, {"operation_id": "3818d821-5901-4147-a474-f5f5aec1d54e"}], "uniqueItems": true}}}}, "responses": {"200": {"description": "Delete multiple operations response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common"}}}}, "4XX": {"description": "Delete multiple operations response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Endpoint Management"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.operations", "x-fern-sdk-method-name": "bulk-delete", "x-forge-hidden": true}
```
