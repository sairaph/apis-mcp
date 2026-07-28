---
title: Delete an operation
page_id: operation-delete-zones-zone-id-api-gateway-operations-operation-id-21651802
path: operations/api-shield-endpoint-management
description: Removes a single API operation from API Shield endpoint management. The operation will no longer be tracked or protected by API Shield rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/api_gateway/operations/{operation_id}
operation_ids:
    - api-shield-endpoint-management-delete-an-operation
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an operation

`DELETE /zones/{zone_id}/api_gateway/operations/{operation_id}`

Operation ID: `api-shield-endpoint-management-delete-an-operation`

Removes a single API operation from API Shield endpoint management. The operation will no longer be tracked or protected by API Shield rules.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_operation_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-endpoint-management-delete-an-operation", "summary": "Delete an operation", "description": "Removes a single API operation from API Shield endpoint management. The operation will no longer be tracked or protected by API Shield rules.", "responses": {"200": {"description": "Delete an operation response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common"}}}}, "4XX": {"description": "Delete an operation response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Endpoint Management"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.operations", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
