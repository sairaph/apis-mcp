---
title: cloudflare-pipelines_r2_data_catalogTablePublic
page_id: schema-cloudflare-pipelines-r2-data-catalogtablepublic-cf853dc8
path: schemas
description: R2 Data Catalog Sink public configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudflare-pipelines_r2_data_catalogTablePublic

R2 Data Catalog Sink public configuration.

```yaml
{"description": "R2 Data Catalog Sink public configuration.", "type": "object", "properties": {"account_id": {"description": "Cloudflare Account ID", "type": "string", "format": "uri", "example": "6791973bbe876d621f9e557f891e90b3"}, "bucket": {"description": "The R2 Bucket that hosts this catalog", "type": "string", "example": "my-bucket"}, "namespace": {"description": "Table namespace", "type": "string", "title": "Namespace"}, "rolling_policy": {"description": "Rolling policy for file sinks (when & why to close a file and open a new one).", "type": "object", "additionalProperties": false, "properties": {"file_size_bytes": {"description": "Files will be rolled after reaching this number of bytes", "type": "integer", "format": "uint64", "minimum": 0, "title": "File Size"}, "inactivity_seconds": {"description": "Number of seconds of inactivity to wait before rolling over to a new file", "type": "integer", "format": "uint64", "minimum": 1, "title": "Inactivity Seconds"}, "interval_seconds": {"description": "Number of seconds to wait before rolling over to a new file", "type": "integer", "format": "uint64", "minimum": 1, "title": "Interval Seconds"}}, "title": "File Rolling Policy"}, "table_name": {"description": "Table name", "type": "string", "title": "Table Name"}}, "additionalProperties": false, "required": ["account_id", "bucket", "table_name"], "title": "R2DataCatalogSinkTablePublic"}
```
