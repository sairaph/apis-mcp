---
title: List trusted email domains
page_id: operation-get-accounts-account-id-email-security-settings-trusted-domains-9e2ad71d
path: operations/email-security-settings
description: Returns a paginated list of trusted domain patterns. Trusted domains prevent false positives for recently registered domains and lookalike domain detections. Patterns can use regular expressions for flexible matching.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/settings/trusted_domains
operation_ids:
    - email_security_list_trusted_domains
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List trusted email domains

`GET /accounts/{account_id}/email-security/settings/trusted_domains`

Operation ID: `email_security_list_trusted_domains`

Returns a paginated list of trusted domain patterns. Trusted domains prevent false positives for recently registered domains and lookalike domain detections. Patterns can use regular expressions for flexible matching.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_list_trusted_domains", "summary": "List trusted email domains", "description": "Returns a paginated list of trusted domain patterns. Trusted domains prevent false positives for recently registered domains and lookalike domain detections. Patterns can use regular expressions for flexible matching.", "parameters": [{"$ref": "#/components/parameters/email-security_page"}, {"$ref": "#/components/parameters/email-security_per_page"}, {"$ref": "#/components/parameters/email-security_search"}, {"name": "order", "in": "query", "description": "Field to sort by.", "schema": {"type": "string", "enum": ["pattern", "created_at"]}}, {"$ref": "#/components/parameters/email-security_direction"}, {"name": "is_recent", "in": "query", "description": "Filter to show only recently registered domains that are trusted to prevent triggering Suspicious or Malicious dispositions.", "schema": {"type": "boolean"}}, {"name": "is_similarity", "in": "query", "description": "Filter to show only proximity domains (partner or approved domains with similar spelling to connected domains) that prevent Spoof dispositions.", "schema": {"type": "boolean"}}, {"name": "pattern", "in": "query", "schema": {"type": "string"}}], "responses": {"200": {"description": "List of trusted domains.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_TrustedDomainList"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-stability": "beta"}
```
