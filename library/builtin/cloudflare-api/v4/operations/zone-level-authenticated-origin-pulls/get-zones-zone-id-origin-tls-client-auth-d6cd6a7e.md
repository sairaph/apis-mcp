---
title: List Certificates
page_id: operation-get-zones-zone-id-origin-tls-client-auth-c7c225f3
path: operations/zone-level-authenticated-origin-pulls
description: Lists all client certificates configured for zone-level authenticated origin pulls.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/origin_tls_client_auth
operation_ids:
    - zone-level-authenticated-origin-pulls-list-certificates
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Certificates

`GET /zones/{zone_id}/origin_tls_client_auth`

Operation ID: `zone-level-authenticated-origin-pulls-list-certificates`

Lists all client certificates configured for zone-level authenticated origin pulls.

## Definition

```yaml
{"operationId": "zone-level-authenticated-origin-pulls-list-certificates", "summary": "List Certificates", "description": "Lists all client certificates configured for zone-level authenticated origin pulls.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "List Certificates response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_collection-3"}}}}, "4XX": {"description": "List Certificates response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_collection-3"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone-Level Authenticated Origin Pulls"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "origin-tls-client-auth.zone-certificates", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
