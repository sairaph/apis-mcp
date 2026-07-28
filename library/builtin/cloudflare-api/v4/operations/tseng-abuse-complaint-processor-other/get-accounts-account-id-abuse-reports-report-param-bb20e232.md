---
title: Abuse Report Details
page_id: operation-get-accounts-account-id-abuse-reports-report-param-f79e7650
path: operations/tseng-abuse-complaint-processor-other
description: Retrieve the details of an abuse report.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/abuse-reports/{report_param}
operation_ids:
    - GetAbuseReport
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Abuse Report Details

`GET /accounts/{account_id}/abuse-reports/{report_param}`

Operation ID: `GetAbuseReport`

Retrieve the details of an abuse report.

## Definition

```yaml
{"operationId": "GetAbuseReport", "summary": "Abuse Report Details", "description": "Retrieve the details of an abuse report.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare Account ID", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32}}, {"name": "report_param", "in": "path", "description": "Identifier of the abuse report", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Report submitted successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_ErrorMessage"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "result": {"$ref": "#/components/schemas/abuse-reports_AbuseReport"}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "Report submitted with an error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_ErrorMessage"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "result": {"$ref": "#/components/schemas/abuse-reports_AbuseReport"}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}, "401": {"description": "The request is not authorized. Missing or invalid credentials may be\nrejected with the standard Cloudflare API authentication error. If\ncredentials are valid, the API token does not have the Abuse Reports\nread permission required for this account.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/abuse-reports_UnauthorizedErrorResponse"}}}}, "500": {"description": "Report submitted with an error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_ErrorMessage"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "result": {"$ref": "#/components/schemas/abuse-reports_AbuseReport"}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["tseng-abuse-complaint-processor_other"], "x-api-token-group": ["Trust and Safety Write", "Trust and Safety Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "abuse-reports", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
