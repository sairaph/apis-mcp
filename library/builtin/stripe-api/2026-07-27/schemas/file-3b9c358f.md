---
title: file
page_id: schema-file-3b9c358f
path: schemas
description: |-
    This object represents files hosted on Stripe's servers. You can upload
    files with the [create file](https://api.stripe.com#create_file) request
    (for example, when uploading dispute evidence). Stripe also
    creates files independently (for example, the results of a [Sigma scheduled
    query](#scheduled_queries)).

    Related guide: [File upload guide](https://docs.stripe.com/file-upload)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# file

This object represents files hosted on Stripe's servers. You can upload
files with the [create file](https://api.stripe.com#create_file) request
(for example, when uploading dispute evidence). Stripe also
creates files independently (for example, the results of a [Sigma scheduled
query](#scheduled_queries)).

Related guide: [File upload guide](https://docs.stripe.com/file-upload)

```yaml
{"title": "File", "required": ["created", "id", "object", "purpose", "size"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "expires_at": {"type": "integer", "description": "The file expires and isn't available at this time in epoch seconds.", "format": "unix-time", "nullable": true}, "filename": {"maxLength": 5000, "type": "string", "description": "The suitable name for saving the file to a filesystem.", "nullable": true}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "links": {"title": "FileResourceFileLinkList", "required": ["data", "has_more", "object", "url"], "type": "object", "properties": {"data": {"type": "array", "description": "Details about each object.", "items": {"$ref": "#/components/schemas/file_link"}}, "has_more": {"type": "boolean", "description": "True if this list has another page of items after this one that can be fetched."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value. Always has the value `list`.", "enum": ["list"]}, "url": {"maxLength": 5000, "pattern": "^/v1/file_links", "type": "string", "description": "The URL where this list can be accessed."}}, "description": "A list of [file links](https://api.stripe.com#file_links) that point at this file.", "nullable": true, "x-expandableFields": ["data"]}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["file"]}, "purpose": {"type": "string", "description": "The [purpose](https://docs.stripe.com/file-upload#uploading-a-file) of the uploaded file.", "enum": ["account_requirement", "additional_verification", "business_icon", "business_logo", "customer_signature", "dispute_evidence", "document_provider_identity_document", "finance_report_run", "financial_account_statement", "identity_document", "identity_document_downloadable", "issuing_regulatory_reporting", "pci_document", "platform_terms_of_service", "selfie", "sigma_scheduled_query", "tax_document_user_upload", "terminal_android_apk", "terminal_reader_splashscreen", "terminal_wifi_certificate", "terminal_wifi_private_key"], "x-stripeBypassValidation": true}, "size": {"type": "integer", "description": "The size of the file object in bytes."}, "title": {"maxLength": 5000, "type": "string", "description": "A suitable title for the document.", "nullable": true}, "type": {"maxLength": 5000, "type": "string", "description": "The returned file type (for example, `csv`, `pdf`, `jpg`, or `png`).", "nullable": true}, "url": {"maxLength": 5000, "type": "string", "description": "Use your live secret API key to download the file from this URL.", "nullable": true}}, "description": "This object represents files hosted on Stripe's servers. You can upload\nfiles with the [create file](https://api.stripe.com#create_file) request\n(for example, when uploading dispute evidence). Stripe also\ncreates files independently (for example, the results of a [Sigma scheduled\nquery](#scheduled_queries)).\n\nRelated guide: [File upload guide](https://docs.stripe.com/file-upload)", "x-expandableFields": ["links"], "x-resourceId": "file"}
```
