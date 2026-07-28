---
title: Retrieve discovered operations on a zone rendered as OpenAPI schemas
page_id: operation-get-zones-zone-id-api-gateway-discovery-cdf3b6bd
path: operations/api-shield-api-discovery
description: Retrieve the most up to date view of discovered operations, rendered as OpenAPI schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/api_gateway/discovery
operation_ids:
    - api-shield-api-discovery-retrieve-discovered-operations-on-a-zone-as-openapi
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve discovered operations on a zone rendered as OpenAPI schemas

`GET /zones/{zone_id}/api_gateway/discovery`

Operation ID: `api-shield-api-discovery-retrieve-discovered-operations-on-a-zone-as-openapi`

Retrieve the most up to date view of discovered operations, rendered as OpenAPI schemas

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-api-discovery-retrieve-discovered-operations-on-a-zone-as-openapi", "summary": "Retrieve discovered operations on a zone rendered as OpenAPI schemas", "description": "Retrieve the most up to date view of discovered operations, rendered as OpenAPI schemas", "responses": {"200": {"description": "Retrieve discovered operations on a zone, rendered as OpenAPI schemas response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_schema_response_discovery"}}}}, "4XX": {"description": "Retrieve discovered operations on a zone, rendered as OpenAPI schemas response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_schema_response_discovery"}, {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield API Discovery"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.discovery", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
