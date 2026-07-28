---
title: Update CSAM Scanner setting
page_id: operation-patch-zones-zone-id-settings-csam-scanner-third-party-df1dfb32
path: operations/csam-scanner-settings
description: |-
    Update the CSAM Scanner configuration for a zone. Allows enabling or
    disabling CSAM scanning, updating the notification email, and
    configuring scanning sources.

    When a new email is provided, email verification is triggered
    automatically. The `enabled` field is a toggle; the server may
    adjust it based on whether the notification email is verified.

    Returns 403 if the zone or account is locked by Trust & Safety.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/settings/csam_scanner_third_party
operation_ids:
    - csam-scanner-update-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update CSAM Scanner setting

`PATCH /zones/{zone_id}/settings/csam_scanner_third_party`

Operation ID: `csam-scanner-update-setting`

Update the CSAM Scanner configuration for a zone. Allows enabling or
disabling CSAM scanning, updating the notification email, and
configuring scanning sources.

When a new email is provided, email verification is triggered
automatically. The `enabled` field is a toggle; the server may
adjust it based on whether the notification email is verified.

Returns 403 if the zone or account is locked by Trust & Safety.

## Definition

```yaml
{"operationId": "csam-scanner-update-setting", "summary": "Update CSAM Scanner setting", "description": "Update the CSAM Scanner configuration for a zone. Allows enabling or\ndisabling CSAM scanning, updating the notification email, and\nconfiguring scanning sources.\n\nWhen a new email is provided, email verification is triggered\nautomatically. The `enabled` field is a toggle; the server may\nadjust it based on whether the notification email is verified.\n\nReturns 403 if the zone or account is locked by Trust & Safety.\n", "parameters": [{"$ref": "#/components/parameters/csam-config-service_zone_id"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/csam-config-service_csam_scanner_third_party_update_request"}}}}, "responses": {"200": {"description": "CSAM Scanner setting updated", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/csam-config-service_csam_scanner_single_response"}}}}, "400": {"description": "Bad request - validation failure or malformed request body", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/csam-config-service_api_response_common_failure"}}}}, "401": {"description": "Unauthorized", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/csam-config-service_api_response_common_failure"}}}}, "403": {"description": "Forbidden - zone or account is locked", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/csam-config-service_api_response_common_failure"}}}}, "404": {"description": "Zone not found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/csam-config-service_api_response_common_failure"}}}}, "500": {"description": "Internal server error", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/csam-config-service_api_response_common_failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["CSAM Scanner Settings"], "x-api-token-group": ["Zone Settings Write"]}
```
