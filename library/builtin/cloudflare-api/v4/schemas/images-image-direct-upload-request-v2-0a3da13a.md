---
title: images_image_direct_upload_request_v2
page_id: schema-images-image-direct-upload-request-v2-0a3da13a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# images_image_direct_upload_request_v2

```yaml
{"type": "object", "properties": {"creator": {"description": "Can set the creator field with an internal user ID.", "type": "string"}, "expiry": {"description": "The date after which the upload will not be accepted. Minimum: Now + 2 minutes. Maximum: Now + 6 hours.", "type": "string", "format": "date-time", "example": "2021-01-02T02:20:00Z", "default": "Now + 30 minutes", "x-auditable": true}, "id": {"description": "Optional Image Custom ID. Up to 1024 chars. Can include any number of subpaths, and utf8 characters. Cannot start nor end with a / (forward slash). Cannot be a UUID.", "type": "string", "example": "this/is/my-customid", "maxLength": 1024, "x-auditable": true}, "metadata": {"description": "User modifiable key-value store. Can be used for keeping references to another system of record, for managing images.", "type": "object"}, "requireSignedURLs": {"description": "Indicates whether the image requires a signature token to be accessed.", "type": "boolean", "example": true, "default": false, "x-auditable": true}}}
```
