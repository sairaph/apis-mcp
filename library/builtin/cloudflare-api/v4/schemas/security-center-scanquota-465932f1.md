---
title: security-center_scanQuota
page_id: schema-security-center-scanquota-465932f1
path: schemas
description: Quota information for on-demand scans. Scans are rate limited per account per 24-hour rolling window.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# security-center_scanQuota

Quota information for on-demand scans. Scans are rate limited per account per 24-hour rolling window.

```yaml
{"description": "Quota information for on-demand scans. Scans are rate limited per account per 24-hour rolling window.", "type": "object", "properties": {"available": {"description": "The number of on-demand scans remaining in the current 24-hour window.", "type": "integer", "example": 2}, "used": {"description": "The number of on-demand scans initiated in the current 24-hour window.", "type": "integer", "example": 3}}, "required": ["used", "available"]}
```
