---
title: Update a trusted email domain
page_id: operation-patch-accounts-account-id-email-security-settings-trusted-domains-truste-e4b575c3
path: operations/email-security-settings
description: Updates an existing trusted domain pattern. Only provided fields will be modified. Changes take effect for new emails matching the pattern.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/email-security/settings/trusted_domains/{trusted_domain_id}
operation_ids:
    - email_security_update_trusted_domain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a trusted email domain

`PATCH /accounts/{account_id}/email-security/settings/trusted_domains/{trusted_domain_id}`

Operation ID: `email_security_update_trusted_domain`

Updates an existing trusted domain pattern. Only provided fields will be modified. Changes take effect for new emails matching the pattern.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "trusted_domain_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_TrustedDomainId"}}]
```

## Definition

```yaml
{"operationId": "email_security_update_trusted_domain", "summary": "Update a trusted email domain", "description": "Updates an existing trusted domain pattern. Only provided fields will be modified. Changes take effect for new emails matching the pattern.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-security_UpdateTrustedDomain"}}}}, "responses": {"200": {"description": "Updated trusted domain.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_TrustedDomain"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.settings.trusted-domains", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true, "x-stability": "beta"}
```
