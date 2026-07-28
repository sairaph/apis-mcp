---
title: iam_scim_bulk_feature
page_id: schema-iam-scim-bulk-feature-e992fa4c
path: schemas
description: Configuration for SCIM bulk operations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_bulk_feature

Configuration for SCIM bulk operations.

```yaml
{"description": "Configuration for SCIM bulk operations.", "type": "object", "properties": {"maxOperations": {"description": "The maximum number of operations in a bulk request.", "type": "integer", "example": 0}, "maxPayloadSize": {"description": "The maximum payload size in bytes for a bulk request.", "type": "integer", "example": 0}, "supported": {"description": "Whether bulk operations are supported.", "type": "boolean", "example": false}}, "required": ["supported", "maxOperations", "maxPayloadSize"], "title": "SCIM Bulk Feature"}
```
