---
title: Generate takedown letter
page_id: operation-post-accounts-account-id-cloudforce-one-v2-brand-protection-letter-gener-8c8f4270
path: operations/brand-protection
description: Generate a takedown letter from a template. Returns V4 JSON for text format or a PDF binary for pdf format.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/letter/generate
operation_ids:
    - post_LetterGenerate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Generate takedown letter

`POST /accounts/{account_id}/cloudforce-one/v2/brand-protection/letter/generate`

Operation ID: `post_LetterGenerate`

Generate a takedown letter from a template. Returns V4 JSON for text format or a PDF binary for pdf format.

## Definition

```yaml
{"operationId": "post_LetterGenerate", "summary": "Generate takedown letter", "description": "Generate a takedown letter from a template. Returns V4 JSON for text format or a PDF binary for pdf format.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"createNotice": {"type": "boolean", "default": false}, "fields": {"type": "object", "properties": {"discoveryDate": {"type": "string", "pattern": "^\\d{4}-\\d{2}-\\d{2}$"}, "domain": {"type": "string", "maxLength": 500}, "generationDate": {"type": "string", "pattern": "^\\d{4}-\\d{2}-\\d{2}$"}, "jurisdiction": {"type": "string", "maxLength": 500}, "registrantEmail": {"type": "string", "maxLength": 500}, "registrar": {"type": "string", "maxLength": 500}, "registrarEmail": {"type": "string", "maxLength": 500}, "resolutionByDate": {"type": "string", "pattern": "^\\d{4}-\\d{2}-\\d{2}$"}, "senderCompany": {"type": "string", "maxLength": 500}, "senderEmail": {"type": "string", "maxLength": 500}, "senderName": {"type": "string", "maxLength": 500}, "senderTitle": {"type": "string", "maxLength": 500}, "trademarkName": {"type": "string", "maxLength": 500}, "trademarkNumber": {"type": "string", "maxLength": 500}, "trademarkOwner": {"type": "string", "maxLength": 500}}}, "format": {"type": "string", "default": "text", "enum": ["text", "pdf"]}, "noticeParams": {"type": "object", "properties": {"domain": {"type": "string", "maxLength": 253, "minLength": 1}, "queryId": {"type": "integer", "exclusiveMinimum": true, "minimum": 0}, "status": {"type": "string", "default": "sent", "enum": ["draft", "sent", "resolved", "expired"]}}, "required": ["domain"]}, "templateId": {"type": "string", "minLength": 1}}, "required": ["templateId", "fields"]}}}}, "responses": {"200": {"description": "Letter generated successfully. When format=text, returns V4 JSON with { result: { letter: string } }. When format=pdf, returns binary with Content-Type: application/pdf.", "content": {"application/json": {"schema": {"type": "object", "properties": {"letter": {"type": "string"}, "letterId": {"type": "integer", "exclusiveMinimum": true, "minimum": 0}, "notice": {"type": "object", "properties": {"createdAt": {"type": "string", "nullable": true}, "domain": {"type": "string"}, "id": {"type": "number"}, "matchId": {"type": "number", "nullable": true}, "matchType": {"type": "string", "enum": ["logo", "domain"], "nullable": true}, "queryId": {"type": "number", "nullable": true}, "status": {"type": "string", "enum": ["draft", "sent", "resolved", "expired"]}, "updatedAt": {"type": "string", "nullable": true}}, "required": ["id", "domain", "queryId", "status", "matchId", "matchType", "createdAt", "updatedAt"]}}, "required": ["letter"]}}, "application/pdf": {"schema": {"type": "string", "format": "binary"}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write"]}
```
