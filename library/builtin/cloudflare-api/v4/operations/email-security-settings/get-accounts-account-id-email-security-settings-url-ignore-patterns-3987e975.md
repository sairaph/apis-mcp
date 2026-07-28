---
title: List URL ignore patterns
page_id: operation-get-accounts-account-id-email-security-settings-url-ignore-patterns-18fb1362
path: operations/email-security-settings
description: Returns a paginated list of URL rewrite ignore patterns for the account. URLs matching these patterns will not be rewritten.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/settings/url_ignore_patterns
operation_ids:
    - email_security_list_url_ignore_patterns
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List URL ignore patterns

`GET /accounts/{account_id}/email-security/settings/url_ignore_patterns`

Operation ID: `email_security_list_url_ignore_patterns`

Returns a paginated list of URL rewrite ignore patterns for the account. URLs matching these patterns will not be rewritten.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_list_url_ignore_patterns", "summary": "List URL ignore patterns", "description": "Returns a paginated list of URL rewrite ignore patterns for the account. URLs matching these patterns will not be rewritten.", "parameters": [{"$ref": "#/components/parameters/email-security_page"}, {"$ref": "#/components/parameters/email-security_per_page"}], "responses": {"200": {"description": "List of URL ignore patterns for the provided account.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_UrlIgnorePattern"}}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-stability": "beta"}
```
