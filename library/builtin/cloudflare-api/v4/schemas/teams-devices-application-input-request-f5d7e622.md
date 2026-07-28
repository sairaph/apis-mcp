---
title: teams-devices_application_input_request
page_id: schema-teams-devices-application-input-request-f5d7e622
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_application_input_request

```yaml
{"type": "object", "properties": {"operating_system": {"description": "Operating system.", "type": "string", "example": "mac", "enum": ["windows", "linux", "mac"], "x-auditable": true}, "path": {"description": "Path for the application.", "type": "string", "example": "/bin/cat", "x-auditable": true}, "sha256": {"description": "SHA-256.", "type": "string", "example": "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c", "x-auditable": true}, "thumbprint": {"description": "Signing certificate thumbprint.", "type": "string", "example": "0aabab210bdb998e9cf45da2c9ce352977ab531c681b74cf1e487be1bbe9fe6e", "x-auditable": true}}, "required": ["path", "operating_system"], "title": "Application"}
```
