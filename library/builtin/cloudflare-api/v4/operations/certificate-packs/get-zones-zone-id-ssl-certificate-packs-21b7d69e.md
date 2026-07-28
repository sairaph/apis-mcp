---
title: List Certificate Packs
page_id: operation-get-zones-zone-id-ssl-certificate-packs-913be470
path: operations/certificate-packs
description: For a given zone, list all active certificate packs.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/ssl/certificate_packs
operation_ids:
    - certificate-packs-list-certificate-packs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Certificate Packs

`GET /zones/{zone_id}/ssl/certificate_packs`

Operation ID: `certificate-packs-list-certificate-packs`

For a given zone, list all active certificate packs.

## Definition

```yaml
{"operationId": "certificate-packs-list-certificate-packs", "summary": "List Certificate Packs", "description": "For a given zone, list all active certificate packs.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Number of certificate packs per page.", "type": "number", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "status", "in": "query", "schema": {"description": "Include Certificate Packs of all statuses, not just active ones.", "type": "string", "example": "all", "enum": ["all"]}}, {"name": "deploy", "in": "query", "schema": {"description": "Specify the deployment environment for the certificate packs.", "type": "string", "enum": ["staging", "production"]}}], "responses": {"200": {"description": "List Certificate Packs response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_pack_response_collection"}}}}, "4XX": {"description": "List Certificate Packs response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_pack_response_collection"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Certificate Packs"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ssl.certificate-packs", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
