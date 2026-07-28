---
title: List blocked email senders
page_id: operation-get-accounts-account-id-email-security-settings-block-senders-70cf11c1
path: operations/email-security-settings
description: Returns a paginated list of blocked email sender patterns. These patterns prevent emails from matching senders from being delivered. Supports filtering by pattern type and searching across patterns.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/settings/block_senders
operation_ids:
    - email_security_list_blocked_senders
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List blocked email senders

`GET /accounts/{account_id}/email-security/settings/block_senders`

Operation ID: `email_security_list_blocked_senders`

Returns a paginated list of blocked email sender patterns. These patterns prevent emails from matching senders from being delivered. Supports filtering by pattern type and searching across patterns.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_list_blocked_senders", "summary": "List blocked email senders", "description": "Returns a paginated list of blocked email sender patterns. These patterns prevent emails from matching senders from being delivered. Supports filtering by pattern type and searching across patterns.", "parameters": [{"$ref": "#/components/parameters/email-security_page"}, {"$ref": "#/components/parameters/email-security_per_page"}, {"$ref": "#/components/parameters/email-security_search"}, {"name": "order", "in": "query", "description": "Field to sort by.", "schema": {"type": "string", "enum": ["pattern", "created_at"]}}, {"$ref": "#/components/parameters/email-security_direction"}, {"name": "pattern_type", "in": "query", "description": "Filter by pattern type.", "schema": {"allOf": [{"$ref": "#/components/schemas/email-security_PatternType"}]}}, {"name": "pattern", "in": "query", "description": "Filter by pattern value.", "schema": {"type": "string"}}], "responses": {"200": {"description": "List of blocked senders.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_BlockedSenderList"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-stability": "beta"}
```
