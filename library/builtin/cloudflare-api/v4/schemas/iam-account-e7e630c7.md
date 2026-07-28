---
title: iam_account
page_id: schema-iam-account-e7e630c7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_account

```yaml
{"type": "object", "properties": {"created_on": {"description": "Timestamp for the creation of the account", "type": "string", "format": "date-time", "example": "2014-03-01T12:21:02.0000Z", "readOnly": true, "x-auditable": true}, "id": {"$ref": "#/components/schemas/iam_common_components-schemas-identifier"}, "managed_by": {"description": "Parent container details", "type": "object", "properties": {"parent_org_id": {"description": "ID of the parent Organization, if one exists", "type": "string", "example": "4536bcfad5faccb111b47003c79917fa", "maxLength": 32, "readOnly": true, "x-auditable": true}, "parent_org_name": {"description": "Name of the parent Organization, if one exists", "type": "string", "example": "Demo Parent Organization", "readOnly": true, "x-auditable": true}}}, "name": {"description": "Account name", "type": "string", "example": "Demo Account", "maxLength": 100, "x-auditable": true}, "settings": {"description": "Account settings", "type": "object", "properties": {"abuse_contact_email": {"description": "Sets an abuse contact email to notify for abuse reports.", "type": "string", "x-auditable": true}, "enforce_twofactor": {"description": "Indicates whether membership in this account requires that\nTwo-Factor Authentication is enabled", "type": "boolean", "default": false, "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}}}, "type": {"$ref": "#/components/schemas/iam_account-type"}}, "required": ["id", "name", "type"]}
```
