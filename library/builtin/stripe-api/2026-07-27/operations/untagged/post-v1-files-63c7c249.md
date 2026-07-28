---
title: Create a file
page_id: operation-post-v1-files-f6421712
path: operations/untagged
description: |-
    <p>To upload a file to Stripe, you need to send a request of type <code>multipart/form-data</code>. Include the file you want to upload in the request, and the parameters for creating a file.</p>

    <p>All of Stripe’s officially supported Client libraries support sending <code>multipart/form-data</code>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/files
operation_ids:
    - PostFiles
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a file

`POST /v1/files`

Operation ID: `PostFiles`

<p>To upload a file to Stripe, you need to send a request of type <code>multipart/form-data</code>. Include the file you want to upload in the request, and the parameters for creating a file.</p>

<p>All of Stripe’s officially supported Client libraries support sending <code>multipart/form-data</code>.</p>

## Definition

```yaml
{"summary": "Create a file", "description": "<p>To upload a file to Stripe, you need to send a request of type <code>multipart/form-data</code>. Include the file you want to upload in the request, and the parameters for creating a file.</p>\n\n<p>All of Stripe’s officially supported Client libraries support sending <code>multipart/form-data</code>.</p>", "operationId": "PostFiles", "requestBody": {"content": {"multipart/form-data": {"schema": {"required": ["file", "purpose"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "file": {"type": "string", "description": "A file to upload. Make sure that the specifications follow RFC 2388, which defines file transfers for the `multipart/form-data` protocol.", "format": "binary"}, "file_link_data": {"title": "file_link_creation_params", "required": ["create"], "type": "object", "properties": {"create": {"type": "boolean"}, "expires_at": {"type": "integer", "format": "unix-time"}, "metadata": {"anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}}, "description": "Optional parameters that automatically create a [file link](https://api.stripe.com#file_links) for the newly created file."}, "purpose": {"type": "string", "description": "The [purpose](https://docs.stripe.com/file-upload#uploading-a-file) of the uploaded file.", "enum": ["account_requirement", "additional_verification", "business_icon", "business_logo", "customer_signature", "dispute_evidence", "identity_document", "issuing_regulatory_reporting", "pci_document", "platform_terms_of_service", "tax_document_user_upload", "terminal_android_apk", "terminal_reader_splashscreen", "terminal_wifi_certificate", "terminal_wifi_private_key"], "x-stripeBypassValidation": true}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "file_link_data": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/file"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}, "servers": [{"url": "https://files.stripe.com/"}]}
```
