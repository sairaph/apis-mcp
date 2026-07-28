---
title: Delete a trusted email domain
page_id: operation-delete-accounts-account-id-email-security-settings-trusted-domains-trust-834e0c56
path: operations/email-security-settings
description: Removes a trusted domain pattern. After deletion, emails from this domain will be subject to normal recent domain and similarity checks.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/email-security/settings/trusted_domains/{trusted_domain_id}
operation_ids:
    - email_security_delete_trusted_domain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a trusted email domain

`DELETE /accounts/{account_id}/email-security/settings/trusted_domains/{trusted_domain_id}`

Operation ID: `email_security_delete_trusted_domain`

Removes a trusted domain pattern. After deletion, emails from this domain will be subject to normal recent domain and similarity checks.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "trusted_domain_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_TrustedDomainId"}}]
```

## Definition

```yaml
{"operationId": "email_security_delete_trusted_domain", "summary": "Delete a trusted email domain", "description": "Removes a trusted domain pattern. After deletion, emails from this domain will be subject to normal recent domain and similarity checks.", "responses": {"200": {"description": "Deleted trusted domain.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_DeletedTrustedDomain"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.settings.trusted-domains", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true, "x-stability": "beta"}
```
