---
title: List Access SCIM update logs
page_id: operation-get-accounts-account-id-access-logs-scim-updates-f44f5807
path: operations/access-scim-update-logs
description: Lists Access SCIM update logs that maintain a record of updates made to User and Group resources synced to Cloudflare via the System for Cross-domain Identity Management (SCIM).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/logs/scim/updates
operation_ids:
    - access-scim-update-logs-list-access-scim-update-logs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Access SCIM update logs

`GET /accounts/{account_id}/access/logs/scim/updates`

Operation ID: `access-scim-update-logs-list-access-scim-update-logs`

Lists Access SCIM update logs that maintain a record of updates made to User and Group resources synced to Cloudflare via the System for Cross-domain Identity Management (SCIM).

## Definition

```yaml
{"operationId": "access-scim-update-logs-list-access-scim-update-logs", "summary": "List Access SCIM update logs", "description": "Lists Access SCIM update logs that maintain a record of updates made to User and Group resources synced to Cloudflare via the System for Cross-domain Identity Management (SCIM).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "limit", "in": "query", "schema": {"$ref": "#/components/schemas/access_limit"}}, {"name": "direction", "in": "query", "schema": {"$ref": "#/components/schemas/access_direction"}, "example": "desc"}, {"name": "since", "in": "query", "schema": {"$ref": "#/components/schemas/access_since"}}, {"name": "until", "in": "query", "schema": {"$ref": "#/components/schemas/access_until"}}, {"name": "idp_id", "in": "query", "required": true, "schema": {"$ref": "#/components/schemas/access_idp_id"}, "explode": true, "style": "form"}, {"name": "status", "in": "query", "schema": {"$ref": "#/components/schemas/access_status-2"}, "explode": true, "style": "form"}, {"name": "resource_type", "in": "query", "schema": {"$ref": "#/components/schemas/access_resource_type"}, "explode": true, "style": "form"}, {"name": "request_method", "in": "query", "schema": {"$ref": "#/components/schemas/access_request_method"}, "explode": true, "style": "form"}, {"name": "resource_user_email", "in": "query", "schema": {"$ref": "#/components/schemas/access_resource_user_email"}, "explode": true, "style": "form"}, {"name": "resource_group_name", "in": "query", "schema": {"$ref": "#/components/schemas/access_resource_group_name"}, "explode": true, "style": "form"}, {"name": "cf_resource_id", "in": "query", "schema": {"$ref": "#/components/schemas/access_cf_resource_id-3"}, "explode": true, "style": "form"}, {"name": "idp_resource_id", "in": "query", "schema": {"$ref": "#/components/schemas/access_idp_resource_id-3"}, "explode": true, "style": "form"}, {"$ref": "#/components/parameters/access_page"}, {"name": "per_page", "in": "query", "schema": {"description": "Number of results per page.", "type": "integer", "default": 20, "maximum": 1000}}], "responses": {"200": {"description": "Get Access SCIM update logs response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_scim_update_logs_response"}}}}, "4XX": {"description": "Get Access SCIM update logs response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Access SCIM update logs"], "x-api-token-group": ["Access: SCIM Logs Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.logs.scim.updates", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
