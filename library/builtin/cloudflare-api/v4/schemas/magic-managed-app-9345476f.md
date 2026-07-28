---
title: magic_managed_app
page_id: schema-magic-managed-app-9345476f
path: schemas
description: Managed app defined by Cloudflare.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_managed_app

Managed app defined by Cloudflare.

```yaml
{"description": "Managed app defined by Cloudflare.", "type": "object", "properties": {"hostnames": {"$ref": "#/components/schemas/magic_app_hostnames"}, "ip_subnets": {"$ref": "#/components/schemas/magic_app_subnets"}, "managed_app_id": {"$ref": "#/components/schemas/magic_managed_app_id"}, "name": {"$ref": "#/components/schemas/magic_app_name"}, "source_subnets": {"$ref": "#/components/schemas/magic_app_source_subnets"}, "type": {"$ref": "#/components/schemas/magic_app_type"}}, "required": ["managed_app_id"], "title": "Managed App"}
```
