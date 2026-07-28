---
title: Attach label(s) on an operation in endpoint management
page_id: operation-post-zones-zone-id-api-gateway-operations-operation-id-labels-d4dcf6a9
path: operations/api-shield-labels
description: Attach label(s) on an operation in endpoint management
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/api_gateway/operations/{operation_id}/labels
operation_ids:
    - api-shield-operations-post-labels-to-operation
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Attach label(s) on an operation in endpoint management

`POST /zones/{zone_id}/api_gateway/operations/{operation_id}/labels`

Operation ID: `api-shield-operations-post-labels-to-operation`

Attach label(s) on an operation in endpoint management

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_operation_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-operations-post-labels-to-operation", "summary": "Attach label(s) on an operation in endpoint management", "description": "Attach label(s) on an operation in endpoint management", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_post_labels_on_operation_request"}}}}, "responses": {"200": {"description": "Attach label(s) on an operation in endpoint management response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/api-shield_operation_with_labels_only"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Attach label(s) on an operation in endpoint management response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Labels"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.operations.labels", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
