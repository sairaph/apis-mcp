---
title: Add operations to a zone
page_id: operation-post-zones-zone-id-api-gateway-operations-86cb0214
path: operations/api-shield-endpoint-management
description: Add one or more operations to a zone. Endpoints can contain path variables. Host, method, endpoint will be normalized to a canoncial form when creating an operation and must be unique on the zone. Inserting an operation that matches an existing one will return the record of the already existing operation and update its last_updated date.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/api_gateway/operations
operation_ids:
    - api-shield-endpoint-management-add-operations-to-a-zone
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add operations to a zone

`POST /zones/{zone_id}/api_gateway/operations`

Operation ID: `api-shield-endpoint-management-add-operations-to-a-zone`

Add one or more operations to a zone. Endpoints can contain path variables. Host, method, endpoint will be normalized to a canoncial form when creating an operation and must be unique on the zone. Inserting an operation that matches an existing one will return the record of the already existing operation and update its last_updated date.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-endpoint-management-add-operations-to-a-zone", "summary": "Add operations to a zone", "description": "Add one or more operations to a zone. Endpoints can contain path variables. Host, method, endpoint will be normalized to a canoncial form when creating an operation and must be unique on the zone. Inserting an operation that matches an existing one will return the record of the already existing operation and update its last_updated date.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_basic_operation"}}}}}, "responses": {"200": {"description": "Add operations to a zone response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_multiple-operation-response"}}}}, "4XX": {"description": "Add operations to a zone response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Endpoint Management"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.operations", "x-fern-sdk-method-name": "bulk-create", "x-forge-hidden": true}
```
