---
title: Get Zero Trust SSH settings
page_id: operation-get-accounts-account-id-gateway-audit-ssh-settings-d706270d
path: operations/zero-trust-ssh-settings
description: Retrieve all Zero Trust Audit SSH and SSH with Access for Infrastructure settings for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/audit_ssh_settings
operation_ids:
    - zero-trust-get-audit-ssh-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Zero Trust SSH settings

`GET /accounts/{account_id}/gateway/audit_ssh_settings`

Operation ID: `zero-trust-get-audit-ssh-settings`

Retrieve all Zero Trust Audit SSH and SSH with Access for Infrastructure settings for an account.

## Definition

```yaml
{"operationId": "zero-trust-get-audit-ssh-settings", "summary": "Get Zero Trust SSH settings", "description": "Retrieve all Zero Trust Audit SSH and SSH with Access for Infrastructure settings for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "responses": {"200": {"description": "Get Zero Trust SSH settings response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response-6"}}}}, "4XX": {"description": "Get Zero Trust SSH settings response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_single_response-6"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust SSH Settings"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.audit-ssh-settings", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
