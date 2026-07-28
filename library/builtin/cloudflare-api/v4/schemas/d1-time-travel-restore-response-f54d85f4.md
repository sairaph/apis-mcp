---
title: d1_time-travel-restore-response
page_id: schema-d1-time-travel-restore-response-f54d85f4
path: schemas
description: Response from a time travel restore operation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# d1_time-travel-restore-response

Response from a time travel restore operation.

```yaml
{"description": "Response from a time travel restore operation.", "type": "object", "properties": {"bookmark": {"allOf": [{"$ref": "#/components/schemas/d1_time-travel-bookmark"}, {"description": "The new bookmark representing the state of the database after the restore operation."}]}, "message": {"description": "A message describing the result of the restore operation.", "type": "string", "example": "Database restored successfully", "x-auditable": true}, "previous_bookmark": {"allOf": [{"$ref": "#/components/schemas/d1_time-travel-bookmark"}, {"description": "The bookmark representing the state of the database before the restore operation. Can be used to undo the restore if needed."}]}}}
```
