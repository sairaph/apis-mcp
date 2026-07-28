---
title: resource-sharing_create_share_recipient_request
page_id: schema-resource-sharing-create-share-recipient-request-7c3c09c7
path: schemas
description: |-
    Optionally specify `recipient_account_id` to target a specific account, or `organization_id` to target the caller's whole organization. If neither is provided, the caller's organization is used.
    The legacy field `account_id` is accepted as a synonym for `recipient_account_id` during the deprecation period (see `x-sunset` on that field).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-sharing_create_share_recipient_request

Optionally specify `recipient_account_id` to target a specific account, or `organization_id` to target the caller's whole organization. If neither is provided, the caller's organization is used.
The legacy field `account_id` is accepted as a synonym for `recipient_account_id` during the deprecation period (see `x-sunset` on that field).

```yaml
{"description": "Optionally specify `recipient_account_id` to target a specific account, or `organization_id` to target the caller's whole organization. If neither is provided, the caller's organization is used.\nThe legacy field `account_id` is accepted as a synonym for `recipient_account_id` during the deprecation period (see `x-sunset` on that field).\n", "type": "object", "properties": {"account_id": {"description": "Deprecated alias for `recipient_account_id`. Use `recipient_account_id` instead.\nThe body field collided with the URL path parameter of the same name, which prevented SDK generators from distinguishing the source account (in the URL) from the recipient account (in the body). Both names will continue to be accepted until 2027-05-26 (see `x-sunset`).\n", "allOf": [{"$ref": "#/components/schemas/resource-sharing_account_id"}], "deprecated": true, "x-stainless-deprecation-message": "This field has been renamed to `recipient_account_id`. Both names are accepted during the deprecation period.", "x-stainless-skip": ["terraform"]}, "organization_id": {"$ref": "#/components/schemas/resource-sharing_organization_id"}, "recipient_account_id": {"description": "The account that will receive the share.", "allOf": [{"$ref": "#/components/schemas/resource-sharing_account_id"}]}}}
```
