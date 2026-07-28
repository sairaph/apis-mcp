---
title: vuln_scanner_credential
page_id: schema-vuln-scanner-credential-d7e83353
path: schemas
description: |-
    A credential attached to API requests during scanning. The
    credential `value` is write-only and never returned in responses.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_credential

A credential attached to API requests during scanning. The
credential `value` is write-only and never returned in responses.

```yaml
{"description": "A credential attached to API requests during scanning. The\ncredential `value` is write-only and never returned in responses.\n", "type": "object", "properties": {"credential_set_id": {"description": "Parent credential set identifier.", "type": "string", "format": "uuid", "x-auditable": true}, "id": {"description": "Credential identifier.", "type": "string", "format": "uuid", "x-auditable": true}, "location": {"$ref": "#/components/schemas/vuln_scanner_credential-location"}, "location_name": {"description": "Name of the header or cookie where the credential is attached.\n", "type": "string", "example": "Authorization", "x-auditable": true}, "name": {"description": "Human-readable name.", "type": "string", "example": "Admin API key", "x-auditable": true}}, "required": ["id", "credential_set_id", "name", "location", "location_name"]}
```
