---
title: Download takedown letter PDF
page_id: operation-get-accounts-account-id-cloudforce-one-v2-brand-protection-takedown-noti-4410ccb2
path: operations/brand-protection
description: Download the PDF for a stored takedown letter. If no PDF exists yet, it is generated from the stored letter text and cached for future downloads.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices/{notice_id}/letters/{letter_id}/pdf
operation_ids:
    - get_TakedownLetterPdfGet
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Download takedown letter PDF

`GET /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices/{notice_id}/letters/{letter_id}/pdf`

Operation ID: `get_TakedownLetterPdfGet`

Download the PDF for a stored takedown letter. If no PDF exists yet, it is generated from the stored letter text and cached for future downloads.

## Definition

```yaml
{"operationId": "get_TakedownLetterPdfGet", "summary": "Download takedown letter PDF", "description": "Download the PDF for a stored takedown letter. If no PDF exists yet, it is generated from the stored letter text and cached for future downloads.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "notice_id", "in": "path", "required": true, "schema": {"type": "integer", "exclusiveMinimum": true, "minimum": 0}}, {"name": "letter_id", "in": "path", "required": true, "schema": {"type": "integer", "exclusiveMinimum": true, "minimum": 0}}], "responses": {"200": {"description": "Takedown letter PDF downloaded successfully", "content": {"application/pdf": {"schema": {"type": "string", "format": "binary"}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
