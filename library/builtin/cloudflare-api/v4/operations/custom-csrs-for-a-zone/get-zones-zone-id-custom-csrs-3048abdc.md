---
title: List Custom CSRs
page_id: operation-get-zones-zone-id-custom-csrs-e5610864
path: operations/custom-csrs-for-a-zone
description: List all custom Certificate Signing Requests (CSRs) for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/custom_csrs
operation_ids:
    - custom-csrs-for-a-zone-list-custom-csrs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Custom CSRs

`GET /zones/{zone_id}/custom_csrs`

Operation ID: `custom-csrs-for-a-zone-list-custom-csrs`

List all custom Certificate Signing Requests (CSRs) for a zone.

## Definition

```yaml
{"operationId": "custom-csrs-for-a-zone-list-custom-csrs", "summary": "List Custom CSRs", "description": "List all custom Certificate Signing Requests (CSRs) for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Number of custom CSRs per page.", "type": "number", "default": 20, "maximum": 50, "minimum": 5}}], "responses": {"200": {"description": "List Custom CSRs response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_response_collection"}}}}, "4XX": {"description": "List Custom CSRs response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_response_collection"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom CSRs for a Zone"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
