---
title: iam_create-account
page_id: schema-iam-create-account-39dcde68
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_create-account

```yaml
{"type": "object", "properties": {"name": {"description": "Account name", "type": "string"}, "type": {"$ref": "#/components/schemas/iam_account-type"}, "unit": {"description": "information related to the tenant unit, and optionally, an id of the unit to create the account on. see https://developers.cloudflare.com/tenant/how-to/manage-accounts/", "type": "object", "properties": {"id": {"description": "Tenant unit ID", "type": "string", "example": "f267e341f3dd4697bd3b9f71dd96247f", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}}, "x-stainless-terraform-configurability": "computed_optional"}}, "required": ["name"], "title": "Create account"}
```
