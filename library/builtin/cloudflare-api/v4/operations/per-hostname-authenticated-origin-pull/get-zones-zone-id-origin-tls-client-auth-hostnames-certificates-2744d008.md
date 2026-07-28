---
title: List Certificates
page_id: operation-get-zones-zone-id-origin-tls-client-auth-hostnames-certificates-1fe3c1d4
path: operations/per-hostname-authenticated-origin-pull
description: Lists all client certificates configured for per-hostname authenticated origin pulls on the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/origin_tls_client_auth/hostnames/certificates
operation_ids:
    - per-hostname-authenticated-origin-pull-list-certificates
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Certificates

`GET /zones/{zone_id}/origin_tls_client_auth/hostnames/certificates`

Operation ID: `per-hostname-authenticated-origin-pull-list-certificates`

Lists all client certificates configured for per-hostname authenticated origin pulls on the zone.

## Definition

```yaml
{"operationId": "per-hostname-authenticated-origin-pull-list-certificates", "summary": "List Certificates", "description": "Lists all client certificates configured for per-hostname authenticated origin pulls on the zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "List Certificates response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_collection-4"}}}}, "4XX": {"description": "List Certificates response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_collection-4"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Per-hostname Authenticated Origin Pull"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "origin-tls-client-auth.hostname-certificates", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
