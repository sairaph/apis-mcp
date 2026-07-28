---
title: Get application
page_id: operation-get-accounts-account-id-resource-library-applications-id-75f58cde
path: operations/applications
description: Get application by ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/resource-library/applications/{id}
operation_ids:
    - getApplicationById
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get application

`GET /accounts/{account_id}/resource-library/applications/{id}`

Operation ID: `getApplicationById`

Get application by ID.

## Definition

```yaml
{"operationId": "getApplicationById", "summary": "Get application", "description": "Get application by ID.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"type": "string"}, "example": "023e105f4ecef8ad9ca31a8372d0c353"}, {"name": "id", "in": "path", "description": "Application ID.", "required": true, "schema": {"type": "string"}, "example": "0b63249c-95bf-4cc0-a7cc-d7faaaf1dac0"}], "responses": {"200": {"description": "Get the application response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/alexandria_get_application_response"}}}}, "4XX": {"description": "Get application by id response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/alexandria_get_application_response"}, {"$ref": "#/components/schemas/alexandria_api_response_common_failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Applications"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "accounts.applications.get.by", "x-fern-sdk-method-name": "id"}
```
