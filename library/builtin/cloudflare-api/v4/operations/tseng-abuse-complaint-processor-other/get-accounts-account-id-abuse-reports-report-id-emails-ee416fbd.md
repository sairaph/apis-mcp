---
title: List abuse report emails
page_id: operation-get-accounts-account-id-abuse-reports-report-id-emails-2036b1fd
path: operations/tseng-abuse-complaint-processor-other
description: List emails sent to the customer for an abuse report. Returns all successful customer emails sent for the specified abuse report. Does not include emails sent to hosts or submitters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/abuse-reports/{report_id}/emails
operation_ids:
    - ListEmails
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List abuse report emails

`GET /accounts/{account_id}/abuse-reports/{report_id}/emails`

Operation ID: `ListEmails`

List emails sent to the customer for an abuse report. Returns all successful customer emails sent for the specified abuse report. Does not include emails sent to hosts or submitters.

## Definition

```yaml
{"operationId": "ListEmails", "summary": "List abuse report emails", "description": "List emails sent to the customer for an abuse report. Returns all successful customer emails sent for the specified abuse report. Does not include emails sent to hosts or submitters.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare Account ID", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32}}, {"name": "report_id", "in": "path", "description": "Abuse Report ID", "required": true, "schema": {"type": "string"}}, {"name": "page", "in": "query", "description": "Page number to retrieve (default 1)", "schema": {"type": "integer"}}, {"name": "per_page", "in": "query", "description": "Number of emails per page (default 20, max 100)", "schema": {"type": "integer"}}], "responses": {"200": {"description": "List abuse report emails successful", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "result": {"type": "object", "properties": {"emails": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_EmailListItem"}}}, "required": ["emails"]}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}, "total_pages": {"type": "number"}}, "required": ["page", "per_page", "count", "total_count", "total_pages"]}, "success": {"type": "boolean"}}, "required": ["success"]}}}}, "400": {"description": "Bad request - invalid parameters or report not found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "success": {"type": "boolean"}}, "required": ["success"]}}}}, "401": {"description": "The request is not authorized. Missing or invalid credentials may be\nrejected with the standard Cloudflare API authentication error. If\ncredentials are valid, the API token does not have the Abuse Reports\nread permission required for this account.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/abuse-reports_UnauthorizedErrorResponse"}}}}, "500": {"description": "Failed to list abuse report emails", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "success": {"type": "boolean"}}, "required": ["success"]}}}}}, "security": [{"api_token": []}], "tags": ["tseng-abuse-complaint-processor_other"], "x-api-token-group": ["Trust and Safety Write", "Trust and Safety Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "abuse-reports.emails", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
