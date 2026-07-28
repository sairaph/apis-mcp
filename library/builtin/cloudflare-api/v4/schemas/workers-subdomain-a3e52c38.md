---
title: workers_subdomain
page_id: schema-workers-subdomain-a3e52c38
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_subdomain

```yaml
{"type": "object", "properties": {"enabled": {"description": "Whether the Worker is available on the workers.dev subdomain.", "type": "boolean", "example": false, "default": false, "x-auditable": true}, "previews_enabled": {"description": "Whether the Worker's Preview URLs are available on the workers.dev subdomain.", "type": "boolean", "example": false, "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}}, "required": ["enabled", "previews_enabled"]}
```
