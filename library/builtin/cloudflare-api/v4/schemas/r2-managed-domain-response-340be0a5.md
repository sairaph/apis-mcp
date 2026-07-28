---
title: r2_managed_domain_response
page_id: schema-r2-managed-domain-response-340be0a5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_managed_domain_response

```yaml
{"type": "object", "properties": {"bucketId": {"description": "Bucket ID.", "type": "string", "maxLength": 32, "x-auditable": true}, "domain": {"description": "Domain name of the bucket's r2.dev domain.", "type": "string", "x-auditable": true}, "enabled": {"description": "Whether this bucket is publicly accessible at the r2.dev domain.", "type": "boolean", "x-auditable": true}}, "example": {"bucketId": "0113a9e4549cf9b1ff1bf56e04da0cef", "domain": "pub-0113a9e4549cf9b1ff1bf56e04da0cef.r2.dev", "enabled": true}, "required": ["bucketId", "domain", "enabled"]}
```
