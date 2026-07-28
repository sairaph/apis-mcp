---
title: Unprotect multiple email domains
page_id: operation-delete-accounts-account-id-email-security-settings-domains-f0626156
path: operations/email-security-settings
description: Removes protection from multiple email domains. Deprecated; use the batch endpoint instead.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/email-security/settings/domains
operation_ids:
    - email_security_delete_domains
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Unprotect multiple email domains

`DELETE /accounts/{account_id}/email-security/settings/domains`

Operation ID: `email_security_delete_domains`

Removes protection from multiple email domains. Deprecated; use the batch endpoint instead.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_delete_domains", "summary": "Unprotect multiple email domains", "description": "Removes protection from multiple email domains. Deprecated; use the batch endpoint instead.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"properties": {"id": {"$ref": "#/components/schemas/email-security_DomainId"}}, "required": ["id"], "type": "object"}, "minItems": 1}}}}, "responses": {"200": {"description": "Deleted domains.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_DeletedDomain"}}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
