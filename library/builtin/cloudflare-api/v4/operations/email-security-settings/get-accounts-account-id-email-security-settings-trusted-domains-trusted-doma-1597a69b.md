---
title: Get a trusted email domain
page_id: operation-get-accounts-account-id-email-security-settings-trusted-domains-trusted-87def99e
path: operations/email-security-settings
description: Retrieves details for a specific trusted domain pattern including its pattern value, whether it uses regex matching, and which detection types it affects.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/settings/trusted_domains/{trusted_domain_id}
operation_ids:
    - email_security_get_trusted_domain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a trusted email domain

`GET /accounts/{account_id}/email-security/settings/trusted_domains/{trusted_domain_id}`

Operation ID: `email_security_get_trusted_domain`

Retrieves details for a specific trusted domain pattern including its pattern value, whether it uses regex matching, and which detection types it affects.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "trusted_domain_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_TrustedDomainId"}}]
```

## Definition

```yaml
{"operationId": "email_security_get_trusted_domain", "summary": "Get a trusted email domain", "description": "Retrieves details for a specific trusted domain pattern including its pattern value, whether it uses regex matching, and which detection types it affects.", "responses": {"200": {"description": "Trusted domain details.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_TrustedDomain"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.settings.trusted-domains", "x-fern-sdk-method-name": "get", "x-forge-hidden": true, "x-stability": "beta"}
```
