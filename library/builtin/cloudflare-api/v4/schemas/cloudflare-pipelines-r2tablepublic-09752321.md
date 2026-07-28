---
title: cloudflare-pipelines_r2TablePublic
page_id: schema-cloudflare-pipelines-r2tablepublic-09752321
path: schemas
description: R2 Sink public configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudflare-pipelines_r2TablePublic

R2 Sink public configuration.

```yaml
{"description": "R2 Sink public configuration.", "type": "object", "properties": {"account_id": {"description": "Cloudflare Account ID for the bucket", "type": "string", "title": "Account ID"}, "bucket": {"description": "R2 Bucket to write to", "type": "string", "title": "Bucket"}, "file_naming": {"description": "Controls filename prefix/suffix and strategy.", "type": "object", "additionalProperties": false, "properties": {"prefix": {"description": "The prefix to use in file name. i.e prefix-<uuid>.parquet", "type": "string", "title": "Filename Prefix"}, "strategy": {"description": "Filename generation strategy.", "type": "string", "enum": ["serial", "uuid", "uuid_v7", "ulid"], "title": "Filename Strategy"}, "suffix": {"description": "This will overwrite the default file suffix. i.e .parquet, use with caution", "type": "string", "title": "Filename Suffix"}}, "title": "File Naming"}, "jurisdiction": {"description": "Jurisdiction this bucket is hosted in", "type": "string", "title": "Jurisdiction"}, "partitioning": {"description": "Data-layout partitioning for sinks.", "type": "object", "additionalProperties": false, "properties": {"time_pattern": {"description": "The pattern of the date string", "type": "string", "example": "year=%Y/month=%m/day=%d/hour=%H", "title": "Time Partition Pattern"}}, "title": "Partitioning Configuration"}, "path": {"description": "Subpath within the bucket to write to", "type": "string", "title": "Path"}, "rolling_policy": {"description": "Rolling policy for file sinks (when & why to close a file and open a new one).", "type": "object", "additionalProperties": false, "properties": {"file_size_bytes": {"description": "Files will be rolled after reaching this number of bytes", "type": "integer", "format": "uint64", "minimum": 0, "title": "File Size"}, "inactivity_seconds": {"description": "Number of seconds of inactivity to wait before rolling over to a new file", "type": "integer", "format": "uint64", "minimum": 1, "title": "Inactivity Seconds"}, "interval_seconds": {"description": "Number of seconds to wait before rolling over to a new file", "type": "integer", "format": "uint64", "minimum": 1, "title": "Interval Seconds"}}, "title": "File Rolling Policy"}}, "additionalProperties": false, "required": ["account_id", "bucket"], "title": "R2 Sink Public"}
```
