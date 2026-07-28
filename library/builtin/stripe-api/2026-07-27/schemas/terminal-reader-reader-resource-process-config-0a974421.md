---
title: terminal_reader_reader_resource_process_config
page_id: schema-terminal-reader-reader-resource-process-config-0a974421
path: schemas
description: Represents a per-transaction override of a reader configuration
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_process_config

Represents a per-transaction override of a reader configuration

```yaml
{"title": "TerminalReaderReaderResourceProcessConfig", "type": "object", "properties": {"enable_customer_cancellation": {"type": "boolean", "description": "Enable customer-initiated cancellation when processing this payment."}, "return_url": {"maxLength": 5000, "type": "string", "description": "If the customer doesn't abandon authenticating the payment, they're redirected to this URL after completion."}, "skip_tipping": {"type": "boolean", "description": "Override showing a tipping selection screen on this transaction."}, "tipping": {"$ref": "#/components/schemas/terminal_reader_reader_resource_tipping_config"}}, "description": "Represents a per-transaction override of a reader configuration", "x-expandableFields": ["tipping"]}
```
