---
title: Get an impersonation registry entry
page_id: operation-get-accounts-account-id-email-security-settings-impersonation-registry-i-11c53339
path: operations/email-security-settings
description: Retrieves details for a specific impersonation registry entry including the protected identity, email pattern, and synchronization source if directory-synced.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/settings/impersonation_registry/{impersonation_registry_id}
operation_ids:
    - email_security_get_impersonation_registry
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an impersonation registry entry

`GET /accounts/{account_id}/email-security/settings/impersonation_registry/{impersonation_registry_id}`

Operation ID: `email_security_get_impersonation_registry`

Retrieves details for a specific impersonation registry entry including the protected identity, email pattern, and synchronization source if directory-synced.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "impersonation_registry_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_ImpersonationRegistryId"}}]
```

## Definition

```yaml
{"operationId": "email_security_get_impersonation_registry", "summary": "Get an impersonation registry entry", "description": "Retrieves details for a specific impersonation registry entry including the protected identity, email pattern, and synchronization source if directory-synced.", "responses": {"200": {"description": "Impersonation registry entry details.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_ImpersonationRegistry"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-stability": "beta"}
```
