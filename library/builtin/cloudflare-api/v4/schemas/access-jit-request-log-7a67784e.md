---
title: access_jit_request_log
page_id: schema-access-jit-request-log-7a67784e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_jit_request_log

```yaml
{"type": "object", "properties": {"app_aud": {"description": "Audience of the Access application.", "type": "string"}, "app_hostname": {"type": "string"}, "approvals_received": {"description": "Number of unique actors who approved the request.", "type": "integer"}, "approver_emails": {"description": "Unique actors who approved the request.", "type": "array", "items": {"format": "email", "type": "string"}}, "created_at": {"$ref": "#/components/schemas/access_timestamp"}, "expires_at": {"$ref": "#/components/schemas/access_timestamp"}, "knock_request_id": {"description": "Unique identifier for the JIT request.", "type": "string"}, "purpose_justification": {"type": "string"}, "requester_email": {"type": "string", "format": "email"}, "status": {"$ref": "#/components/schemas/access_jit_request_status"}}}
```
