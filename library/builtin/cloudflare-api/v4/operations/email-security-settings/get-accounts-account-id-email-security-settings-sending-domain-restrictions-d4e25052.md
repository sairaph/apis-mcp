---
title: List sending domain restrictions
page_id: operation-get-accounts-account-id-email-security-settings-sending-domain-restricti-c3f19e3c
path: operations/email-security-settings
description: Returns a paginated list of sending domain restrictions. These restrictions enforce TLS requirements for emails from specific domains. Mail without TLS from restricted domains will be dropped unless the subdomain is in the exclude list. Supports sorting and searching.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/settings/sending_domain_restrictions
operation_ids:
    - email_security_list_sending_domain_restrictions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List sending domain restrictions

`GET /accounts/{account_id}/email-security/settings/sending_domain_restrictions`

Operation ID: `email_security_list_sending_domain_restrictions`

Returns a paginated list of sending domain restrictions. These restrictions enforce TLS requirements for emails from specific domains. Mail without TLS from restricted domains will be dropped unless the subdomain is in the exclude list. Supports sorting and searching.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_list_sending_domain_restrictions", "summary": "List sending domain restrictions", "description": "Returns a paginated list of sending domain restrictions. These restrictions enforce TLS requirements for emails from specific domains. Mail without TLS from restricted domains will be dropped unless the subdomain is in the exclude list. Supports sorting and searching.", "parameters": [{"$ref": "#/components/parameters/email-security_page"}, {"$ref": "#/components/parameters/email-security_per_page"}, {"$ref": "#/components/parameters/email-security_search"}, {"name": "order", "in": "query", "description": "Field to sort by.", "schema": {"type": "string", "enum": ["domain", "created_at"]}}, {"$ref": "#/components/parameters/email-security_direction"}], "responses": {"200": {"description": "List of sending domain restrictions.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_SendingDomainRestrictionList"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-stability": "beta"}
```
