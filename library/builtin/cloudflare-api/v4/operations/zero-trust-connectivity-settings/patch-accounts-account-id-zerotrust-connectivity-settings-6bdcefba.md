---
title: Updates the Zero Trust Connectivity Settings
page_id: operation-patch-accounts-account-id-zerotrust-connectivity-settings-3aebbae2
path: operations/zero-trust-connectivity-settings
description: Updates the Zero Trust Connectivity Settings for the given account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/zerotrust/connectivity_settings
operation_ids:
    - zero-trust-accounts-patch-connectivity-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Updates the Zero Trust Connectivity Settings

`PATCH /accounts/{account_id}/zerotrust/connectivity_settings`

Operation ID: `zero-trust-accounts-patch-connectivity-settings`

Updates the Zero Trust Connectivity Settings for the given account.

## Definition

```yaml
{"operationId": "zero-trust-accounts-patch-connectivity-settings", "summary": "Updates the Zero Trust Connectivity Settings", "description": "Updates the Zero Trust Connectivity Settings for the given account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"icmp_proxy_enabled": {"$ref": "#/components/schemas/tunnel_icmp_proxy_enabled"}, "offramp_warp_enabled": {"$ref": "#/components/schemas/tunnel_offramp_warp_enabled"}}}}}}, "responses": {"200": {"description": "Update Zero Trust Connectivity Settings response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_zero_trust_connectivity_settings_response"}}}}, "4XX": {"description": "Update Zero Trust Connectivity Settings response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Zero Trust Connectivity Settings"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.connectivity-settings", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
