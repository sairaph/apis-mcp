---
title: r2_temp_access_creds_request
page_id: schema-r2-temp-access-creds-request-1a266f4f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_temp_access_creds_request

```yaml
{"type": "object", "properties": {"bucket": {"description": "Name of the R2 bucket.", "type": "string", "x-auditable": true}, "objects": {"description": "Optional object paths to scope the credentials to.", "type": "array", "items": {"type": "string", "x-auditable": true}}, "parentAccessKeyId": {"description": "The parent access key id to use for signing.", "type": "string"}, "permission": {"description": "Permissions allowed on the credentials.", "type": "string", "enum": ["admin-read-write", "admin-read-only", "object-read-write", "object-read-only"], "x-auditable": true}, "prefixes": {"description": "Optional prefix paths to scope the credentials to.", "type": "array", "items": {"type": "string", "x-auditable": true}}, "ttlSeconds": {"description": "How long the credentials will live for in seconds.", "type": "number", "default": 900, "maximum": 604800, "x-auditable": true}}, "example": {"bucket": "example-bucket", "objects": ["example-object"], "parentAccessKeyId": "example-access-key-id", "permission": "object-read-write", "prefixes": ["example-prefix/"], "ttlSeconds": 3600}, "required": ["bucket", "permission", "ttlSeconds", "parentAccessKeyId"]}
```
