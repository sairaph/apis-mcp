---
title: Retrieve information about all operations on a zone
page_id: operation-get-zones-zone-id-api-gateway-operations-2b0e0d0e
path: operations/api-shield-endpoint-management
description: Lists all API operations tracked by API Shield for a zone with pagination. Returns operation details including method, path, and feature configurations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/api_gateway/operations
operation_ids:
    - api-shield-endpoint-management-retrieve-information-about-all-operations-on-a-zone
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve information about all operations on a zone

`GET /zones/{zone_id}/api_gateway/operations`

Operation ID: `api-shield-endpoint-management-retrieve-information-about-all-operations-on-a-zone`

Lists all API operations tracked by API Shield for a zone with pagination. Returns operation details including method, path, and feature configurations.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-endpoint-management-retrieve-information-about-all-operations-on-a-zone", "summary": "Retrieve information about all operations on a zone", "description": "Lists all API operations tracked by API Shield for a zone with pagination. Returns operation details including method, path, and feature configurations.", "parameters": [{"$ref": "#/components/parameters/api-shield_page"}, {"$ref": "#/components/parameters/api-shield_per_page"}, {"name": "order", "in": "query", "schema": {"description": "Field to order by. When requesting a feature, the feature keys are available for ordering as well, e.g., `thresholds.suggested_threshold`.", "type": "string", "example": "method", "enum": ["method", "host", "endpoint", "thresholds.$key"]}}, {"$ref": "#/components/parameters/api-shield_direction_parameter"}, {"$ref": "#/components/parameters/api-shield_host_parameter"}, {"$ref": "#/components/parameters/api-shield_method_parameter"}, {"$ref": "#/components/parameters/api-shield_endpoint_parameter"}, {"$ref": "#/components/parameters/api-shield_operation_feature_parameter"}], "responses": {"200": {"description": "Retrieve information about all operations on a zone response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_multiple-operation-response-paginated"}}}}, "4XX": {"description": "Retrieve information about all operations on a zone response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Endpoint Management"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.operations", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
