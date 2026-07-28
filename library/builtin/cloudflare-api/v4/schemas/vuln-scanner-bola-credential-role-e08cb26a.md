---
title: vuln_scanner_bola-credential-role
page_id: schema-vuln-scanner-bola-credential-role-e08cb26a
path: schemas
description: Identifies the role a request was made with. The credential set governs this role. `owner` is the resource owner, `attacker` attempts to access resources.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-credential-role

Identifies the role a request was made with. The credential set governs this role. `owner` is the resource owner, `attacker` attempts to access resources.

```yaml
{"description": "Identifies the role a request was made with. The credential set governs this role. `owner` is the resource owner, `attacker` attempts to access resources.", "type": "string", "enum": ["owner", "attacker"]}
```
