---
title: Create user labels
page_id: operation-post-zones-zone-id-api-gateway-labels-user-f2a88b2f
path: operations/api-shield-labels
description: Create user labels
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/api_gateway/labels/user
operation_ids:
    - api-shield-labels-create-user-labels
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create user labels

`POST /zones/{zone_id}/api_gateway/labels/user`

Operation ID: `api-shield-labels-create-user-labels`

Create user labels

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-labels-create-user-labels", "summary": "Create user labels", "description": "Create user labels", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_label_request"}}}}}, "responses": {"200": {"description": "Create user labels response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_label"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Create user labels response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Labels"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.labels.user", "x-fern-sdk-method-name": "bulk-create", "x-forge-hidden": true}
```
