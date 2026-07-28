---
title: Get CSAM Scanner setting
page_id: operation-get-zones-zone-id-settings-csam-scanner-third-party-2e7334eb
path: operations/csam-scanner-settings
description: |-
    Retrieve the current CSAM Scanner configuration for a zone.

    The notification email is masked by default in responses.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/settings/csam_scanner_third_party
operation_ids:
    - csam-scanner-get-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get CSAM Scanner setting

`GET /zones/{zone_id}/settings/csam_scanner_third_party`

Operation ID: `csam-scanner-get-setting`

Retrieve the current CSAM Scanner configuration for a zone.

The notification email is masked by default in responses.

## Definition

```yaml
{"operationId": "csam-scanner-get-setting", "summary": "Get CSAM Scanner setting", "description": "Retrieve the current CSAM Scanner configuration for a zone.\n\nThe notification email is masked by default in responses.\n", "parameters": [{"$ref": "#/components/parameters/csam-config-service_zone_id"}], "responses": {"200": {"description": "CSAM Scanner setting response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/csam-config-service_csam_scanner_single_response"}}}}, "400": {"description": "Bad request - malformed zone ID in URL", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/csam-config-service_api_response_common_failure"}}}}, "401": {"description": "Unauthorized", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/csam-config-service_api_response_common_failure"}}}}, "403": {"description": "Forbidden", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/csam-config-service_api_response_common_failure"}}}}, "404": {"description": "Zone not found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/csam-config-service_api_response_common_failure"}}}}, "500": {"description": "Internal server error", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/csam-config-service_api_response_common_failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["CSAM Scanner Settings"], "x-api-token-group": ["Zone Settings Read"]}
```
