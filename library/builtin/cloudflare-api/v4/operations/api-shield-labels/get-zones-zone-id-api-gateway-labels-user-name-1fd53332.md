---
title: Retrieve user label
page_id: operation-get-zones-zone-id-api-gateway-labels-user-name-2b2ef210
path: operations/api-shield-labels
description: Retrieve user label
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/api_gateway/labels/user/{name}
operation_ids:
    - api-shield-labels-get-user-label
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve user label

`GET /zones/{zone_id}/api_gateway/labels/user/{name}`

Operation ID: `api-shield-labels-get-user-label`

Retrieve user label

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_label_name_parameter"}]
```

## Definition

```yaml
{"operationId": "api-shield-labels-get-user-label", "summary": "Retrieve user label", "description": "Retrieve user label", "parameters": [{"$ref": "#/components/parameters/api-shield_with_mapped_resource_counts_parameter"}], "responses": {"200": {"description": "Retrieve user label response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/api-shield_full_label"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Retrieve user label response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Labels"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.labels.user", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
