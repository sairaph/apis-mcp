---
title: Resume integration
page_id: operation-post-accounts-account-id-one-integrations-id-resume-ffeab92c
path: operations/integrations
description: Resumes a paused integration, restarting crawlers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/one/integrations/{id}/resume
operation_ids:
    - resume_integration_v2
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Resume integration

`POST /accounts/{account_id}/one/integrations/{id}/resume`

Operation ID: `resume_integration_v2`

Resumes a paused integration, restarting crawlers.

## Definition

```yaml
{"operationId": "resume_integration_v2", "summary": "Resume integration", "description": "Resumes a paused integration, restarting crawlers.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account identifier.", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353"}}, {"name": "id", "in": "path", "description": "Integration ID.", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Integration resumed successfully.", "content": {"application/json": {"examples": {"SampleSuccessfulReturn": {"summary": "Example integration resume response", "value": {"errors": [], "messages": [], "result": {"application": {"category": "Productivity", "display_name": "Google Workspace", "logo": "https://onprem.cloudflare.com/static/google_workspace.png"}, "auth_method": {"display_name": "OAuth 2.0", "id": "oauth"}, "authorization_link": {"components": {"client_id": "abc", "instance_name": "example"}, "link": "https://example.cloudflare.com/authorize"}, "created": "2025-01-01T00:00:00Z", "credentials_expiry": "2026-01-01T00:00:00Z", "dlp_profiles": ["e91a2360-da51-4fdf-9711-bcdecd462614"], "health_details": [], "id": "019d2e6a-d995-7185-afbd-4feead9e42ec", "is_paused": false, "last_hydrated": "2025-04-10T08:30:00Z", "name": "My Google Workspace", "organization_id": 1, "status": "Healthy", "updated": "2025-04-10T08:30:00Z", "use_cases": [{"description": "Discover and secure SaaS applications", "features": [{"description": "Automatically remediate security issues (requires write permissions)", "id": "auto_remediation", "is_enabled": true, "name": "Auto Remediation", "permissions": [{"display_name": "Manage users", "scope": "https://www.googleapis.com/auth/admin.directory.user", "status": "granted"}]}], "id": "casb", "is_enabled": true, "name": "Cloud Access Security Broker", "permissions": [{"display_name": "Drive (Read Only)", "scope": "https://www.googleapis.com/auth/drive.readonly", "status": "granted"}, {"display_name": "Gmail (Read Only)", "scope": "https://www.googleapis.com/auth/gmail.readonly", "status": "missing"}]}, {"description": "Protect against email-based threats", "features": [], "id": "ces", "is_enabled": false, "name": "Cloud Email Security", "permissions": []}]}, "result_info": {"count": 1, "next": null, "page": 1, "per_page": 1, "previous": null, "total_count": 1}, "success": true}}}, "schema": {"$ref": "#/components/schemas/one_IntegrationV2Detail"}}}}, "400": {"description": "Invalid request."}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Integrations"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero_trust.casb.integrations", "x-fern-sdk-method-name": "resume", "x-forge-hidden": true, "x-stability": "beta"}
```
