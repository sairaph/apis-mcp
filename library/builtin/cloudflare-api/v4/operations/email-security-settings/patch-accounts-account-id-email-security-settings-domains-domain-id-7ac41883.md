---
title: Update an email domain
page_id: operation-patch-accounts-account-id-email-security-settings-domains-domain-id-309dc7ec
path: operations/email-security-settings
description: Updates configuration for a protected email domain. Only provided fields will be modified. Changes affect delivery mode, security settings, and regional processing.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/email-security/settings/domains/{domain_id}
operation_ids:
    - email_security_update_domain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an email domain

`PATCH /accounts/{account_id}/email-security/settings/domains/{domain_id}`

Operation ID: `email_security_update_domain`

Updates configuration for a protected email domain. Only provided fields will be modified. Changes affect delivery mode, security settings, and regional processing.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "domain_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_DomainId"}}]
```

## Definition

```yaml
{"operationId": "email_security_update_domain", "summary": "Update an email domain", "description": "Updates configuration for a protected email domain. Only provided fields will be modified. Changes affect delivery mode, security settings, and regional processing.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-security_UpdateDomain"}}}}, "responses": {"200": {"description": "Updated domain.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_Domain"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
