---
title: Retrieve managed label
page_id: operation-get-zones-zone-id-api-gateway-labels-managed-name-a3b1be78
path: operations/api-shield-labels
description: Retrieve managed label
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/api_gateway/labels/managed/{name}
operation_ids:
    - api-shield-labels-get-managed-label
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve managed label

`GET /zones/{zone_id}/api_gateway/labels/managed/{name}`

Operation ID: `api-shield-labels-get-managed-label`

Retrieve managed label

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_label_name_parameter"}]
```

## Definition

```yaml
{"operationId": "api-shield-labels-get-managed-label", "summary": "Retrieve managed label", "description": "Retrieve managed label", "parameters": [{"$ref": "#/components/parameters/api-shield_with_mapped_resource_counts_parameter"}], "responses": {"200": {"description": "Retrieve managed label response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/api-shield_full_managed_label"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Retrieve managed label response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Labels"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.labels.managed", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
