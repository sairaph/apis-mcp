---
title: d1_read-replication-details-for-request
page_id: schema-d1-read-replication-details-for-request-0f966dea
path: schemas
description: Configuration for D1 read replication.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# d1_read-replication-details-for-request

Configuration for D1 read replication.

```yaml
{"description": "Configuration for D1 read replication.", "type": "object", "properties": {"mode": {"description": "The read replication mode for the database. Use 'auto' to create replicas and allow D1 automatically place them around the world, or 'disabled' to not use any database replicas (it can take a few hours for all replicas to be deleted).", "type": "string", "example": "auto", "enum": ["auto", "disabled"], "x-auditable": true}}, "required": ["mode"], "x-stainless-terraform-configurability": "computed_optional"}
```
