---
title: Retrieve a discovered operation
page_id: operation-get-zones-zone-id-api-gateway-discovery-operations-discovery-id-279b3dc6
path: operations/api-shield-api-discovery
description: Retrieve a single discovered operation by ID
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/api_gateway/discovery/operations/{discovery_id}
operation_ids:
    - api-shield-api-discovery-retrieve-discovered-operation-by-id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve a discovered operation

`GET /zones/{zone_id}/api_gateway/discovery/operations/{discovery_id}`

Operation ID: `api-shield-api-discovery-retrieve-discovered-operation-by-id`

Retrieve a single discovered operation by ID

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_discovery_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-api-discovery-retrieve-discovered-operation-by-id", "summary": "Retrieve a discovered operation", "description": "Retrieve a single discovered operation by ID", "responses": {"200": {"description": "Retrieve discovered operation response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_discovery_operation"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Retrieve discovered operation response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield API Discovery"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"]}
```
