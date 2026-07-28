---
title: images_sourcingkit_connectivity_precheck_request
page_id: schema-images-sourcingkit-connectivity-precheck-request-96bdc7b2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# images_sourcingkit_connectivity_precheck_request

```yaml
{"type": "object", "properties": {"account": {"description": "Account identifier for the bucket (required for R2 vendor).", "type": "string"}, "bucket": {"description": "The name of the storage bucket.", "type": "string", "maxLength": 128, "minLength": 1}, "region": {"description": "The region hint for the bucket (S3 only).", "type": "string", "nullable": true}, "secret": {"description": "Storage credentials for accessing the bucket.", "type": "object", "writeOnly": true}, "vendor": {"$ref": "#/components/schemas/images_sourcingkit_vendor"}}, "required": ["bucket", "vendor", "secret"]}
```
