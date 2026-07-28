---
title: magic_account_app
page_id: schema-magic-account-app-29154f76
path: schemas
description: Custom app defined for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_account_app

Custom app defined for an account.

```yaml
{"description": "Custom app defined for an account.", "type": "object", "properties": {"account_app_id": {"$ref": "#/components/schemas/magic_account_app_id"}, "hostnames": {"$ref": "#/components/schemas/magic_app_hostnames"}, "ip_subnets": {"$ref": "#/components/schemas/magic_app_subnets"}, "name": {"$ref": "#/components/schemas/magic_app_name"}, "source_subnets": {"$ref": "#/components/schemas/magic_app_source_subnets"}, "type": {"$ref": "#/components/schemas/magic_app_type"}}, "required": ["account_app_id"], "title": "Account App"}
```
