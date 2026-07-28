---
title: images_image_direct_upload_response_v2
page_id: schema-images-image-direct-upload-response-v2-126a2caf
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# images_image_direct_upload_response_v2

```yaml
{"allOf": [{"$ref": "#/components/schemas/images_api-response-single"}, {"properties": {"result": {"properties": {"id": {"description": "Image unique identifier.", "type": "string", "example": "e22e9e6b-c02b-42fd-c405-6c32af5fe600", "maxLength": 32, "readOnly": true, "x-auditable": true}, "uploadURL": {"description": "The URL the unauthenticated upload can be performed to using a single HTTP POST (multipart/form-data) request.", "type": "string", "example": "https://upload.imagedelivery.net/FxUufywByo0m2v3xhKSiU8/e22e9e6b-c02b-42fd-c405-6c32af5fe600", "x-auditable": true}}}}}]}
```
