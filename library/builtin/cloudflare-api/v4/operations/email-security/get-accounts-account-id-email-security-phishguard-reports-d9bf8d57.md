---
title: Get PhishGuard reports
page_id: operation-get-accounts-account-id-email-security-phishguard-reports-ea77cb76
path: operations/email-security
description: Retrieves PhishGuard security alert reports for a specified date range. Reports include detected threats, dispositions, and contextual information. Use for security monitoring and threat analysis.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/phishguard/reports
operation_ids:
    - email_security_get_phishguard_reports
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get PhishGuard reports

`GET /accounts/{account_id}/email-security/phishguard/reports`

Operation ID: `email_security_get_phishguard_reports`

Retrieves PhishGuard security alert reports for a specified date range. Reports include detected threats, dispositions, and contextual information. Use for security monitoring and threat analysis.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_get_phishguard_reports", "summary": "Get PhishGuard reports", "description": "Retrieves PhishGuard security alert reports for a specified date range. Reports include detected threats, dispositions, and contextual information. Use for security monitoring and threat analysis.", "parameters": [{"name": "start", "in": "query", "description": "Start of the time range (RFC3339). Takes precedence over from_date.", "schema": {"type": "string", "format": "date-time"}, "example": "2020-08-01T00:00:00Z"}, {"name": "end", "in": "query", "description": "End of the time range (RFC3339). Takes precedence over to_date.", "schema": {"type": "string", "format": "date-time"}, "example": "2020-09-01T00:00:00Z"}, {"name": "from_date", "in": "query", "description": "Deprecated, use `start` instead. Start date in YYYY-MM-DD format.", "schema": {"type": "string", "format": "date"}, "deprecated": true, "example": "2020-08-01", "x-stainless-deprecation-message": "Use `start` instead.", "x-sunset": "2026-11-01"}, {"name": "to_date", "in": "query", "description": "Deprecated, use `end` instead. End date in YYYY-MM-DD format.", "schema": {"type": "string", "format": "date"}, "deprecated": true, "example": "2020-09-01", "x-stainless-deprecation-message": "Use `end` instead.", "x-sunset": "2026-11-01"}], "responses": {"200": {"description": "List of PhishGuard reports.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_PhishGuardReport"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-stability": "beta"}
```
