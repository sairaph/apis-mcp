---
title: List applications
page_id: operation-get-accounts-account-id-one-applications-1e5da2ee
path: operations/applications
description: Returns a list of available applications with use cases and permissions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/one/applications
operation_ids:
    - list_applications_v2
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List applications

`GET /accounts/{account_id}/one/applications`

Operation ID: `list_applications_v2`

Returns a list of available applications with use cases and permissions.

## Definition

```yaml
{"operationId": "list_applications_v2", "summary": "List applications", "description": "Returns a list of available applications with use cases and permissions.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account identifier.", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353"}}, {"name": "environment", "in": "query", "description": "Filter by supported environment (standard, fedramp).", "schema": {"type": "string"}}], "responses": {"200": {"description": "List of applications.", "content": {"application/json": {"examples": {"SampleSuccessfulReturn": {"summary": "Example application list response", "value": [{"errors": [], "messages": [], "result": [{"auth_methods": [{"display_name": "OAuth 2.0 Admin Consent", "id": "oauth2_standard"}], "category": "Productivity", "description": "Monitor OneDrive, SharePoint, Teams, and Outlook.", "display_name": "Microsoft", "dlp_enabled": true, "id": "microsoft_internal", "logo": "/api/v4/accounts/12345678/casb/static/microsoft_internal.svg", "permissions": [{"display_name": "Read all users' full profiles", "scope": "User.Read.All", "severity": "high"}, {"display_name": "Read all files", "scope": "Files.Read.All", "severity": "high"}, {"display_name": "Read and write mail", "scope": "Mail.ReadWrite", "severity": "critical"}], "supported_environments": ["standard", "fedramp"], "use_cases": [{"display_name": "Cloud Access Security Broker", "id": "casb"}, {"display_name": "Cloud Email Security", "id": "ces"}]}], "result_info": {"count": 1, "next": null, "page": 1, "per_page": 10, "previous": null, "total_count": 1}, "success": true}]}}, "schema": {"type": "array", "items": {"$ref": "#/components/schemas/one_ApplicationList"}}}}}, "400": {"description": "Invalid request."}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Applications"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero_trust.casb.applications", "x-fern-sdk-method-name": "list", "x-forge-hidden": true, "x-stability": "beta"}
```
