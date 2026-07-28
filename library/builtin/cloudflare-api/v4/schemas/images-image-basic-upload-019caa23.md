---
title: images_image_basic_upload
page_id: schema-images-image-basic-upload-019caa23
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# images_image_basic_upload

```yaml
{"type": "object", "properties": {"creator": {"description": "Can set the creator field with an internal user ID.", "type": "string", "maxLength": 1024}, "file": {"description": "An image binary data. Only needed when type is uploading a file.", "type": "string", "format": "binary", "x-auditable": true}, "id": {"description": "An optional custom unique identifier for your image.", "type": "string", "x-auditable": true}, "metadata": {"description": "User modifiable key-value store. Can use used for keeping references to another system of record for managing images.", "type": "object"}, "requireSignedURLs": {"description": "Indicates whether the image requires a signature token for the access.", "type": "boolean", "example": true, "default": false, "x-auditable": true}, "url": {"description": "A URL to fetch an image from origin. Only needed when type is uploading from a URL.", "type": "string", "example": "https://example.com/path/to/logo.png", "x-auditable": true}}}
```
