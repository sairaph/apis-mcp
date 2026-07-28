---
title: Update Zero Trust SSH settings
page_id: operation-put-accounts-account-id-gateway-audit-ssh-settings-7c26c52b
path: operations/zero-trust-ssh-settings
description: Update Zero Trust Audit SSH and SSH with Access for Infrastructure settings for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/gateway/audit_ssh_settings
operation_ids:
    - zero-trust-update-audit-ssh-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Zero Trust SSH settings

`PUT /accounts/{account_id}/gateway/audit_ssh_settings`

Operation ID: `zero-trust-update-audit-ssh-settings`

Update Zero Trust Audit SSH and SSH with Access for Infrastructure settings for an account.

## Definition

```yaml
{"operationId": "zero-trust-update-audit-ssh-settings", "summary": "Update Zero Trust SSH settings", "description": "Update Zero Trust Audit SSH and SSH with Access for Infrastructure settings for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"public_key": {"$ref": "#/components/schemas/zero-trust-gateway_public_key"}}, "required": ["public_key"]}}}}, "responses": {"200": {"description": "Update Zero Trust SSH settings response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response-6"}}}}, "4XX": {"description": "Update Zero Trust SSH settings response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_single_response-6"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust SSH Settings"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.audit-ssh-settings", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
