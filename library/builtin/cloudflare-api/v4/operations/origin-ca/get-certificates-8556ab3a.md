---
title: List Certificates
page_id: operation-get-certificates-2560d2a0
path: operations/origin-ca
description: List all existing Origin CA certificates for a given zone. You can use an Origin CA Key as your User Service Key or an API token when calling this endpoint ([see above](#requests)).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /certificates
operation_ids:
    - origin-ca-list-certificates
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Certificates

`GET /certificates`

Operation ID: `origin-ca-list-certificates`

List all existing Origin CA certificates for a given zone. You can use an Origin CA Key as your User Service Key or an API token when calling this endpoint ([see above](#requests)).

## Definition

```yaml
{"operationId": "origin-ca-list-certificates", "summary": "List Certificates", "description": "List all existing Origin CA certificates for a given zone. You can use an Origin CA Key as your User Service Key or an API token when calling this endpoint ([see above](#requests)).", "parameters": [{"name": "zone_id", "in": "query", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Number of records per page.", "type": "number", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "limit", "in": "query", "schema": {"description": "Limit to the number of records returned.", "type": "integer", "example": 10}}, {"name": "offset", "in": "query", "schema": {"description": "Offset the results.", "type": "integer", "example": 10}}], "responses": {"200": {"description": "List Certificates response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_collection-2"}}}}, "4XX": {"description": "List Certificates response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_collection-2"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"user_service_key": []}], "tags": ["Origin CA"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "origin-ca-certificates", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
