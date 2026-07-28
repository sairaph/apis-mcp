---
title: v2.error.document_purpose_invalid
page_id: schema-v2-error-document-purpose-invalid-0fb39d2b
path: schemas
description: Provided file tokens for documents are of the wrong purpose.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.document_purpose_invalid

Provided file tokens for documents are of the wrong purpose.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["document_purpose_invalid"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Provided file tokens for documents are of the wrong purpose."}
```
