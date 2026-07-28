---
title: security-center_scansListResponse
page_id: schema-security-center-scanslistresponse-6f436c84
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# security-center_scansListResponse

```yaml
{"type": "object", "properties": {"quota": {"$ref": "#/components/schemas/security-center_scanQuota"}, "scans": {"description": "List of on-demand scans.", "type": "array", "items": {"$ref": "#/components/schemas/security-center_scanStatusResponse"}}}, "required": ["scans", "quota"]}
```
