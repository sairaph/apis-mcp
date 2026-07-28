---
title: Update user label
page_id: operation-put-zones-zone-id-api-gateway-labels-user-name-e13d260d
path: operations/api-shield-labels
description: Update all fields on a label
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/api_gateway/labels/user/{name}
operation_ids:
    - api-shield-put-user-label
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update user label

`PUT /zones/{zone_id}/api_gateway/labels/user/{name}`

Operation ID: `api-shield-put-user-label`

Update all fields on a label

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_label_name_parameter"}]
```

## Definition

```yaml
{"operationId": "api-shield-put-user-label", "summary": "Update user label", "description": "Update all fields on a label", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_put_label_request"}}}}, "responses": {"200": {"description": "Update label response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/api-shield_label"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Update label response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Labels"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.labels.user", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
