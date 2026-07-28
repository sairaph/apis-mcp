---
title: iam_scim_filter_feature
page_id: schema-iam-scim-filter-feature-e869b5a9
path: schemas
description: Configuration for SCIM filtering operations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_filter_feature

Configuration for SCIM filtering operations.

```yaml
{"description": "Configuration for SCIM filtering operations.", "type": "object", "properties": {"maxResults": {"description": "The maximum number of filter results per page.", "type": "integer", "example": 100}, "supported": {"description": "Whether filtering is supported.", "type": "boolean", "example": true}}, "required": ["supported", "maxResults"], "title": "SCIM Filter Feature"}
```
