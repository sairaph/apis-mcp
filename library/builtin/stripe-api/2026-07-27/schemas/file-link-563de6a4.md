---
title: file_link
page_id: schema-file-link-563de6a4
path: schemas
description: |-
    To share the contents of a `File` object with non-Stripe users, you can
    create a `FileLink`. `FileLink`s contain a URL that you can use to
    retrieve the contents of the file without authentication.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# file_link

To share the contents of a `File` object with non-Stripe users, you can
create a `FileLink`. `FileLink`s contain a URL that you can use to
retrieve the contents of the file without authentication.

```yaml
{"title": "FileLink", "required": ["created", "expired", "file", "id", "livemode", "metadata", "object"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "expired": {"type": "boolean", "description": "Returns if the link is already expired."}, "expires_at": {"type": "integer", "description": "Time that the link expires.", "format": "unix-time", "nullable": true}, "file": {"description": "The file object this link points to.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/file"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/file"}]}}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["file_link"]}, "url": {"maxLength": 5000, "type": "string", "description": "The publicly accessible URL to download the file.", "nullable": true}}, "description": "To share the contents of a `File` object with non-Stripe users, you can\ncreate a `FileLink`. `FileLink`s contain a URL that you can use to\nretrieve the contents of the file without authentication.", "x-expandableFields": ["file"], "x-resourceId": "file_link"}
```
