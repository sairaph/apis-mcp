---
title: Delete an impersonation registry entry
page_id: operation-delete-accounts-account-id-email-security-settings-impersonation-registr-7c01189e
path: operations/email-security-settings
description: Removes an entry from the impersonation registry. After deletion, this identity will no longer be protected from impersonation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/email-security/settings/impersonation_registry/{impersonation_registry_id}
operation_ids:
    - email_security_delete_impersonation_registry
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an impersonation registry entry

`DELETE /accounts/{account_id}/email-security/settings/impersonation_registry/{impersonation_registry_id}`

Operation ID: `email_security_delete_impersonation_registry`

Removes an entry from the impersonation registry. After deletion, this identity will no longer be protected from impersonation.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "impersonation_registry_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_ImpersonationRegistryId"}}]
```

## Definition

```yaml
{"operationId": "email_security_delete_impersonation_registry", "summary": "Delete an impersonation registry entry", "description": "Removes an entry from the impersonation registry. After deletion, this identity will no longer be protected from impersonation.", "responses": {"200": {"description": "Deleted impersonation registry entry.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_DeletedImpersonationRegistry"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
