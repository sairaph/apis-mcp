---
title: List Custom Origin Trust Store Details
page_id: operation-get-zones-zone-id-acm-custom-trust-store-662a453d
path: operations/custom-origin-trust-store
description: Get Custom Origin Trust Store for a Zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/acm/custom_trust_store
operation_ids:
    - custom-origin-trust-store-list-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Custom Origin Trust Store Details

`GET /zones/{zone_id}/acm/custom_trust_store`

Operation ID: `custom-origin-trust-store-list-details`

Get Custom Origin Trust Store for a Zone.

## Definition

```yaml
{"operationId": "custom-origin-trust-store-list-details", "summary": "List Custom Origin Trust Store Details", "description": "Get Custom Origin Trust Store for a Zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Number of records per page.", "type": "number", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "limit", "in": "query", "schema": {"description": "Limit to the number of records returned.", "type": "integer", "example": 10}}, {"name": "offset", "in": "query", "schema": {"description": "Offset the results.", "type": "integer", "example": 10}}], "responses": {"200": {"description": "Custom Origin Trust Store Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_trust_store_response_collection"}}}}, "4XX": {"description": "Custom Origin Trust Store response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_trust_store_response_collection"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom Origin Trust Store"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "acm.custom-trust-store", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
