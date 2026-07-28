---
title: Initialize new SSO connector
page_id: operation-post-accounts-account-id-sso-connectors-6544d145
path: operations/sso
description: Creates a new SSO connector for logging into Cloudflare through an identity provider.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/sso_connectors
operation_ids:
    - init-new-sso-connector
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Initialize new SSO connector

`POST /accounts/{account_id}/sso_connectors`

Operation ID: `init-new-sso-connector`

Creates a new SSO connector for logging into Cloudflare through an identity provider.

## Definition

```yaml
{"operationId": "init-new-sso-connector", "summary": "Initialize new SSO connector", "description": "Creates a new SSO connector for logging into Cloudflare through an identity provider.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"begin_verification": {"description": "Begin the verification process after creation", "type": "boolean", "example": true, "default": true, "x-stainless-terraform-configurability": "computed_optional"}, "email_domain": {"description": "Email domain of the new SSO connector", "type": "string", "example": "example.com"}, "use_fedramp_language": {"$ref": "#/components/schemas/iam_use_fedramp_language"}}, "required": ["email_domain"]}}}}, "responses": {"200": {"description": "Initialize new SSO connector response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_sso_connector_response"}}}}, "4XX": {"description": "Initialize new SSO connector response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["SSO"], "x-api-token-group": ["SSO Connector Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
