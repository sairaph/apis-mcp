---
title: Get Zero Trust Connectivity Settings
page_id: operation-get-accounts-account-id-zerotrust-connectivity-settings-16fe6e47
path: operations/zero-trust-connectivity-settings
description: Gets the Zero Trust Connectivity Settings for the given account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/zerotrust/connectivity_settings
operation_ids:
    - zero-trust-accounts-get-connectivity-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Zero Trust Connectivity Settings

`GET /accounts/{account_id}/zerotrust/connectivity_settings`

Operation ID: `zero-trust-accounts-get-connectivity-settings`

Gets the Zero Trust Connectivity Settings for the given account.

## Definition

```yaml
{"operationId": "zero-trust-accounts-get-connectivity-settings", "summary": "Get Zero Trust Connectivity Settings", "description": "Gets the Zero Trust Connectivity Settings for the given account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}], "responses": {"200": {"description": "Get Zero Trust Connectivity Settings response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_zero_trust_connectivity_settings_response"}}}}, "4XX": {"description": "Get Zero Trust Connectivity Settings response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Zero Trust Connectivity Settings"], "x-api-token-group": ["Zero Trust Report", "Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.connectivity-settings", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
