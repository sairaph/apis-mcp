---
title: Retrieve discovered operations on a zone
page_id: operation-get-zones-zone-id-api-gateway-discovery-operations-e77c5860
path: operations/api-shield-api-discovery
description: Retrieve the most up to date view of discovered operations
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/api_gateway/discovery/operations
operation_ids:
    - api-shield-api-discovery-retrieve-discovered-operations-on-a-zone
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve discovered operations on a zone

`GET /zones/{zone_id}/api_gateway/discovery/operations`

Operation ID: `api-shield-api-discovery-retrieve-discovered-operations-on-a-zone`

Retrieve the most up to date view of discovered operations

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-api-discovery-retrieve-discovered-operations-on-a-zone", "summary": "Retrieve discovered operations on a zone", "description": "Retrieve the most up to date view of discovered operations", "parameters": [{"$ref": "#/components/parameters/api-shield_page"}, {"$ref": "#/components/parameters/api-shield_per_page"}, {"$ref": "#/components/parameters/api-shield_host_parameter"}, {"$ref": "#/components/parameters/api-shield_method_parameter"}, {"$ref": "#/components/parameters/api-shield_endpoint_parameter"}, {"$ref": "#/components/parameters/api-shield_direction_parameter"}, {"$ref": "#/components/parameters/api-shield_order_parameter"}, {"$ref": "#/components/parameters/api-shield_diff_parameter"}, {"$ref": "#/components/parameters/api-shield_api_discovery_origin_parameter"}, {"$ref": "#/components/parameters/api-shield_api_discovery_state_parameter"}], "responses": {"200": {"description": "Retrieve discovered operations on a zone response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_discovery_operation"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Retrieve discovered operations on a zone response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield API Discovery"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.discovery.operations", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
