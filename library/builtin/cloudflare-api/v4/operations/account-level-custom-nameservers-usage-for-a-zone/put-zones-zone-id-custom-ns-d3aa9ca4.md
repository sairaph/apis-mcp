---
title: Set Account Custom Nameserver Related Zone Metadata
page_id: operation-put-zones-zone-id-custom-ns-7672239e
path: operations/account-level-custom-nameservers-usage-for-a-zone
description: |-
    Set metadata for account-level custom nameservers on a zone.

    If you would like new zones in the account to use account custom nameservers by default, use PUT /accounts/:identifier to set the account setting use_account_custom_ns_by_default to true.

    Deprecated in favor of [Update DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-a-zone-update-dns-settings).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/custom_ns
operation_ids:
    - account-level-custom-nameservers-usage-for-a-zone-set-account-custom-nameserver-related-zone-metadata
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Set Account Custom Nameserver Related Zone Metadata

`PUT /zones/{zone_id}/custom_ns`

Operation ID: `account-level-custom-nameservers-usage-for-a-zone-set-account-custom-nameserver-related-zone-metadata`

Set metadata for account-level custom nameservers on a zone.

If you would like new zones in the account to use account custom nameservers by default, use PUT /accounts/:identifier to set the account setting use_account_custom_ns_by_default to true.

Deprecated in favor of [Update DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-a-zone-update-dns-settings).

## Definition

```yaml
{"operationId": "account-level-custom-nameservers-usage-for-a-zone-set-account-custom-nameserver-related-zone-metadata", "summary": "Set Account Custom Nameserver Related Zone Metadata", "description": "Set metadata for account-level custom nameservers on a zone.\n\nIf you would like new zones in the account to use account custom nameservers by default, use PUT /accounts/:identifier to set the account setting use_account_custom_ns_by_default to true.\n\nDeprecated in favor of [Update DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-a-zone-update-dns-settings).\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-custom-nameservers_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-custom-nameservers_zone_metadata"}}}}, "responses": {"200": {"description": "Set Account Custom Nameserver Related Zone Metadata response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-custom-nameservers_empty_response-2"}}}}, "4XX": {"description": "Set Account Custom Nameserver Related Zone Metadata response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-custom-nameservers_empty_response-2"}, {"$ref": "#/components/schemas/dns-custom-nameservers_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Account-Level Custom Nameservers Usage for a Zone"], "x-api-token-group": ["Zone Write"], "x-cfPermissionsRequired": {"enum": ["#zone:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
