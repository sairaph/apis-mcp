---
title: Get domain verification details
page_id: operation-get-accounts-account-id-email-security-settings-domains-domain-id-verifi-08b06612
path: operations/email-security-settings
description: |-
    Returns DNS TXT verification details for a domain created through the Terraform API.
    Returns `200 OK` with the TXT record name and value when the domain uses `dns_txt`
    verification. Returns `204 No Content` for pre-existing domains or domains verified
    by other means (Cloudflare zone ownership, CASB integration).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/settings/domains/{domain_id}/verification
operation_ids:
    - email_security_get_domain_verification
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get domain verification details

`GET /accounts/{account_id}/email-security/settings/domains/{domain_id}/verification`

Operation ID: `email_security_get_domain_verification`

Returns DNS TXT verification details for a domain created through the Terraform API.
Returns `200 OK` with the TXT record name and value when the domain uses `dns_txt`
verification. Returns `204 No Content` for pre-existing domains or domains verified
by other means (Cloudflare zone ownership, CASB integration).

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "domain_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_DomainId"}}]
```

## Definition

```yaml
{"operationId": "email_security_get_domain_verification", "summary": "Get domain verification details", "description": "Returns DNS TXT verification details for a domain created through the Terraform API.\nReturns `200 OK` with the TXT record name and value when the domain uses `dns_txt`\nverification. Returns `204 No Content` for pre-existing domains or domains verified\nby other means (Cloudflare zone ownership, CASB integration).\n", "responses": {"200": {"description": "DNS TXT verification details.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_DomainVerificationInfo"}}, "type": "object"}]}}}}, "204": {"description": "Domain is verified by other means; no TXT record required."}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-stability": "beta"}
```
