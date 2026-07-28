---
title: Retrieve operations and features as OpenAPI schemas
page_id: operation-get-zones-zone-id-api-gateway-schemas-288c93c1
path: operations/api-shield-endpoint-management
description: Retrieves API operations and their features exported as OpenAPI schemas.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/api_gateway/schemas
operation_ids:
    - api-shield-endpoint-management-retrieve-operations-and-features-as-open-api-schemas
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve operations and features as OpenAPI schemas

`GET /zones/{zone_id}/api_gateway/schemas`

Operation ID: `api-shield-endpoint-management-retrieve-operations-and-features-as-open-api-schemas`

Retrieves API operations and their features exported as OpenAPI schemas.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-endpoint-management-retrieve-operations-and-features-as-open-api-schemas", "summary": "Retrieve operations and features as OpenAPI schemas", "description": "Retrieves API operations and their features exported as OpenAPI schemas.", "parameters": [{"name": "host", "in": "query", "schema": {"description": "Receive schema only for the given host(s).", "type": "array", "items": {"example": "www.example.com", "type": "string"}, "uniqueItems": true}}, {"$ref": "#/components/parameters/api-shield_operation_feature_parameter"}, {"name": "include_schema_kind", "in": "query", "description": "Schema kinds to include in exported OpenAPI schemas.", "schema": {"type": "array", "items": {"enum": ["learned"], "type": "string"}, "uniqueItems": true}, "explode": true, "style": "form"}], "responses": {"200": {"description": "Retrieve operations and features as OpenAPI schemas response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_schema-response-with-thresholds"}}}}, "4XX": {"description": "Retrieve operations and features as OpenAPI schemas response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Endpoint Management"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.schemas", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
