---
title: Delete user label
page_id: operation-delete-zones-zone-id-api-gateway-labels-user-name-93f89f99
path: operations/api-shield-labels
description: Delete user label
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/api_gateway/labels/user/{name}
operation_ids:
    - api-shield-delete-user-label
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete user label

`DELETE /zones/{zone_id}/api_gateway/labels/user/{name}`

Operation ID: `api-shield-delete-user-label`

Delete user label

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_label_name_parameter"}]
```

## Definition

```yaml
{"operationId": "api-shield-delete-user-label", "summary": "Delete user label", "description": "Delete user label", "responses": {"200": {"description": "Delete user label response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/api-shield_label"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Delete label response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Labels"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.labels.user", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
