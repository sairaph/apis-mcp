---
title: workers_script-response-upload-single
page_id: schema-workers-script-response-upload-single-fb46d204
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_script-response-upload-single

```yaml
{"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_script-response-upload"}}, "required": ["result"], "type": "object", "x-cfLinkErrors": {"10001": "Unsupported or unexpected Content Type", "10002": "Unexpected internal server error", "10003": "Missing required URL parameter", "10004": "Malformed URL parameter", "10006": "Unparseable script body", "10007": "Resource not found (similar to HTTP 404)", "10015": "The current account is not authorized to use workers", "10018": "Attempted to update a script where the e-tag does not match", "10021": "Script content failed validation checks, but was otherwise parseable", "10023": "Unauthorized access attempt", "10027": "Script body was too large", "10075": "Requires a Workers Paid plan"}}]}
```
