---
title: terminal.reader
page_id: schema-terminal-reader-a623bd90
path: schemas
description: |-
    A Reader represents a physical device for accepting payment details.

    Related guide: [Connecting to a reader](https://docs.stripe.com/terminal/payments/connect-reader)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal.reader

A Reader represents a physical device for accepting payment details.

Related guide: [Connecting to a reader](https://docs.stripe.com/terminal/payments/connect-reader)

```yaml
{"title": "TerminalReaderReader", "required": ["device_type", "id", "label", "livemode", "metadata", "object", "serial_number"], "type": "object", "properties": {"action": {"description": "The most recent action performed by the reader.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/terminal_reader_reader_resource_reader_action"}]}, "device_sw_version": {"maxLength": 5000, "type": "string", "description": "The current software version of the reader.", "nullable": true}, "device_type": {"type": "string", "description": "Device type of the reader.", "enum": ["bbpos_chipper2x", "bbpos_wisepad3", "bbpos_wisepos_e", "mobile_phone_reader", "simulated_stripe_s700", "simulated_stripe_s710", "simulated_verifone_m425", "simulated_verifone_p630", "simulated_verifone_ux700", "simulated_verifone_v660p", "simulated_wisepos_e", "stripe_m2", "stripe_s700", "stripe_s710", "verifone_P400", "verifone_m425", "verifone_p630", "verifone_ux700", "verifone_v660p"], "x-stripeBypassValidation": true}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "ip_address": {"maxLength": 5000, "type": "string", "description": "The local IP address of the reader.", "nullable": true}, "label": {"maxLength": 5000, "type": "string", "description": "Custom label given to the reader for easier identification."}, "last_seen_at": {"type": "integer", "description": "The last time this reader reported to Stripe backend. Timestamp is measured in milliseconds since the Unix epoch. Unlike most other Stripe timestamp fields which use seconds, this field uses milliseconds.", "format": "unix-time", "nullable": true}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "location": {"description": "The location identifier of the reader.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/terminal.location"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/terminal.location"}]}}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["terminal.reader"]}, "serial_number": {"maxLength": 5000, "type": "string", "description": "Serial number of the reader."}, "status": {"type": "string", "description": "The networking status of the reader. We do not recommend using this field in flows that may block taking payments.", "nullable": true, "enum": ["offline", "online"]}}, "description": "A Reader represents a physical device for accepting payment details.\n\nRelated guide: [Connecting to a reader](https://docs.stripe.com/terminal/payments/connect-reader)", "x-expandableFields": ["action", "location"], "x-resourceId": "terminal.reader"}
```
