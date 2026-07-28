---
title: d1_database-details-response
page_id: schema-d1-database-details-response-4b65594f
path: schemas
description: The details of the D1 database.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# d1_database-details-response

The details of the D1 database.

```yaml
{"description": "The details of the D1 database.", "type": "object", "properties": {"created_at": {"$ref": "#/components/schemas/d1_created-at"}, "file_size": {"$ref": "#/components/schemas/d1_file-size"}, "jurisdiction": {"$ref": "#/components/schemas/d1_jurisdiction-nullable"}, "name": {"$ref": "#/components/schemas/d1_database-name"}, "num_tables": {"$ref": "#/components/schemas/d1_table-count"}, "read_replication": {"$ref": "#/components/schemas/d1_read-replication-details-for-response"}, "uuid": {"$ref": "#/components/schemas/d1_database-identifier"}, "version": {"$ref": "#/components/schemas/d1_database-version"}}}
```
