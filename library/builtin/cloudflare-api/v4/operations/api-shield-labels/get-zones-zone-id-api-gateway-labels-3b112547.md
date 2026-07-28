---
title: Retrieve all labels
page_id: operation-get-zones-zone-id-api-gateway-labels-efbb80bf
path: operations/api-shield-labels
description: Retrieve all labels
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/api_gateway/labels
operation_ids:
    - api-shield-labels-get-labels
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve all labels

`GET /zones/{zone_id}/api_gateway/labels`

Operation ID: `api-shield-labels-get-labels`

Retrieve all labels

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-labels-get-labels", "summary": "Retrieve all labels", "description": "Retrieve all labels", "parameters": [{"$ref": "#/components/parameters/api-shield_page"}, {"$ref": "#/components/parameters/api-shield_per_page"}, {"$ref": "#/components/parameters/api-shield_order_parameter-2"}, {"$ref": "#/components/parameters/api-shield_direction_parameter"}, {"$ref": "#/components/parameters/api-shield_source_parameter"}, {"$ref": "#/components/parameters/api-shield_filter_parameter"}, {"$ref": "#/components/parameters/api-shield_with_mapped_resource_counts_parameter"}], "responses": {"200": {"description": "Retrieve all labels response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_full_label"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Retrieve all labels response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Labels"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.labels", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
