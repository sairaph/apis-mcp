---
title: Get an email domain
page_id: operation-get-accounts-account-id-email-security-settings-domains-domain-id-112e822b
path: operations/email-security-settings
description: Retrieves detailed information for a specific protected email domain including its delivery configuration, SPF/DMARC status, and authorization state.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/settings/domains/{domain_id}
operation_ids:
    - email_security_get_domain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an email domain

`GET /accounts/{account_id}/email-security/settings/domains/{domain_id}`

Operation ID: `email_security_get_domain`

Retrieves detailed information for a specific protected email domain including its delivery configuration, SPF/DMARC status, and authorization state.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "domain_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_DomainId"}}]
```

## Definition

```yaml
{"operationId": "email_security_get_domain", "summary": "Get an email domain", "description": "Retrieves detailed information for a specific protected email domain including its delivery configuration, SPF/DMARC status, and authorization state.", "responses": {"200": {"description": "Domain details.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_Domain"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.settings.domains", "x-fern-sdk-method-name": "get", "x-forge-hidden": true, "x-stability": "beta"}
```
