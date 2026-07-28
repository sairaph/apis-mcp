---
title: Add one operation to a zone
page_id: operation-post-zones-zone-id-api-gateway-operations-item-569ac91b
path: operations/api-shield-endpoint-management
description: Add one operation to a zone. Endpoints can contain path variables. Host, method, endpoint will be normalized to a canoncial form when creating an operation and must be unique on the zone. Inserting an operation that matches an existing one will return the record of the already existing operation and update its last_updated date.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/api_gateway/operations/item
operation_ids:
    - api-shield-endpoint-management-add-operation-to-a-zone
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add one operation to a zone

`POST /zones/{zone_id}/api_gateway/operations/item`

Operation ID: `api-shield-endpoint-management-add-operation-to-a-zone`

Add one operation to a zone. Endpoints can contain path variables. Host, method, endpoint will be normalized to a canoncial form when creating an operation and must be unique on the zone. Inserting an operation that matches an existing one will return the record of the already existing operation and update its last_updated date.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-endpoint-management-add-operation-to-a-zone", "summary": "Add one operation to a zone", "description": "Add one operation to a zone. Endpoints can contain path variables. Host, method, endpoint will be normalized to a canoncial form when creating an operation and must be unique on the zone. Inserting an operation that matches an existing one will return the record of the already existing operation and update its last_updated date.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_basic_operation"}}}}, "responses": {"200": {"description": "Add one operation to a zone response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_single-operation-response"}}}}, "4XX": {"description": "Add one operation to a zone response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Endpoint Management"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.operations", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
