---
title: Unprotect an email domain
page_id: operation-delete-accounts-account-id-email-security-settings-domains-domain-id-bd8d7e90
path: operations/email-security-settings
description: Removes email security protection from a domain. After deletion, emails for this domain will no longer be processed by Email Security. This action cannot be undone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/email-security/settings/domains/{domain_id}
operation_ids:
    - email_security_delete_domain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Unprotect an email domain

`DELETE /accounts/{account_id}/email-security/settings/domains/{domain_id}`

Operation ID: `email_security_delete_domain`

Removes email security protection from a domain. After deletion, emails for this domain will no longer be processed by Email Security. This action cannot be undone.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "domain_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_DomainId"}}]
```

## Definition

```yaml
{"operationId": "email_security_delete_domain", "summary": "Unprotect an email domain", "description": "Removes email security protection from a domain. After deletion, emails for this domain will no longer be processed by Email Security. This action cannot be undone.", "responses": {"200": {"description": "Deleted domain.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_DeletedDomain"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
