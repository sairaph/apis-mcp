---
title: Patch Zero Trust account configuration
page_id: operation-patch-accounts-account-id-gateway-configuration-0b4399c6
path: operations/zero-trust-accounts
description: Update (PATCH) a single subcollection of settings such as `antivirus`, `tls_decrypt`, `activity_log`, `block_page`, `browser_isolation`, `fips`, `body_scanning`, `certificate`, or `max_ttl_secs` without updating the entire configuration object. This endpoint returns an error if any settings collection lacks proper configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/gateway/configuration
operation_ids:
    - zero-trust-accounts-patch-zero-trust-account-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Zero Trust account configuration

`PATCH /accounts/{account_id}/gateway/configuration`

Operation ID: `zero-trust-accounts-patch-zero-trust-account-configuration`

Update (PATCH) a single subcollection of settings such as `antivirus`, `tls_decrypt`, `activity_log`, `block_page`, `browser_isolation`, `fips`, `body_scanning`, `certificate`, or `max_ttl_secs` without updating the entire configuration object. This endpoint returns an error if any settings collection lacks proper configuration.

## Definition

```yaml
{"operationId": "zero-trust-accounts-patch-zero-trust-account-configuration", "summary": "Patch Zero Trust account configuration", "description": "Update (PATCH) a single subcollection of settings such as `antivirus`, `tls_decrypt`, `activity_log`, `block_page`, `browser_isolation`, `fips`, `body_scanning`, `certificate`, or `max_ttl_secs` without updating the entire configuration object. This endpoint returns an error if any settings collection lacks proper configuration.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_gateway-account-settings"}}}}, "responses": {"200": {"description": "Zero Trust account configuration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_gateway_account_config"}}}}, "4XX": {"description": "Zero Trust account configuration response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_gateway_account_config"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust accounts"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.configurations", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
