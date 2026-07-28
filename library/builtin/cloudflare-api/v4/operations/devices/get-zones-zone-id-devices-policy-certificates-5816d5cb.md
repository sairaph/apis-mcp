---
title: Get device certificate provisioning status
page_id: operation-get-zones-zone-id-devices-policy-certificates-a9d1b2c7
path: operations/devices
description: Fetches device certificate provisioning.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/devices/policy/certificates
operation_ids:
    - devices-get-policy-certificates
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get device certificate provisioning status

`GET /zones/{zone_id}/devices/policy/certificates`

Operation ID: `devices-get-policy-certificates`

Fetches device certificate provisioning.

## Definition

```yaml
{"operationId": "devices-get-policy-certificates", "summary": "Get device certificate provisioning status", "description": "Fetches device certificate provisioning.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "Get WARP client provision certificates enabled status response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_devices_policy_certificates_single"}}}}, "4XX": {"description": "Get WARP client provision certificates enabled status failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_devices_policy_certificates_single"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.policies.default.certificates", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
