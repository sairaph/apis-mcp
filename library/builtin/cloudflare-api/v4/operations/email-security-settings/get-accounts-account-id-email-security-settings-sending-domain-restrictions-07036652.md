---
title: Get a sending domain restriction
page_id: operation-get-accounts-account-id-email-security-settings-sending-domain-restricti-4f3c6bba
path: operations/email-security-settings
description: Retrieves details for a specific sending domain restriction including the domain requiring TLS and any excluded subdomains exempt from the TLS requirement.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/settings/sending_domain_restrictions/{sending_domain_restriction_id}
operation_ids:
    - email_security_get_sending_domain_restriction
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a sending domain restriction

`GET /accounts/{account_id}/email-security/settings/sending_domain_restrictions/{sending_domain_restriction_id}`

Operation ID: `email_security_get_sending_domain_restriction`

Retrieves details for a specific sending domain restriction including the domain requiring TLS and any excluded subdomains exempt from the TLS requirement.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "sending_domain_restriction_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_SendingDomainRestrictionId"}}]
```

## Definition

```yaml
{"operationId": "email_security_get_sending_domain_restriction", "summary": "Get a sending domain restriction", "description": "Retrieves details for a specific sending domain restriction including the domain requiring TLS and any excluded subdomains exempt from the TLS requirement.", "responses": {"200": {"description": "Sending domain restriction details.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_SendingDomainRestriction"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-stability": "beta"}
```
