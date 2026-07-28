---
title: vuln_scanner_bola-test-credential-set
page_id: schema-vuln-scanner-bola-test-credential-set-947e69ca
path: schemas
description: Credential set that was used.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-test-credential-set

Credential set that was used.

```yaml
{"description": "Credential set that was used.", "type": "object", "properties": {"id": {"description": "ID of the credential set.", "type": "string", "format": "uuid"}, "role": {"description": "Role of the credential set.", "allOf": [{"$ref": "#/components/schemas/vuln_scanner_bola-credential-role"}]}}, "required": ["id", "role"]}
```
