---
title: Rotate Zero Trust SSH account seed
page_id: operation-post-accounts-account-id-gateway-audit-ssh-settings-rotate-seed-d488c4e4
path: operations/zero-trust-ssh-settings
description: Rotate the SSH account seed that generates the host key identity when connecting through the Cloudflare SSH Proxy.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/gateway/audit_ssh_settings/rotate_seed
operation_ids:
    - zero-trust-rotate-ssh-account-seed
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Rotate Zero Trust SSH account seed

`POST /accounts/{account_id}/gateway/audit_ssh_settings/rotate_seed`

Operation ID: `zero-trust-rotate-ssh-account-seed`

Rotate the SSH account seed that generates the host key identity when connecting through the Cloudflare SSH Proxy.

## Definition

```yaml
{"operationId": "zero-trust-rotate-ssh-account-seed", "summary": "Rotate Zero Trust SSH account seed", "description": "Rotate the SSH account seed that generates the host key identity when connecting through the Cloudflare SSH Proxy.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "responses": {"200": {"description": "Rotate Zero Trust SSH account seed response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response-6"}}}}, "4XX": {"description": "Rotate Zero Trust SSH account seed response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_single_response-6"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust SSH Settings"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.audit-ssh-settings", "x-fern-sdk-method-name": "rotate-seed", "x-forge-hidden": true}
```
