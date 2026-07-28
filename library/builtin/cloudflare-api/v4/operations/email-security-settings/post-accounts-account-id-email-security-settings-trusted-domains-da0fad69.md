---
title: Create trusted email domain
page_id: operation-post-accounts-account-id-email-security-settings-trusted-domains-c9dd7ecd
path: operations/email-security-settings
description: Creates a new trusted domain pattern. Use for partner domains or approved senders that should bypass recent domain registration and similarity checks. Configure whether it prevents recent domain or spoof dispositions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email-security/settings/trusted_domains
operation_ids:
    - email_security_create_trusted_domain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create trusted email domain

`POST /accounts/{account_id}/email-security/settings/trusted_domains`

Operation ID: `email_security_create_trusted_domain`

Creates a new trusted domain pattern. Use for partner domains or approved senders that should bypass recent domain registration and similarity checks. Configure whether it prevents recent domain or spoof dispositions.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_create_trusted_domain", "summary": "Create trusted email domain", "description": "Creates a new trusted domain pattern. Use for partner domains or approved senders that should bypass recent domain registration and similarity checks. Configure whether it prevents recent domain or spoof dispositions.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-security_CreateTrustedDomain"}}}}, "responses": {"201": {"description": "Created trusted domain.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_TrustedDomain"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.settings.trusted-domains", "x-fern-sdk-method-name": "create", "x-forge-hidden": true, "x-stability": "beta"}
```
