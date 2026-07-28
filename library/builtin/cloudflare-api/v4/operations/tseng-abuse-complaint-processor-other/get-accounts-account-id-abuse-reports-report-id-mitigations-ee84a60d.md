---
title: List abuse report mitigations
page_id: operation-get-accounts-account-id-abuse-reports-report-id-mitigations-da0d1a1e
path: operations/tseng-abuse-complaint-processor-other
description: List mitigations done to remediate the abuse report.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/abuse-reports/{report_id}/mitigations
operation_ids:
    - ListMitigations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List abuse report mitigations

`GET /accounts/{account_id}/abuse-reports/{report_id}/mitigations`

Operation ID: `ListMitigations`

List mitigations done to remediate the abuse report.

## Definition

```yaml
{"operationId": "ListMitigations", "summary": "List abuse report mitigations", "description": "List mitigations done to remediate the abuse report.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare Account ID", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32}}, {"name": "report_id", "in": "path", "description": "Abuse Report ID", "required": true, "schema": {"type": "string"}}, {"name": "page", "in": "query", "description": "Where in pagination to start listing abuse reports", "schema": {"type": "integer"}}, {"name": "per_page", "in": "query", "description": "How many abuse reports per page to list", "schema": {"type": "integer"}}, {"name": "sort", "in": "query", "description": "A property to sort by, followed by the order", "schema": {"type": "string", "enum": ["type,asc", "type,desc", "effective_date,asc", "effective_date,desc", "status,asc", "status,desc", "entity_type,asc", "entity_type,desc"]}, "example": "type,desc"}, {"name": "type", "in": "query", "description": "Filter by the type of mitigation. This filter parameter can be specified multiple times to include multiple types of mitigations in the result set, e.g. ?type=rate_limit_cache&type=legal_block.", "schema": {"$ref": "#/components/schemas/abuse-reports_MitigationType"}}, {"name": "effective_before", "in": "query", "description": "Returns mitigations that were dispatched before the given date", "schema": {"description": "Time in RFC 3339 format (https://www.rfc-editor.org/rfc/rfc3339.html)", "type": "string", "example": "2009-11-10T23:00:00Z"}, "example": "2009-11-10T23:00:00Z"}, {"name": "effective_after", "in": "query", "description": "Returns mitigation that were dispatched after the given date", "schema": {"description": "Time in RFC 3339 format (https://www.rfc-editor.org/rfc/rfc3339.html)", "type": "string", "example": "2009-11-10T23:00:00Z"}, "example": "2009-11-10T23:00:00Z"}, {"name": "status", "in": "query", "description": "Filter by the status of the mitigation.", "schema": {"$ref": "#/components/schemas/abuse-reports_MitigationStatus"}}, {"name": "entity_type", "in": "query", "description": "Filter by the type of entity the mitigation impacts.", "schema": {"$ref": "#/components/schemas/abuse-reports_MitigatedEntityType"}}], "responses": {"200": {"description": "List abuse report mitigations successful", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "result": {"type": "object", "properties": {"mitigations": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_MitigationListItem"}}}, "required": ["mitigations"]}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}, "total_pages": {"type": "number"}}, "required": ["page", "per_page", "count", "total_count", "total_pages"]}, "success": {"type": "boolean"}}, "required": ["success"]}}}}, "401": {"description": "The request is not authorized. Missing or invalid credentials may be\nrejected with the standard Cloudflare API authentication error. If\ncredentials are valid, the API token does not have the Abuse Reports\nread permission required for this account.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/abuse-reports_UnauthorizedErrorResponse"}}}}, "500": {"description": "Failed to list abuse report mitigations", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "success": {"type": "boolean"}}, "required": ["success"]}}}}}, "security": [{"api_token": []}], "tags": ["tseng-abuse-complaint-processor_other"], "x-api-token-group": ["Trust and Safety Write", "Trust and Safety Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "abuse-reports.mitigations", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
