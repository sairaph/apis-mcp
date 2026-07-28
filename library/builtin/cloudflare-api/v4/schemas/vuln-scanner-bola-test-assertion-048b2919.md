---
title: vuln_scanner_bola-test-assertion
page_id: schema-vuln-scanner-bola-test-assertion-048b2919
path: schemas
description: Assertion that was made against the received response.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-test-assertion

Assertion that was made against the received response.

```yaml
{"description": "Assertion that was made against the received response.", "type": "object", "properties": {"description": {"description": "Human-readable description of the assertion, explaining what was checked.", "type": "string"}, "kind": {"description": "Kind of assertion.", "allOf": [{"$ref": "#/components/schemas/vuln_scanner_bola-assertion-kind"}]}, "observed": {"description": "Observed value on which the assertion was made.", "type": "integer", "nullable": true}, "outcome": {"description": "Outcome of the assertion.", "allOf": [{"$ref": "#/components/schemas/vuln_scanner_bola-outcome"}]}}, "required": ["description", "kind", "observed", "outcome"]}
```
