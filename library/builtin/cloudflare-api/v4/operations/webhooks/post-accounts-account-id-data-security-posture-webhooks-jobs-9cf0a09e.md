---
title: Create webhook jobs
page_id: operation-post-accounts-account-id-data-security-posture-webhooks-jobs-81a2a9c7
path: operations/webhooks
description: Creates webhook jobs to send a finding instance to one or more configured webhooks.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/data-security/posture/webhooks/jobs
operation_ids:
    - CreateWebhookJobs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create webhook jobs

`POST /accounts/{account_id}/data-security/posture/webhooks/jobs`

Operation ID: `CreateWebhookJobs`

Creates webhook jobs to send a finding instance to one or more configured webhooks.

## Definition

```yaml
{"operationId": "CreateWebhookJobs", "summary": "Create webhook jobs", "description": "Creates webhook jobs to send a finding instance to one or more configured webhooks.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_CreateWebhookJobsRequest"}}}}, "responses": {"200": {"description": "OK: Webhook jobs successfully created", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_CreateWebhookJobsResponse"}}}}, "400": {"description": "Bad Request: Invalid request parameters", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "401": {"description": "Unauthorized: Authentication required", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "403": {"description": "Forbidden: Insufficient permissions", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "404": {"description": "Not Found: Webhook or finding instance not found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}}, "security": [{"api_token": []}], "tags": ["webhooks"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
