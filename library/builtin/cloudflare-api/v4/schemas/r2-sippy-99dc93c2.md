---
title: r2_sippy
page_id: schema-r2-sippy-99dc93c2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_sippy

```yaml
{"type": "object", "properties": {"destination": {"description": "Details about the configured destination bucket.", "type": "object", "properties": {"accessKeyId": {"description": "ID of the Cloudflare API token used when writing objects to this\nbucket.\n", "type": "string"}, "account": {"type": "string", "x-auditable": true}, "bucket": {"description": "Name of the bucket on the provider.", "type": "string", "x-auditable": true}, "provider": {"type": "string", "enum": ["r2"], "x-auditable": true}}}, "enabled": {"description": "State of Sippy for this bucket.", "type": "boolean", "x-auditable": true}, "source": {"description": "Details about the configured source bucket.", "type": "object", "properties": {"bucket": {"description": "Name of the bucket on the provider (AWS, GCS only).", "type": "string", "nullable": true, "x-auditable": true}, "bucketUrl": {"description": "S3-compatible URL (Generic S3-compatible providers only).", "type": "string", "nullable": true, "x-auditable": true}, "container": {"description": "Name of the Azure Blob Storage container (Azure only).", "type": "string", "nullable": true, "x-auditable": true}, "provider": {"type": "string", "enum": ["aws", "gcs", "s3", "azure"], "x-auditable": true}, "region": {"description": "Region where the bucket resides (AWS only).", "type": "string", "nullable": true, "x-auditable": true}}}}}
```
