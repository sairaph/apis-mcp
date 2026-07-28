---
title: workers_secret
page_id: schema-workers-secret-647d05cf
path: schemas
description: A secret value accessible through a binding.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_secret

A secret value accessible through a binding.

```yaml
{"description": "A secret value accessible through a binding.", "type": "object", "discriminator": {"mapping": {"secret_key": "#/components/schemas/workers_binding_kind_secret_key", "secret_text": "#/components/schemas/workers_binding_kind_secret_text"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/workers_binding_kind_secret_text"}, {"$ref": "#/components/schemas/workers_binding_kind_secret_key"}]}
```
