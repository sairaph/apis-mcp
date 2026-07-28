---
title: vuln_scanner_bola-credential-sets
page_id: schema-vuln-scanner-bola-credential-sets-baa24515
path: schemas
description: |-
    Credential set references for a BOLA scan. The scanner uses the
    `owner` credentials for legitimate requests and the `attacker`
    credentials to attempt unauthorized access.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-credential-sets

Credential set references for a BOLA scan. The scanner uses the
`owner` credentials for legitimate requests and the `attacker`
credentials to attempt unauthorized access.

```yaml
{"description": "Credential set references for a BOLA scan. The scanner uses the\n`owner` credentials for legitimate requests and the `attacker`\ncredentials to attempt unauthorized access.\n", "type": "object", "properties": {"attacker": {"description": "Credential set ID for the attacker.", "type": "string", "format": "uuid"}, "owner": {"description": "Credential set ID for the resource owner.", "type": "string", "format": "uuid"}}, "required": ["owner", "attacker"]}
```
