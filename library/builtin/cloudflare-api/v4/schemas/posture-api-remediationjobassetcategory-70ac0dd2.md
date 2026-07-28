---
title: posture-api_RemediationJobAssetCategory
page_id: schema-posture-api-remediationjobassetcategory-70ac0dd2
path: schemas
description: Category information for a remediation job asset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_RemediationJobAssetCategory

Category information for a remediation job asset.

```yaml
{"description": "Category information for a remediation job asset.", "type": "object", "properties": {"service": {"description": "Specific service within the vendor.", "type": "string", "example": "OneDrive"}, "type": {"description": "Asset type.", "type": "string", "example": "SaaS"}, "vendor": {"$ref": "#/components/schemas/posture-api_VendorsDisplayNameEnum"}}, "required": ["type", "vendor", "service"]}
```
