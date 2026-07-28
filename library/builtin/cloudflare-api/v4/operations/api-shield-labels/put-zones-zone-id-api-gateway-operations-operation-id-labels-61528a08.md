---
title: Replace label(s) on an operation in endpoint management
page_id: operation-put-zones-zone-id-api-gateway-operations-operation-id-labels-ba1b0b35
path: operations/api-shield-labels
description: Replace label(s) on an operation in endpoint management
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/api_gateway/operations/{operation_id}/labels
operation_ids:
    - api-shield-operations-put-labels-to-operation
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Replace label(s) on an operation in endpoint management

`PUT /zones/{zone_id}/api_gateway/operations/{operation_id}/labels`

Operation ID: `api-shield-operations-put-labels-to-operation`

Replace label(s) on an operation in endpoint management

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_operation_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-operations-put-labels-to-operation", "summary": "Replace label(s) on an operation in endpoint management", "description": "Replace label(s) on an operation in endpoint management", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_put_labels_on_operation_request"}}}}, "responses": {"200": {"description": "Replace label(s) on an operation in endpoint management response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/api-shield_operation_with_labels_only"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Replace label(s) on an operation in endpoint management response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Labels"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.operations.labels", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
