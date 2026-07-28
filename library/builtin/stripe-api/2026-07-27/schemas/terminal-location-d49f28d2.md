---
title: terminal.location
page_id: schema-terminal-location-d49f28d2
path: schemas
description: |-
    A Location represents a grouping of readers.

    Related guide: [Fleet management](https://docs.stripe.com/terminal/fleet/locations)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal.location

A Location represents a grouping of readers.

Related guide: [Fleet management](https://docs.stripe.com/terminal/fleet/locations)

```yaml
{"title": "TerminalLocationLocation", "required": ["address", "display_name", "id", "livemode", "metadata", "object"], "type": "object", "properties": {"address": {"$ref": "#/components/schemas/address"}, "address_kana": {"$ref": "#/components/schemas/legal_entity_japan_address"}, "address_kanji": {"$ref": "#/components/schemas/legal_entity_japan_address"}, "configuration_overrides": {"maxLength": 5000, "type": "string", "description": "The ID of a configuration that will be used to customize all readers in this location."}, "display_name": {"maxLength": 5000, "type": "string", "description": "The display name of the location."}, "display_name_kana": {"maxLength": 5000, "type": "string", "description": "The Kana variation of the display name of the location."}, "display_name_kanji": {"maxLength": 5000, "type": "string", "description": "The Kanji variation of the display name of the location."}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["terminal.location"]}, "phone": {"maxLength": 5000, "type": "string", "description": "The phone number of the location."}}, "description": "A Location represents a grouping of readers.\n\nRelated guide: [Fleet management](https://docs.stripe.com/terminal/fleet/locations)", "x-expandableFields": ["address", "address_kana", "address_kanji"], "x-resourceId": "terminal.location"}
```
