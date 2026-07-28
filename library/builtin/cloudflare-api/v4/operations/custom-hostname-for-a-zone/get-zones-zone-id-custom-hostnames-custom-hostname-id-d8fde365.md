---
title: Custom Hostname Details
page_id: operation-get-zones-zone-id-custom-hostnames-custom-hostname-id-b5ecfea0
path: operations/custom-hostname-for-a-zone
description: Retrieves detailed information about a specific custom hostname, including SSL certificate status, ownership verification, and origin configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/custom_hostnames/{custom_hostname_id}
operation_ids:
    - custom-hostname-for-a-zone-custom-hostname-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Custom Hostname Details

`GET /zones/{zone_id}/custom_hostnames/{custom_hostname_id}`

Operation ID: `custom-hostname-for-a-zone-custom-hostname-details`

Retrieves detailed information about a specific custom hostname, including SSL certificate status, ownership verification, and origin configuration.

## Definition

```yaml
{"operationId": "custom-hostname-for-a-zone-custom-hostname-details", "summary": "Custom Hostname Details", "description": "Retrieves detailed information about a specific custom hostname, including SSL certificate status, ownership verification, and origin configuration.", "parameters": [{"name": "custom_hostname_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Custom Hostname Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_hostname_response_single"}}}}, "4XX": {"description": "Custom Hostname Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_hostname_response_single"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom Hostname for a Zone"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-hostnames", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
