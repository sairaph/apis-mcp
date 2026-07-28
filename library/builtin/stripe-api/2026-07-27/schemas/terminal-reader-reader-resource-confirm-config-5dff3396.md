---
title: terminal_reader_reader_resource_confirm_config
page_id: schema-terminal-reader-reader-resource-confirm-config-5dff3396
path: schemas
description: Represents a per-transaction override of a reader configuration
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_confirm_config

Represents a per-transaction override of a reader configuration

```yaml
{"title": "TerminalReaderReaderResourceConfirmConfig", "type": "object", "properties": {"return_url": {"maxLength": 5000, "type": "string", "description": "If the customer doesn't abandon authenticating the payment, they're redirected to this URL after completion."}}, "description": "Represents a per-transaction override of a reader configuration", "x-expandableFields": []}
```
