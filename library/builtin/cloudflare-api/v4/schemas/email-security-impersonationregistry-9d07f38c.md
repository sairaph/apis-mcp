---
title: email-security_ImpersonationRegistry
page_id: schema-email-security-impersonationregistry-9d07f38c
path: schemas
description: An impersonation registry entry.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_ImpersonationRegistry

An impersonation registry entry.

```yaml
{"description": "An impersonation registry entry.", "type": "object", "properties": {"comments": {"type": "string", "nullable": true}, "created_at": {"allOf": [{"$ref": "#/components/schemas/email-security_timestamp"}], "readOnly": true}, "directory_id": {"type": "integer", "nullable": true}, "directory_node_id": {"type": "integer", "nullable": true}, "email": {"type": "string", "example": "john.doe@example.com", "x-auditable": true}, "external_directory_node_id": {"type": "string", "deprecated": true, "nullable": true, "x-stainless-deprecation-message": "This field is deprecated."}, "id": {"allOf": [{"$ref": "#/components/schemas/email-security_ImpersonationRegistryId"}], "readOnly": true}, "is_email_regex": {"type": "boolean", "example": false, "x-auditable": true}, "last_modified": {"description": "Deprecated, use `modified_at` instead. End of life: November 1, 2026.", "allOf": [{"$ref": "#/components/schemas/email-security_timestamp"}], "deprecated": true, "readOnly": true, "x-stainless-deprecation-message": "Use `modified_at` instead."}, "modified_at": {"allOf": [{"$ref": "#/components/schemas/email-security_timestamp"}], "readOnly": true}, "name": {"type": "string", "example": "John Doe", "maxLength": 1024, "x-auditable": true}, "provenance": {"$ref": "#/components/schemas/email-security_Provenance"}}}
```
