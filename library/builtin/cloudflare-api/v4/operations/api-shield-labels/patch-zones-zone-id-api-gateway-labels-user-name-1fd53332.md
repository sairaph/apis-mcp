---
title: Patch user label
page_id: operation-patch-zones-zone-id-api-gateway-labels-user-name-5e755f85
path: operations/api-shield-labels
description: Update certain fields on a label
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/api_gateway/labels/user/{name}
operation_ids:
    - api-shield-patch-user-label
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch user label

`PATCH /zones/{zone_id}/api_gateway/labels/user/{name}`

Operation ID: `api-shield-patch-user-label`

Update certain fields on a label

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_label_name_parameter"}]
```

## Definition

```yaml
{"operationId": "api-shield-patch-user-label", "summary": "Patch user label", "description": "Update certain fields on a label", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_patch_label_request"}}}}, "responses": {"200": {"description": "Patch label response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/api-shield_label"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Patch label response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Labels"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.labels.user", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
