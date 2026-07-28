---
title: Update an impersonation registry entry
page_id: operation-patch-accounts-account-id-email-security-settings-impersonation-registry-366c7ebb
path: operations/email-security-settings
description: Updates an existing impersonation registry entry. Only provided fields will be modified. Directory-synced entries can't be updated.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/email-security/settings/impersonation_registry/{impersonation_registry_id}
operation_ids:
    - email_security_update_impersonation_registry
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an impersonation registry entry

`PATCH /accounts/{account_id}/email-security/settings/impersonation_registry/{impersonation_registry_id}`

Operation ID: `email_security_update_impersonation_registry`

Updates an existing impersonation registry entry. Only provided fields will be modified. Directory-synced entries can't be updated.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "impersonation_registry_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_ImpersonationRegistryId"}}]
```

## Definition

```yaml
{"operationId": "email_security_update_impersonation_registry", "summary": "Update an impersonation registry entry", "description": "Updates an existing impersonation registry entry. Only provided fields will be modified. Directory-synced entries can't be updated.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-security_UpdateImpersonationRegistry"}}}}, "responses": {"200": {"description": "Updated impersonation registry entry.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_ImpersonationRegistry"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
