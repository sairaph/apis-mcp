---
title: Retrieve information about an operation
page_id: operation-get-zones-zone-id-api-gateway-operations-operation-id-68db4ef2
path: operations/api-shield-endpoint-management
description: Gets detailed information about a specific API operation in API Shield, including its schema validation settings and traffic statistics.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/api_gateway/operations/{operation_id}
operation_ids:
    - api-shield-endpoint-management-retrieve-information-about-an-operation
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve information about an operation

`GET /zones/{zone_id}/api_gateway/operations/{operation_id}`

Operation ID: `api-shield-endpoint-management-retrieve-information-about-an-operation`

Gets detailed information about a specific API operation in API Shield, including its schema validation settings and traffic statistics.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_operation_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-endpoint-management-retrieve-information-about-an-operation", "summary": "Retrieve information about an operation", "description": "Gets detailed information about a specific API operation in API Shield, including its schema validation settings and traffic statistics.", "parameters": [{"$ref": "#/components/parameters/api-shield_operation_feature_parameter"}, {"$ref": "#/components/parameters/api-shield_with_schemas_parameter"}], "responses": {"200": {"description": "Retrieve information about an operation response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_single-operation-response"}}}}, "4XX": {"description": "Retrieve information about an operation response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Endpoint Management"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.operations", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
