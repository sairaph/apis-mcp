---
title: teams-devices_file_input_request
page_id: schema-teams-devices-file-input-request-40638304
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_file_input_request

```yaml
{"type": "object", "properties": {"exists": {"description": "Whether or not file exists.", "type": "boolean", "example": true, "x-auditable": true}, "operating_system": {"description": "Operating system.", "type": "string", "example": "mac", "enum": ["windows", "linux", "mac"], "x-auditable": true}, "path": {"description": "File path.", "type": "string", "example": "/bin/cat", "x-auditable": true}, "sha256": {"description": "SHA-256.", "type": "string", "example": "https://api.us-2.crowdstrike.com"}, "thumbprint": {"description": "Signing certificate thumbprint.", "type": "string", "example": "0aabab210bdb998e9cf45da2c9ce352977ab531c681b74cf1e487be1bbe9fe6e"}}, "required": ["path", "operating_system"], "title": "File Check"}
```
