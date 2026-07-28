---
title: cc_ApplicationHealth
page_id: schema-cc-applicationhealth-1037a9e5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_ApplicationHealth

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/cc_ApplicationHealthErrors"}, "instances": {"$ref": "#/components/schemas/cc_ApplicationHealthInstances"}, "summary": {"description": "High-level health assessment. Only populated for \"new_instances\" strategy.\nBased on a sample of target-version instances (not a full count).\n- \"pending\": No target-version instances exist yet.\n- \"healthy\": All sampled target-version instances are running/active.\n- \"degraded\": Some sampled instances are still starting or scheduling.\n- \"unhealthy\": One or more sampled instances have failed.\n", "type": "string", "enum": ["healthy", "degraded", "unhealthy", "pending"]}}, "required": ["instances", "errors"]}
```
