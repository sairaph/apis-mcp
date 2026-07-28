---
title: List abuse reports
page_id: operation-get-accounts-account-id-abuse-reports-722f7d96
path: operations/tseng-abuse-complaint-processor-other
description: List the abuse reports for a given account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/abuse-reports
operation_ids:
    - ListAbuseReports
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List abuse reports

`GET /accounts/{account_id}/abuse-reports`

Operation ID: `ListAbuseReports`

List the abuse reports for a given account

## Definition

```yaml
{"operationId": "ListAbuseReports", "summary": "List abuse reports", "description": "List the abuse reports for a given account", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare Account ID", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32}}, {"name": "page", "in": "query", "description": "Where in pagination to start listing abuse reports", "schema": {"type": "integer"}}, {"name": "per_page", "in": "query", "description": "How many abuse reports per page to list", "schema": {"type": "integer"}}, {"name": "sort", "in": "query", "description": "A property to sort by, followed by the order (id, cdate, domain, type, status)", "schema": {"type": "string"}, "example": "cdate,desc"}, {"name": "domain", "in": "query", "description": "Filter by domain name related to the abuse report", "schema": {"type": "string"}, "example": "example.com"}, {"name": "created_before", "in": "query", "description": "Returns reports created before the specified date", "schema": {"description": "Time in RFC 3339 format (https://www.rfc-editor.org/rfc/rfc3339.html)", "type": "string", "example": "2009-11-10T23:00:00Z"}, "example": "2009-11-10T23:00:00Z"}, {"name": "created_after", "in": "query", "description": "Returns reports created after the specified date", "schema": {"description": "Time in RFC 3339 format (https://www.rfc-editor.org/rfc/rfc3339.html)", "type": "string", "example": "2009-11-10T23:00:00Z"}, "example": "2009-11-10T23:00:00Z"}, {"name": "status", "in": "query", "description": "Filter by the status of the report.", "schema": {"$ref": "#/components/schemas/abuse-reports_ReportStatus"}, "example": "denied"}, {"name": "type", "in": "query", "description": "Filter by the type of the report.", "schema": {"$ref": "#/components/schemas/abuse-reports_ReportType"}, "example": "denied"}, {"name": "mitigation_status", "in": "query", "description": "Filter reports that have any mitigations in the given status.", "schema": {"$ref": "#/components/schemas/abuse-reports_MitigationStatus"}, "example": "active"}], "responses": {"200": {"description": "Abuse report list successful", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "result": {"type": "object", "properties": {"reports": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_AbuseReport"}}}, "required": ["reports"]}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}, "total_pages": {"type": "number"}}, "required": ["page", "per_page", "count", "total_count", "total_pages"]}, "success": {"type": "boolean"}}, "required": ["success"]}}}}, "401": {"description": "The request is not authorized. Missing or invalid credentials may be\nrejected with the standard Cloudflare API authentication error. If\ncredentials are valid, the API token does not have the Abuse Reports\nread permission required for this account.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/abuse-reports_UnauthorizedErrorResponse"}}}}, "500": {"description": "Failed to retrieve abuse reports", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "success": {"type": "boolean"}}, "required": ["success"]}}}}}, "security": [{"api_token": []}], "tags": ["tseng-abuse-complaint-processor_other"], "x-api-token-group": ["Trust and Safety Write", "Trust and Safety Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "abuse-reports", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
