---
title: magic_app_config
page_id: schema-magic-app-config-41971d96
path: schemas
description: Traffic decision configuration for an app.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_app_config

Traffic decision configuration for an app.

```yaml
{"description": "Traffic decision configuration for an app.", "type": "object", "allOf": [{"properties": {"breakout": {"$ref": "#/components/schemas/magic_app_breakout"}, "id": {"$ref": "#/components/schemas/magic_identifier"}, "preferred_wans": {"$ref": "#/components/schemas/magic_app_breakout_preferred_wans"}, "priority": {"$ref": "#/components/schemas/magic_app_priority"}, "site_id": {"$ref": "#/components/schemas/magic_identifier"}}, "required": ["id", "site_id"]}, {"oneOf": [{"properties": {"account_app_id": {"$ref": "#/components/schemas/magic_account_app_id"}}, "required": ["account_app_id"], "title": "Account App"}, {"properties": {"managed_app_id": {"$ref": "#/components/schemas/magic_managed_app_id"}}, "required": ["managed_app_id"], "title": "Managed App"}]}]}
```
