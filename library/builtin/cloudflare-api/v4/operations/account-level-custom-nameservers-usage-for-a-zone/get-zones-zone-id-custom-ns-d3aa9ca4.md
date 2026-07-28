---
title: Get Account Custom Nameserver Related Zone Metadata
page_id: operation-get-zones-zone-id-custom-ns-23b01fc3
path: operations/account-level-custom-nameservers-usage-for-a-zone
description: |-
    Get metadata for account-level custom nameservers on a zone.

    Deprecated in favor of [Show DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-a-zone-list-dns-settings).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/custom_ns
operation_ids:
    - account-level-custom-nameservers-usage-for-a-zone-get-account-custom-nameserver-related-zone-metadata
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Account Custom Nameserver Related Zone Metadata

`GET /zones/{zone_id}/custom_ns`

Operation ID: `account-level-custom-nameservers-usage-for-a-zone-get-account-custom-nameserver-related-zone-metadata`

Get metadata for account-level custom nameservers on a zone.

Deprecated in favor of [Show DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-a-zone-list-dns-settings).

## Definition

```yaml
{"operationId": "account-level-custom-nameservers-usage-for-a-zone-get-account-custom-nameserver-related-zone-metadata", "summary": "Get Account Custom Nameserver Related Zone Metadata", "description": "Get metadata for account-level custom nameservers on a zone.\n\nDeprecated in favor of [Show DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-a-zone-list-dns-settings).\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-custom-nameservers_identifier-2"}}], "responses": {"200": {"description": "Get Account Custom Nameserver Related Zone Metadata response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-custom-nameservers_get_response"}}}}, "4XX": {"description": "Get Account Custom Nameserver Related Zone Metadata response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-custom-nameservers_get_response"}, {"$ref": "#/components/schemas/dns-custom-nameservers_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Account-Level Custom Nameservers Usage for a Zone"], "x-api-token-group": ["Trust and Safety Write", "Trust and Safety Read", "Zero Trust: PII Read", "Zaraz Edit", "Zaraz Read", "Zaraz Admin", "Access: Apps and Policies Revoke", "Access: Apps and Policies Write", "Access: Apps and Policies Read", "Access: Apps and Policies Revoke", "Access: Mutual TLS Certificates Write", "Access: Organizations, Identity Providers, and Groups Write", "Zone Settings Write", "Zone Settings Read", "Zone Read", "DNS Read", "Workers Scripts Write", "Workers Scripts Read", "Zone Write", "Workers Routes Write", "Workers Routes Read", "Stream Write", "Stream Read", "SSL and Certificates Write", "SSL and Certificates Read", "Logs Write", "Logs Read", "Cache Purge", "Page Rules Write", "Page Rules Read", "Load Balancers Write", "Load Balancers Read", "Firewall Services Write", "Firewall Services Read", "DNS Write", "Apps Write", "Analytics Read", "Access: Apps and Policies Write", "Access: Apps and Policies Read"], "x-cfPermissionsRequired": {"enum": ["#zone:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
