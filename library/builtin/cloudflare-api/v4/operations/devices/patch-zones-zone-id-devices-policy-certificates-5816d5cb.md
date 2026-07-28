---
title: Update device certificate provisioning status
page_id: operation-patch-zones-zone-id-devices-policy-certificates-5a9e9272
path: operations/devices
description: Enable Zero Trust Clients to provision a certificate, containing a x509 subject, and referenced by Access device posture policies when the client visits MTLS protected domains. This facilitates device posture without a WARP session.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/devices/policy/certificates
operation_ids:
    - devices-update-policy-certificates
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update device certificate provisioning status

`PATCH /zones/{zone_id}/devices/policy/certificates`

Operation ID: `devices-update-policy-certificates`

Enable Zero Trust Clients to provision a certificate, containing a x509 subject, and referenced by Access device posture policies when the client visits MTLS protected domains. This facilitates device posture without a WARP session.

## Definition

```yaml
{"operationId": "devices-update-policy-certificates", "summary": "Update device certificate provisioning status", "description": "Enable Zero Trust Clients to provision a certificate, containing a x509 subject, and referenced by Access device posture policies when the client visits MTLS protected domains. This facilitates device posture without a WARP session.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_devices_policy_certificates"}}}}, "responses": {"200": {"description": "Update a zone to toggle permission for devices to provision certificates response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_devices_policy_certificates_single"}}}}, "4XX": {"description": "Patch a zone to toggle permission for devices to provision certificates failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_devices_policy_certificates_single"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-api-token-group": ["SSL and Certificates Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.policies.default.certificates", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
