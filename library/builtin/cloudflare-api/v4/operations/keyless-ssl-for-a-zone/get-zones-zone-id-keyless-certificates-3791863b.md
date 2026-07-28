---
title: List Keyless SSL Configurations
page_id: operation-get-zones-zone-id-keyless-certificates-9ac77b24
path: operations/keyless-ssl-for-a-zone
description: List all Keyless SSL configurations for a given zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/keyless_certificates
operation_ids:
    - keyless-ssl-for-a-zone-list-keyless-ssl-configurations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Keyless SSL Configurations

`GET /zones/{zone_id}/keyless_certificates`

Operation ID: `keyless-ssl-for-a-zone-list-keyless-ssl-configurations`

List all Keyless SSL configurations for a given zone.

## Definition

```yaml
{"operationId": "keyless-ssl-for-a-zone-list-keyless-ssl-configurations", "summary": "List Keyless SSL Configurations", "description": "List all Keyless SSL configurations for a given zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "List Keyless SSL Configurations response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_keyless_response_collection"}}}}, "4XX": {"description": "List Keyless SSL Configurations response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_keyless_response_collection"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Keyless SSL for a Zone"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "keyless-certificates", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
