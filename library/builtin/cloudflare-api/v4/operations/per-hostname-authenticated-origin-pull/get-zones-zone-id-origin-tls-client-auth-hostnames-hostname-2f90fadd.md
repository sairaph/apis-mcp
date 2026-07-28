---
title: Get the Hostname Status for Client Authentication
page_id: operation-get-zones-zone-id-origin-tls-client-auth-hostnames-hostname-8e11e170
path: operations/per-hostname-authenticated-origin-pull
description: Retrieves the client certificate authentication status for a specific hostname, showing whether authenticated origin pulls are enabled.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/origin_tls_client_auth/hostnames/{hostname}
operation_ids:
    - per-hostname-authenticated-origin-pull-get-the-hostname-status-for-client-authentication
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get the Hostname Status for Client Authentication

`GET /zones/{zone_id}/origin_tls_client_auth/hostnames/{hostname}`

Operation ID: `per-hostname-authenticated-origin-pull-get-the-hostname-status-for-client-authentication`

Retrieves the client certificate authentication status for a specific hostname, showing whether authenticated origin pulls are enabled.

## Definition

```yaml
{"operationId": "per-hostname-authenticated-origin-pull-get-the-hostname-status-for-client-authentication", "summary": "Get the Hostname Status for Client Authentication", "description": "Retrieves the client certificate authentication status for a specific hostname, showing whether authenticated origin pulls are enabled.", "parameters": [{"name": "hostname", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_hostname-2"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Get the Hostname Status for Client Authentication response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_hostname_aop_single_response"}}}}, "4XX": {"description": "Get the Hostname Status for Client Authentication response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_hostname_aop_single_response"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Per-hostname Authenticated Origin Pull"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "origin-tls-client-auth.hostnames", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
