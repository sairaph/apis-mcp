---
title: Get application details
page_id: operation-get-accounts-account-id-one-applications-application-id-0c8f4c49
path: operations/applications
description: Returns full application details including auth methods, use cases, and permissions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/one/applications/{application_id}
operation_ids:
    - get_application_v2
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get application details

`GET /accounts/{account_id}/one/applications/{application_id}`

Operation ID: `get_application_v2`

Returns full application details including auth methods, use cases, and permissions.

## Definition

```yaml
{"operationId": "get_application_v2", "summary": "Get application details", "description": "Returns full application details including auth methods, use cases, and permissions.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account identifier.", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353"}}, {"name": "application_id", "in": "path", "description": "Application/vendor identifier.", "required": true, "schema": {"type": "string", "enum": ["ANTHROPIC", "BITBUCKET", "BOX", "CONFLUENCE", "DROPBOX", "GITHUB", "GOOGLE_CLOUD_PLATFORM", "GOOGLE_WORKSPACE", "JIRA", "MICROSOFT_INTERNAL", "OPENAI", "SALESFORCE", "SLACK"]}}], "responses": {"200": {"description": "Application details.", "content": {"application/json": {"examples": {"Microsoft365": {"summary": "Microsoft 365 application detail", "value": {"auth_methods": [{"display_name": "OAuth 2.0 Admin Consent", "id": "oauth2", "is_default": true, "supported_environments": ["standard", "fedramp"]}], "category": "Productivity", "description": "Monitor OneDrive, SharePoint, Teams, and Outlook.", "display_name": "Microsoft", "dlp_enabled": true, "id": "microsoft_internal", "instructions": "You'll need a Microsoft 365 admin account with Global Admin or Application Admin role.", "logo": "/api/v4/accounts/12345678/casb/static/microsoft_internal.svg", "use_cases": [{"base_scopes": [{"display_name": "Read all users' full profiles", "scope": "User.Read.All", "severity": "high"}, {"display_name": "Read all files", "scope": "Files.Read.All", "severity": "high"}], "description": "Discover and secure SaaS applications", "display_name": "Cloud Access Security Broker", "features": [{"description": "Automatically remediate security issues", "display_name": "Auto Remediation", "id": "auto_remediation", "scopes": [{"display_name": "Read and write all files", "scope": "Files.ReadWrite.All", "severity": "critical"}]}], "id": "casb"}]}}}, "schema": {"$ref": "#/components/schemas/one_ApplicationDetail"}}}}, "404": {"description": "Application not found."}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Applications"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero_trust.casb.applications", "x-fern-sdk-method-name": "get", "x-forge-hidden": true, "x-stability": "beta"}
```
