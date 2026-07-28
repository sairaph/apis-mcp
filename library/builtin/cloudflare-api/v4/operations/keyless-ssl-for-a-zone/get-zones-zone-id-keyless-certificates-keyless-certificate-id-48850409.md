---
title: Get Keyless SSL Configuration
page_id: operation-get-zones-zone-id-keyless-certificates-keyless-certificate-id-9904b4dc
path: operations/keyless-ssl-for-a-zone
description: Get details for one Keyless SSL configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/keyless_certificates/{keyless_certificate_id}
operation_ids:
    - keyless-ssl-for-a-zone-get-keyless-ssl-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Keyless SSL Configuration

`GET /zones/{zone_id}/keyless_certificates/{keyless_certificate_id}`

Operation ID: `keyless-ssl-for-a-zone-get-keyless-ssl-configuration`

Get details for one Keyless SSL configuration.

## Definition

```yaml
{"operationId": "keyless-ssl-for-a-zone-get-keyless-ssl-configuration", "summary": "Get Keyless SSL Configuration", "description": "Get details for one Keyless SSL configuration.", "parameters": [{"name": "keyless_certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Get Keyless SSL Configuration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_keyless_response_single"}}}}, "4XX": {"description": "Get Keyless SSL Configuration response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_keyless_response_single"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Keyless SSL for a Zone"], "x-api-token-group": ["Trust and Safety Write", "Trust and Safety Read", "Zero Trust: PII Read", "Zaraz Edit", "Zaraz Read", "Zaraz Admin", "Access: Apps and Policies Revoke", "Access: Apps and Policies Write", "Access: Apps and Policies Read", "Access: Apps and Policies Revoke", "Access: Mutual TLS Certificates Write", "Access: Organizations, Identity Providers, and Groups Write", "Zone Settings Write", "Zone Settings Read", "Zone Read", "DNS Read", "Workers Scripts Write", "Workers Scripts Read", "Zone Write", "Workers Routes Write", "Workers Routes Read", "Stream Write", "Stream Read", "SSL and Certificates Write", "SSL and Certificates Read", "Logs Write", "Logs Read", "Cache Purge", "Page Rules Write", "Page Rules Read", "Load Balancers Write", "Load Balancers Read", "Firewall Services Write", "Firewall Services Read", "DNS Write", "Apps Write", "Analytics Read", "Access: Apps and Policies Write", "Access: Apps and Policies Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "keyless-certificates", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
