---
title: d1_read-replication-details-for-response
page_id: schema-d1-read-replication-details-for-response-e1863624
path: schemas
description: Configuration for D1 read replication.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# d1_read-replication-details-for-response

Configuration for D1 read replication.

```yaml
{"description": "Configuration for D1 read replication.", "type": "object", "properties": {"mode": {"description": "The read replication mode for the database. Mode 'auto' denotes that D1 creates replicas and automatically places them around the world. Mode 'disabled' denotes that no database replicas are used.", "type": "string", "example": "auto", "enum": ["auto", "disabled"], "x-auditable": true}}, "required": ["mode"]}
```
