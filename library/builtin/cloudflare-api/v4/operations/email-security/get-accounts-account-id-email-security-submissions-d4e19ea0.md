---
title: Get reclassify submissions
page_id: operation-get-accounts-account-id-email-security-submissions-c8c3a8e2
path: operations/email-security
description: Returns information for submissions made to reclassify emails. Shows the status, outcome, and disposition changes for reclassification requests made by users or the security team. Useful for tracking false positive/negative reports.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/submissions
operation_ids:
    - email_security_submissions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get reclassify submissions

`GET /accounts/{account_id}/email-security/submissions`

Operation ID: `email_security_submissions`

Returns information for submissions made to reclassify emails. Shows the status, outcome, and disposition changes for reclassification requests made by users or the security team. Useful for tracking false positive/negative reports.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_submissions", "summary": "Get reclassify submissions", "description": "Returns information for submissions made to reclassify emails. Shows the status, outcome, and disposition changes for reclassification requests made by users or the security team. Useful for tracking false positive/negative reports.", "parameters": [{"name": "start", "in": "query", "description": "The beginning of the search date range. Defaults to `now - 30 days`.", "schema": {"type": "string", "format": "date-time"}, "example": "2022-06-25T14:30:00Z"}, {"name": "end", "in": "query", "description": "The end of the search date range. Defaults to `now`.", "schema": {"type": "string", "format": "date-time"}, "example": "2022-07-25T14:30:00Z"}, {"name": "type", "in": "query", "schema": {"type": "string", "enum": ["TEAM", "USER"]}}, {"name": "submission_id", "in": "query", "schema": {"type": "string"}}, {"name": "original_disposition", "in": "query", "schema": {"$ref": "#/components/schemas/email-security_SubmissionDisposition"}}, {"name": "requested_disposition", "in": "query", "schema": {"$ref": "#/components/schemas/email-security_SubmissionDisposition"}}, {"name": "outcome_disposition", "in": "query", "schema": {"$ref": "#/components/schemas/email-security_SubmissionDisposition"}}, {"name": "status", "in": "query", "schema": {"type": "string"}}, {"name": "query", "in": "query", "schema": {"type": "string", "nullable": true}}, {"name": "escalated_from_user", "in": "query", "description": "When true, return only submissions that were escalated by an end user (vs. by the security team). When false, return only submissions that were not escalated by an end user. When omitted, no filter is applied.", "schema": {"type": "boolean"}}, {"$ref": "#/components/parameters/email-security_page"}, {"$ref": "#/components/parameters/email-security_per_page"}], "responses": {"200": {"description": "List of submissions.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_Submission"}}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security"], "x-api-token-group": null, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.submissions", "x-fern-sdk-method-name": "list", "x-forge-hidden": true, "x-stability": "beta"}
```
