---
title: Create a sending domain restriction
page_id: operation-post-accounts-account-id-email-security-settings-sending-domain-restrict-cf30f267
path: operations/email-security-settings
description: Creates a new sending domain restriction to enforce TLS requirements for a domain. Emails without TLS from this domain will be dropped unless the subdomain is in the exclude list.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email-security/settings/sending_domain_restrictions
operation_ids:
    - email_security_create_sending_domain_restriction
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a sending domain restriction

`POST /accounts/{account_id}/email-security/settings/sending_domain_restrictions`

Operation ID: `email_security_create_sending_domain_restriction`

Creates a new sending domain restriction to enforce TLS requirements for a domain. Emails without TLS from this domain will be dropped unless the subdomain is in the exclude list.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_create_sending_domain_restriction", "summary": "Create a sending domain restriction", "description": "Creates a new sending domain restriction to enforce TLS requirements for a domain. Emails without TLS from this domain will be dropped unless the subdomain is in the exclude list.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-security_CreateSendingDomainRestriction"}}}}, "responses": {"201": {"description": "Created sending domain restriction.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_SendingDomainRestriction"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
