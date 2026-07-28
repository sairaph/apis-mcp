---
title: r2_r2_delete_objects_by_prefix_result
page_id: schema-r2-r2-delete-objects-by-prefix-result-a93432ed
path: schemas
description: |-
    Descriptor of an asynchronous prefix-delete job, returned when the bulk
    delete-objects endpoint is invoked with a `prefix` query parameter.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_r2_delete_objects_by_prefix_result

Descriptor of an asynchronous prefix-delete job, returned when the bulk
delete-objects endpoint is invoked with a `prefix` query parameter.

```yaml
{"description": "Descriptor of an asynchronous prefix-delete job, returned when the bulk\ndelete-objects endpoint is invoked with a `prefix` query parameter.\n", "type": "object", "properties": {"endTime": {"description": "When the job finished, if it has finished.", "type": "string", "format": "date-time"}, "id": {"description": "Unique identifier for the prefix-delete job.", "type": "string"}, "jobType": {"description": "The job kind. Always `prefixDelete` for this endpoint.", "type": "string", "enum": ["prefixDelete"]}, "prefixDelete": {"description": "Details specific to the prefix-delete job.", "type": "object", "properties": {"deletedObjects": {"description": "Number of objects deleted by the job so far.", "type": "integer", "example": 0}, "isBucketClear": {"description": "True when the job was created to clear the entire bucket\n(i.e. an empty prefix). Distinguishes bucket-level clear operations\nfrom prefix-specific delete operations.\n", "type": "boolean"}, "prefix": {"description": "The prefix matched against object keys for deletion.", "type": "string", "example": "path/to/"}}}, "startTime": {"description": "When the job was created.", "type": "string", "format": "date-time"}, "status": {"description": "Lifecycle status of the job.", "type": "string", "enum": ["ENQUEUED", "RUNNING", "COMPLETED", "FAILED", "CANCELLED"]}}}
```
