---
title: List entries in impersonation registry
page_id: operation-get-accounts-account-id-email-security-settings-impersonation-registry-d839bddc
path: operations/email-security-settings
description: Returns a paginated list of protected identities in the impersonation registry. These entries define identities and email addresses to protect from impersonation attacks. Can be manually added or automatically synced from directory integrations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/settings/impersonation_registry
operation_ids:
    - email_security_list_impersonation_registry
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List entries in impersonation registry

`GET /accounts/{account_id}/email-security/settings/impersonation_registry`

Operation ID: `email_security_list_impersonation_registry`

Returns a paginated list of protected identities in the impersonation registry. These entries define identities and email addresses to protect from impersonation attacks. Can be manually added or automatically synced from directory integrations.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_list_impersonation_registry", "summary": "List entries in impersonation registry", "description": "Returns a paginated list of protected identities in the impersonation registry. These entries define identities and email addresses to protect from impersonation attacks. Can be manually added or automatically synced from directory integrations.", "parameters": [{"$ref": "#/components/parameters/email-security_page"}, {"$ref": "#/components/parameters/email-security_per_page"}, {"$ref": "#/components/parameters/email-security_search"}, {"name": "order", "in": "query", "description": "Field to sort by.", "schema": {"type": "string", "enum": ["name", "email", "created_at"]}}, {"$ref": "#/components/parameters/email-security_direction"}, {"name": "provenance", "in": "query", "schema": {"$ref": "#/components/schemas/email-security_Provenance"}}], "responses": {"200": {"description": "List of impersonation registry entries.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_ImpersonationRegistryList"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-stability": "beta"}
```
