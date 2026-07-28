---
title: Check whether a report can be appealed
page_id: operation-get-accounts-account-id-abuse-reports-report-id-appeals-eligibility-6508af24
path: operations/tseng-abuse-complaint-processor-other
description: 'Returns whether the report is currently appealable, along with the signals behind that decision: whether it already has an open appeal, how many appeals have been submitted against it, and whether it has at least one mitigation that an appeal could reverse.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/abuse-reports/{report_id}/appeals/eligibility
operation_ids:
    - CheckAppealEligibility
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Check whether a report can be appealed

`GET /accounts/{account_id}/abuse-reports/{report_id}/appeals/eligibility`

Operation ID: `CheckAppealEligibility`

Returns whether the report is currently appealable, along with the signals behind that decision: whether it already has an open appeal, how many appeals have been submitted against it, and whether it has at least one mitigation that an appeal could reverse.

## Definition

```yaml
{"operationId": "CheckAppealEligibility", "summary": "Check whether a report can be appealed", "description": "Returns whether the report is currently appealable, along with the signals behind that decision: whether it already has an open appeal, how many appeals have been submitted against it, and whether it has at least one mitigation that an appeal could reverse.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare Account ID", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32}}, {"name": "report_id", "in": "path", "description": "Abuse Report ID", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Appeal eligibility for the report", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "result": {"$ref": "#/components/schemas/abuse-reports_AppealEligibility"}, "success": {"type": "boolean"}}, "required": ["success"]}}}}, "401": {"description": "The request is not authorized. Missing or invalid credentials may be\nrejected with the standard Cloudflare API authentication error. If\ncredentials are valid, the API token does not have the Abuse Reports\nread permission required for this account.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/abuse-reports_UnauthorizedErrorResponse"}}}}, "404": {"description": "The report was not found, or does not support appeals.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "success": {"type": "boolean"}}, "required": ["success"]}}}}, "500": {"description": "Failed to check appeal eligibility.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "success": {"type": "boolean"}}, "required": ["success"]}}}}}, "security": [{"api_token": []}], "tags": ["tseng-abuse-complaint-processor_other"], "x-api-token-group": ["Trust and Safety Write", "Trust and Safety Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "abuse-reports.appeals", "x-fern-sdk-method-name": "eligibility", "x-forge-hidden": true}
```
