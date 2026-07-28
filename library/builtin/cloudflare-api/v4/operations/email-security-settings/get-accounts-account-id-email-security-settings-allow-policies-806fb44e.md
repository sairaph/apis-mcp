---
title: List email allow policies
page_id: operation-get-accounts-account-id-email-security-settings-allow-policies-17c35b13
path: operations/email-security-settings
description: Returns a paginated list of email allow policies. These policies exempt matching emails from security detection, allowing them to bypass disposition actions. Supports filtering by pattern type and policy attributes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/settings/allow_policies
operation_ids:
    - email_security_list_allow_policies
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List email allow policies

`GET /accounts/{account_id}/email-security/settings/allow_policies`

Operation ID: `email_security_list_allow_policies`

Returns a paginated list of email allow policies. These policies exempt matching emails from security detection, allowing them to bypass disposition actions. Supports filtering by pattern type and policy attributes.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_list_allow_policies", "summary": "List email allow policies", "description": "Returns a paginated list of email allow policies. These policies exempt matching emails from security detection, allowing them to bypass disposition actions. Supports filtering by pattern type and policy attributes.", "parameters": [{"$ref": "#/components/parameters/email-security_page"}, {"$ref": "#/components/parameters/email-security_per_page"}, {"$ref": "#/components/parameters/email-security_search"}, {"name": "order", "in": "query", "description": "Field to sort by.", "schema": {"type": "string", "enum": ["pattern", "created_at"]}}, {"$ref": "#/components/parameters/email-security_direction"}, {"name": "is_exempt_recipient", "in": "query", "description": "Filter to show only policies where messages to the recipient bypass all detections.", "schema": {"type": "boolean"}}, {"name": "is_trusted_sender", "in": "query", "description": "Filter to show only policies where messages from the sender bypass all detections and link following.", "schema": {"type": "boolean"}}, {"name": "is_acceptable_sender", "in": "query", "description": "Filter to show only policies where messages from the sender are exempted from Spam, Spoof, and Bulk dispositions (not Malicious or Suspicious).", "schema": {"type": "boolean"}}, {"name": "verify_sender", "in": "query", "description": "Filter to show only policies that enforce DMARC, SPF, or DKIM authentication.", "schema": {"type": "boolean"}}, {"name": "pattern_type", "in": "query", "schema": {"allOf": [{"$ref": "#/components/schemas/email-security_PatternType"}]}}, {"name": "pattern", "in": "query", "schema": {"type": "string"}}], "responses": {"200": {"description": "List of allow policies.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_AllowPolicyList"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.settings.allow-policies", "x-fern-sdk-method-name": "list", "x-forge-hidden": true, "x-stability": "beta"}
```
