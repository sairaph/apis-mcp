---
title: vuln_scanner_create-credential-request
page_id: schema-vuln-scanner-create-credential-request-046f532e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_create-credential-request

```yaml
{"type": "object", "properties": {"location": {"$ref": "#/components/schemas/vuln_scanner_credential-location"}, "location_name": {"description": "Name of the header or cookie where the credential is attached.\n", "type": "string", "example": "Authorization"}, "name": {"description": "Human-readable name.", "type": "string", "example": "Admin API key"}, "value": {"description": "The credential value (e.g. API key, session token). Write-only.\nNever returned in responses.\n", "type": "string", "example": "Bearer EXAMPLE_TOKEN", "writeOnly": true, "x-sensitive": true}}, "required": ["name", "location", "location_name", "value"]}
```
