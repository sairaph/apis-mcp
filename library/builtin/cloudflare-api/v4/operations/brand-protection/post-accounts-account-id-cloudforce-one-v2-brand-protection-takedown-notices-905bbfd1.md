---
title: Create takedown notice
page_id: operation-post-accounts-account-id-cloudforce-one-v2-brand-protection-takedown-not-b4183a4a
path: operations/brand-protection
description: Create a new takedown notice for a domain suspected of trademark infringement.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices
operation_ids:
    - post_TakedownNoticeCreate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create takedown notice

`POST /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices`

Operation ID: `post_TakedownNoticeCreate`

Create a new takedown notice for a domain suspected of trademark infringement.

## Definition

```yaml
{"operationId": "post_TakedownNoticeCreate", "summary": "Create takedown notice", "description": "Create a new takedown notice for a domain suspected of trademark infringement.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"domain": {"type": "string", "maxLength": 253, "minLength": 1}, "matchId": {"type": "integer", "exclusiveMinimum": true, "minimum": 0}, "matchType": {"type": "string", "enum": ["logo", "domain"]}, "queryId": {"type": "integer", "exclusiveMinimum": true, "minimum": 0}, "status": {"type": "string", "default": "draft", "enum": ["draft", "sent", "resolved", "expired"]}}, "required": ["domain"]}}}}, "responses": {"200": {"description": "Takedown notice created successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"createdAt": {"type": "string", "nullable": true}, "domain": {"type": "string"}, "id": {"type": "number"}, "matchId": {"type": "number", "nullable": true}, "matchType": {"type": "string", "enum": ["logo", "domain"], "nullable": true}, "queryId": {"type": "number", "nullable": true}, "status": {"type": "string", "enum": ["draft", "sent", "resolved", "expired"]}, "updatedAt": {"type": "string", "nullable": true}}, "required": ["id", "domain", "queryId", "status", "matchId", "matchType", "createdAt", "updatedAt"]}}}}, "409": {"description": "A takedown notice already exists for this domain"}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write"]}
```
