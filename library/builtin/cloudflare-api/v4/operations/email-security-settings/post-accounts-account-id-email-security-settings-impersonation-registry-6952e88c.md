---
title: Create impersonation registry entry
page_id: operation-post-accounts-account-id-email-security-settings-impersonation-registry-cfaafaa6
path: operations/email-security-settings
description: Creates a new entry in the impersonation registry to protect against impersonation. Emails attempting to impersonate this identity will be flagged. Supports regex patterns for flexible email matching.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email-security/settings/impersonation_registry
operation_ids:
    - email_security_create_impersonation_registry
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create impersonation registry entry

`POST /accounts/{account_id}/email-security/settings/impersonation_registry`

Operation ID: `email_security_create_impersonation_registry`

Creates a new entry in the impersonation registry to protect against impersonation. Emails attempting to impersonate this identity will be flagged. Supports regex patterns for flexible email matching.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_create_impersonation_registry", "summary": "Create impersonation registry entry", "description": "Creates a new entry in the impersonation registry to protect against impersonation. Emails attempting to impersonate this identity will be flagged. Supports regex patterns for flexible email matching.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-security_CreateImpersonationRegistry"}}}}, "responses": {"201": {"description": "Created impersonation registry entry.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_ImpersonationRegistry"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
