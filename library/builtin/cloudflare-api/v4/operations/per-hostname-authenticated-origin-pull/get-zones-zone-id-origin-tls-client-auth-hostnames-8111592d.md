---
title: List Hostname Associations
page_id: operation-get-zones-zone-id-origin-tls-client-auth-hostnames-43188a1f
path: operations/per-hostname-authenticated-origin-pull
description: List certificate ID - hostname associations for the given zone. Shows which hostnames are associated to which certificates for authenticated origin pulls.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/origin_tls_client_auth/hostnames
operation_ids:
    - per-hostname-authenticated-origin-pull-list-hostname-associations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Hostname Associations

`GET /zones/{zone_id}/origin_tls_client_auth/hostnames`

Operation ID: `per-hostname-authenticated-origin-pull-list-hostname-associations`

List certificate ID - hostname associations for the given zone. Shows which hostnames are associated to which certificates for authenticated origin pulls.

## Definition

```yaml
{"operationId": "per-hostname-authenticated-origin-pull-list-hostname-associations", "summary": "List Hostname Associations", "description": "List certificate ID - hostname associations for the given zone. Shows which hostnames are associated to which certificates for authenticated origin pulls.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Number of associations per page.", "type": "number", "default": 50, "maximum": 1000, "minimum": 1}}, {"name": "status", "in": "query", "schema": {"description": "Filter associations by status. Use comma-separated values to filter by multiple statuses, or \"all\" to include every status. Defaults to \"active\" if not provided.", "type": "string", "example": "active", "enum": ["active", "pending_deployment", "pending_deletion", "deleted", "deployment_timed_out", "deletion_timed_out", "all"]}}], "responses": {"200": {"description": "List Hostname Associations response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_hostname_assoc_response_collection"}}}}, "4XX": {"description": "List Hostname Associations response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_hostname_assoc_response_collection"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Per-hostname Authenticated Origin Pull"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "origin-tls-client-auth.hostname-associations", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
