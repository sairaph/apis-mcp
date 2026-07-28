---
title: images_image_patch_request
page_id: schema-images-image-patch-request-b718b496
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# images_image_patch_request

```yaml
{"type": "object", "properties": {"creator": {"description": "Can set the creator field with an internal user ID.", "type": "string"}, "metadata": {"description": "User modifiable key-value store. Can be used for keeping references to another system of record for managing images. No change if not specified.", "type": "object"}, "requireSignedURLs": {"description": "Indicates whether the image can be accessed using only its UID. If set to `true`, a signed token needs to be generated with a signing key to view the image. Returns a new UID on a change. No change if not specified.", "type": "boolean", "example": true, "x-auditable": true}}}
```
