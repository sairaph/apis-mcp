---
title: cloudforce-one-requests_request-edit
page_id: schema-cloudforce-one-requests-request-edit-5e2a3d34
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one-requests_request-edit

```yaml
{"type": "object", "properties": {"content": {"$ref": "#/components/schemas/cloudforce-one-requests_request-content"}, "priority": {"description": "Priority for analyzing the request.", "type": "string", "example": "routine", "x-auditable": true}, "request_type": {"$ref": "#/components/schemas/cloudforce-one-requests_request-type"}, "summary": {"$ref": "#/components/schemas/cloudforce-one-requests_request-summary"}, "tlp": {"$ref": "#/components/schemas/cloudforce-one-requests_tlp"}}, "title": "Request Editable Parameters"}
```
