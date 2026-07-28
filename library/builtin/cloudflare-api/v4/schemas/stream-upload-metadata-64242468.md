---
title: stream_upload_metadata
page_id: schema-stream-upload-metadata-64242468
path: schemas
description: |-
    Comma-separated key-value pairs following the TUS protocol specification. Values are Base-64 encoded.
    Supported keys: `name`, `requiresignedurls`, `allowedorigins`, `thumbnailtimestamppct`, `watermark`, `scheduleddeletion`, `maxdurationseconds`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_upload_metadata

Comma-separated key-value pairs following the TUS protocol specification. Values are Base-64 encoded.
Supported keys: `name`, `requiresignedurls`, `allowedorigins`, `thumbnailtimestamppct`, `watermark`, `scheduleddeletion`, `maxdurationseconds`.

```yaml
{"description": "Comma-separated key-value pairs following the TUS protocol specification. Values are Base-64 encoded.\nSupported keys: `name`, `requiresignedurls`, `allowedorigins`, `thumbnailtimestamppct`, `watermark`, `scheduleddeletion`, `maxdurationseconds`.", "type": "string", "example": "name aGVsbG8gd29ybGQ=, requiresignedurls, allowedorigins ZXhhbXBsZS5jb20sdGVzdC5jb20=", "x-auditable": true}
```
