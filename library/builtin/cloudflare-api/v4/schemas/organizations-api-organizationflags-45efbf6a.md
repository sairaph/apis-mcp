---
title: organizations-api_OrganizationFlags
page_id: schema-organizations-api-organizationflags-45efbf6a
path: schemas
description: Enable features for Organizations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# organizations-api_OrganizationFlags

Enable features for Organizations.

```yaml
{"description": "Enable features for Organizations.", "type": "object", "properties": {"account_creation": {"type": "string"}, "account_deletion": {"type": "string"}, "account_migration": {"type": "string"}, "account_mobility": {"type": "string"}, "sub_org_creation": {"type": "string"}}, "required": ["account_creation", "account_deletion", "account_migration", "account_mobility", "sub_org_creation"]}
```
