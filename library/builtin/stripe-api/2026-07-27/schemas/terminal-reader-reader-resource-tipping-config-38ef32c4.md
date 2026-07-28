---
title: terminal_reader_reader_resource_tipping_config
page_id: schema-terminal-reader-reader-resource-tipping-config-38ef32c4
path: schemas
description: Represents a per-transaction tipping configuration
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_tipping_config

Represents a per-transaction tipping configuration

```yaml
{"title": "TerminalReaderReaderResourceTippingConfig", "type": "object", "properties": {"amount_eligible": {"type": "integer", "description": "Amount used to calculate tip suggestions on tipping selection screen for this transaction. Must be a positive integer in the smallest currency unit (e.g., 100 cents to represent $1.00 or 100 to represent ¥100, a zero-decimal currency)."}}, "description": "Represents a per-transaction tipping configuration", "x-expandableFields": []}
```
