---
title: Get organization accounts
page_id: operation-get-organizations-organization-id-accounts-ccd706ce
path: operations/organizations
description: Retrieve a list of accounts that belong to a specific organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /organizations/{organization_id}/accounts
operation_ids:
    - Organizations_getAccounts
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get organization accounts

`GET /organizations/{organization_id}/accounts`

Operation ID: `Organizations_getAccounts`

Retrieve a list of accounts that belong to a specific organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)

## Definition

```yaml
{"operationId": "Organizations_getAccounts", "summary": "Get organization accounts", "description": "Retrieve a list of accounts that belong to a specific organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)", "parameters": [{"name": "organization_id", "in": "path", "description": "The ID of the organization to retrieve a list of accounts for.", "required": true, "schema": {"$ref": "#/components/schemas/organizations-api_OrganizationID"}}, {"name": "account_pubname", "in": "query", "description": "(case-insensitive) Filter the list of accounts to where the account_pubname is equal to\na particular string.", "schema": {"type": "string"}, "deprecated": true, "explode": false}, {"name": "account_pubname.startsWith", "in": "query", "description": "(case-insensitive) Filter the list of accounts to where the account_pubname starts with\na particular string.", "schema": {"type": "string"}, "deprecated": true, "explode": false}, {"name": "account_pubname.endsWith", "in": "query", "description": "(case-insensitive) Filter the list of accounts to where the account_pubname ends with\na particular string.", "schema": {"type": "string"}, "deprecated": true, "explode": false}, {"name": "account_pubname.contains", "in": "query", "description": "(case-insensitive) Filter the list of accounts to where the account_pubname contains\na particular string.", "schema": {"type": "string"}, "deprecated": true, "explode": false}, {"name": "name", "in": "query", "description": "(case-insensitive) Filter the list of accounts to where the name is equal to a\nparticular string.", "schema": {"type": "string"}, "explode": false}, {"name": "name.startsWith", "in": "query", "description": "(case-insensitive) Filter the list of accounts to where the name starts with a\nparticular string.", "schema": {"type": "string"}, "explode": false}, {"name": "name.endsWith", "in": "query", "description": "(case-insensitive) Filter the list of accounts to where the name ends with a particular\nstring.", "schema": {"type": "string"}, "explode": false}, {"name": "name.contains", "in": "query", "description": "(case-insensitive) Filter the list of accounts to where the name contains a particular\nstring.", "schema": {"type": "string"}, "explode": false}, {"name": "order_by", "in": "query", "description": "Field to order results by. Currently supported values: `account_name`.\nWhen not specified, results are ordered by internal account ID.", "schema": {"type": "string", "enum": ["account_name"]}, "explode": false}, {"name": "direction", "in": "query", "description": "Sort direction for the order_by field. Valid values: `asc`, `desc`.\nDefaults to `asc` when order_by is specified.", "schema": {"type": "string", "enum": ["asc", "desc"]}, "explode": false}, {"$ref": "#/components/parameters/organizations-api_PageTokenParamsPageToken"}, {"$ref": "#/components/parameters/organizations-api_PageTokenParamsPageSize"}], "responses": {"200": {"description": "The request has succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_Account"}}, "result_info": {"$ref": "#/components/schemas/organizations-api_PageTokenResultInfo"}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result", "result_info"]}}}}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Organizations"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "organizations.organization-accounts", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
