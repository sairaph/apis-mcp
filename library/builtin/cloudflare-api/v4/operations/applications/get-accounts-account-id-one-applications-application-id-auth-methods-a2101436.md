---
title: Get auth methods
page_id: operation-get-accounts-account-id-one-applications-application-id-auth-methods-57a29f1d
path: operations/applications
description: Returns available auth methods for the specified vendor, including credential schema, instructions, and example payloads. Use this to understand what credentials are required before calling POST /v2/integrations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/one/applications/{application_id}/auth-methods
operation_ids:
    - get_application_auth_methods_v2
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get auth methods

`GET /accounts/{account_id}/one/applications/{application_id}/auth-methods`

Operation ID: `get_application_auth_methods_v2`

Returns available auth methods for the specified vendor, including credential schema, instructions, and example payloads. Use this to understand what credentials are required before calling POST /v2/integrations.

## Definition

```yaml
{"operationId": "get_application_auth_methods_v2", "summary": "Get auth methods", "description": "Returns available auth methods for the specified vendor, including credential schema, instructions, and example payloads. Use this to understand what credentials are required before calling POST /v2/integrations.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account identifier.", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353"}}, {"name": "application_id", "in": "path", "description": "Application/vendor identifier.", "required": true, "schema": {"type": "string", "enum": ["ANTHROPIC", "BITBUCKET", "BOX", "CONFLUENCE", "DROPBOX", "GITHUB", "GOOGLE_CLOUD_PLATFORM", "GOOGLE_WORKSPACE", "JIRA", "MICROSOFT_INTERNAL", "OPENAI", "SALESFORCE", "SLACK"]}}], "responses": {"200": {"description": "Auth methods available for this application.", "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/one_AuthMethodDetail"}}}}}, "404": {"description": "Application not found."}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Applications"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero_trust.casb.applications.auth_methods", "x-fern-sdk-method-name": "list", "x-forge-hidden": true, "x-stability": "beta"}
```
