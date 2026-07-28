---
title: deleted_terminal.reader
page_id: schema-deleted-terminal-reader-5e2bd9cf
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# deleted_terminal.reader

```yaml
{"title": "TerminalReaderDeletedReader", "required": ["deleted", "device_type", "id", "object", "serial_number"], "type": "object", "properties": {"deleted": {"type": "boolean", "description": "Always true for a deleted object", "enum": [true]}, "device_type": {"type": "string", "description": "Device type of the reader.", "enum": ["bbpos_chipper2x", "bbpos_wisepad3", "bbpos_wisepos_e", "mobile_phone_reader", "simulated_stripe_s700", "simulated_stripe_s710", "simulated_verifone_m425", "simulated_verifone_p630", "simulated_verifone_ux700", "simulated_verifone_v660p", "simulated_wisepos_e", "stripe_m2", "stripe_s700", "stripe_s710", "verifone_P400", "verifone_m425", "verifone_p630", "verifone_ux700", "verifone_v660p"], "x-stripeBypassValidation": true}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["terminal.reader"]}, "serial_number": {"maxLength": 5000, "type": "string", "description": "Serial number of the reader."}}, "description": "", "x-expandableFields": [], "x-resourceId": "deleted_terminal.reader"}
```
