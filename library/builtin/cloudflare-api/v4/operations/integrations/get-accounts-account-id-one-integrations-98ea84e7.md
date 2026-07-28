---
title: List integrations
page_id: operation-get-accounts-account-id-one-integrations-afe12cf1
path: operations/integrations
description: Returns a paginated list of integrations for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/one/integrations
operation_ids:
    - list_integrations_v2
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List integrations

`GET /accounts/{account_id}/one/integrations`

Operation ID: `list_integrations_v2`

Returns a paginated list of integrations for the account.

## Definition

```yaml
{"operationId": "list_integrations_v2", "summary": "List integrations", "description": "Returns a paginated list of integrations for the account.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account identifier.", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353"}}, {"name": "application", "in": "query", "description": "Filter by application/vendor (e.g., GOOGLE_WORKSPACE, MICROSOFT_INTERNAL).", "schema": {"type": "string"}}, {"name": "direction", "in": "query", "description": "Direction to order results.", "schema": {"type": "string", "enum": ["asc", "desc"]}}, {"name": "dlp_enabled", "in": "query", "description": "Filter by DLP enabled status (true/false).", "schema": {"type": "boolean"}}, {"name": "order", "in": "query", "description": "Field to order results by.", "schema": {"type": "string", "enum": ["application", "created", "name", "status"]}}, {"name": "page", "in": "query", "description": "Page number within the paginated result set.", "schema": {"type": "integer"}}, {"name": "page_size", "in": "query", "description": "Number of results per page.", "schema": {"type": "integer"}}, {"name": "search", "in": "query", "description": "Search integrations by name or application.", "schema": {"type": "string"}}, {"name": "status", "in": "query", "description": "Filter by integration status.", "schema": {"type": "string", "enum": ["Healthy", "Initializing", "Offline", "Unhealthy"]}}, {"name": "use_cases", "in": "query", "description": "Filter by enabled use cases (e.g., casb, ces). Matches integrations enrolled in any of the specified values. Can be specified multiple times.", "schema": {"type": "string"}}], "responses": {"200": {"description": "List of integrations.", "content": {"application/json": {"examples": {"SampleSuccessfulReturn": {"summary": "Example integration list response", "value": {"errors": [], "messages": [], "result": [{"application": {"category": "Productivity", "display_name": "Google Workspace", "logo": "https://onprem.cloudflare.com/static/google_workspace.png"}, "created": "2025-01-15T10:00:00Z", "id": "019d2e6a-d995-7185-afbd-4feead9e42ec", "is_paused": false, "name": "My Google Workspace", "status": "Healthy", "updated": "2025-04-10T08:30:00Z"}], "result_info": {"count": 1, "next": null, "page": 1, "per_page": 10, "previous": null, "total_count": 1}, "success": true}}}, "schema": {"$ref": "#/components/schemas/one_PaginatedIntegrationV2BaseList"}}}}, "400": {"description": "Invalid request."}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Integrations"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero_trust.casb.integrations", "x-fern-sdk-method-name": "list", "x-forge-hidden": true, "x-stability": "beta"}
```
