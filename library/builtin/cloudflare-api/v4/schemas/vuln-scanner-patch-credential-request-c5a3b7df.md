---
title: vuln_scanner_patch-credential-request
page_id: schema-vuln-scanner-patch-credential-request-c5a3b7df
path: schemas
description: Applies a partial update. Only the provided fields change; omitted fields remain unchanged.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_patch-credential-request

Applies a partial update. Only the provided fields change; omitted fields remain unchanged.

```yaml
{"description": "Applies a partial update. Only the provided fields change; omitted fields remain unchanged.", "type": "object", "properties": {"location": {"$ref": "#/components/schemas/vuln_scanner_credential-location"}, "location_name": {"description": "Name of the header or cookie where the credential is attached.\n", "type": "string", "example": "Authorization"}, "name": {"description": "Human-readable name.", "type": "string", "example": "Admin API key"}, "value": {"description": "The credential value. Write-only. Never returned in responses.\n", "type": "string", "example": "Bearer EXAMPLE_TOKEN", "writeOnly": true, "x-sensitive": true}}}
```
