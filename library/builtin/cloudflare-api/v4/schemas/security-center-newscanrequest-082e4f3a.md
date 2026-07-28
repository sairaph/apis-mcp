---
title: security-center_newScanRequest
page_id: schema-security-center-newscanrequest-082e4f3a
path: schemas
description: Request body for starting an on-demand scan. Specify issue_type or issue_class to scope the scan, or provide an empty object to scan all issue types.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# security-center_newScanRequest

Request body for starting an on-demand scan. Specify issue_type or issue_class to scope the scan, or provide an empty object to scan all issue types.

```yaml
{"description": "Request body for starting an on-demand scan. Specify issue_type or issue_class to scope the scan, or provide an empty object to scan all issue types.", "oneOf": [{"properties": {"issue_type": {"$ref": "#/components/schemas/security-center_issueType"}}, "required": ["issue_type"], "type": "object"}, {"properties": {"issue_class": {"$ref": "#/components/schemas/security-center_issueClass"}}, "required": ["issue_class"], "type": "object"}, {"additionalProperties": false, "type": "object"}]}
```
