---
title: workers_namespace-response
page_id: schema-workers-namespace-response-fe232f14
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_namespace-response

```yaml
{"type": "object", "properties": {"created_by": {"$ref": "#/components/schemas/workers_identifier"}, "created_on": {"$ref": "#/components/schemas/workers_created_on"}, "modified_by": {"$ref": "#/components/schemas/workers_identifier"}, "modified_on": {"$ref": "#/components/schemas/workers_modified_on"}, "namespace_id": {"$ref": "#/components/schemas/workers_uuid"}, "namespace_name": {"$ref": "#/components/schemas/workers_dispatch_namespace_name"}, "script_count": {"$ref": "#/components/schemas/workers_script_count"}, "trusted_workers": {"$ref": "#/components/schemas/workers_trusted_workers"}}}
```
