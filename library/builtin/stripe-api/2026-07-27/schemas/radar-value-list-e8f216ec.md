---
title: radar.value_list
page_id: schema-radar-value-list-e8f216ec
path: schemas
description: |-
    Value lists allow you to group values together which can then be referenced in rules.

    Related guide: [Default Stripe lists](https://docs.stripe.com/radar/lists#managing-list-items)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# radar.value_list

Value lists allow you to group values together which can then be referenced in rules.

Related guide: [Default Stripe lists](https://docs.stripe.com/radar/lists#managing-list-items)

```yaml
{"title": "RadarListList", "required": ["alias", "created", "created_by", "id", "item_type", "list_items", "livemode", "metadata", "name", "object"], "type": "object", "properties": {"alias": {"maxLength": 5000, "type": "string", "description": "The name of the value list for use in rules."}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "created_by": {"maxLength": 5000, "type": "string", "description": "The name or email address of the user who created this value list."}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "item_type": {"type": "string", "description": "The type of items in the value list. One of `card_fingerprint`, `card_bin`, `crypto_fingerprint`, `email`, `ip_address`, `country`, `string`, `case_sensitive_string`, `customer_id`, `account`, `sepa_debit_fingerprint`, or `us_bank_account_fingerprint`.", "enum": ["account", "card_bin", "card_fingerprint", "case_sensitive_string", "country", "crypto_fingerprint", "customer_id", "email", "ip_address", "sepa_debit_fingerprint", "string", "us_bank_account_fingerprint"]}, "list_items": {"title": "RadarListListItemList", "required": ["data", "has_more", "object", "url"], "type": "object", "properties": {"data": {"type": "array", "description": "Details about each object.", "items": {"$ref": "#/components/schemas/radar.value_list_item"}}, "has_more": {"type": "boolean", "description": "True if this list has another page of items after this one that can be fetched."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value. Always has the value `list`.", "enum": ["list"]}, "url": {"maxLength": 5000, "type": "string", "description": "The URL where this list can be accessed."}}, "description": "List of items contained within this value list.", "x-expandableFields": ["data"]}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format."}, "name": {"maxLength": 5000, "type": "string", "description": "The name of the value list."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["radar.value_list"]}}, "description": "Value lists allow you to group values together which can then be referenced in rules.\n\nRelated guide: [Default Stripe lists](https://docs.stripe.com/radar/lists#managing-list-items)", "x-expandableFields": ["list_items"], "x-resourceId": "radar.value_list"}
```
