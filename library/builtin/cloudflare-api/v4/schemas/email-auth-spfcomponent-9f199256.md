---
title: email-auth_SpfComponent
page_id: schema-email-auth-spfcomponent-9f199256
path: schemas
description: |-
    A single SPF component (mechanism) in the inspection tree.

    The `value` field includes the qualifier prefix (e.g., "-all", "~mx:example.com")
    to match the raw SPF record syntax. The qualifier is also available separately
    in the `result` field for programmatic access.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-auth_SpfComponent

A single SPF component (mechanism) in the inspection tree.

The `value` field includes the qualifier prefix (e.g., "-all", "~mx:example.com")
to match the raw SPF record syntax. The qualifier is also available separately
in the `result` field for programmatic access.

```yaml
{"description": "A single SPF component (mechanism) in the inspection tree.\n\nThe `value` field includes the qualifier prefix (e.g., \"-all\", \"~mx:example.com\")\nto match the raw SPF record syntax. The qualifier is also available separately\nin the `result` field for programmatic access.\n", "type": "object", "properties": {"lookup_count": {"description": "Number of DNS lookups this component requires (per RFC 7208).\n- MX, A, EXISTS, INCLUDE, REDIRECT, PTR: 1\n- IP4, IP6, ALL: 0\n", "type": "integer", "example": 0}, "nested": {"description": "Nested SPF tree for INCLUDE or REDIRECT mechanisms.\nOnly present for INCLUDE/REDIRECT components.\n", "allOf": [{"$ref": "#/components/schemas/email-auth_SpfTree"}]}, "result": {"$ref": "#/components/schemas/email-auth_SpfResult"}, "type": {"description": "Component type (UPPERCASE)", "type": "string", "example": "IP4", "enum": ["ALL", "A", "MX", "IP4", "IP6", "EXISTS", "INCLUDE", "PTR", "REDIRECT"]}, "value": {"description": "Component value with qualifier prefix.\nFor IP4/IP6, stores just the IP address (no \"ip4:\" or \"ip6:\" prefix).\nExamples: \"203.0.113.1\", \"203.0.113.0/24\", \"2001:db8::1/64\", \"-all\", \"~mx:example.com\", \"include:_spf.example.com\"\n", "type": "string", "example": "203.0.113.1"}}, "required": ["type", "value", "result", "lookup_count"]}
```
