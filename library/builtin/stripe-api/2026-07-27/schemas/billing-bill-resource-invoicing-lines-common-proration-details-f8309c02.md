---
title: billing_bill_resource_invoicing_lines_common_proration_details
page_id: schema-billing-bill-resource-invoicing-lines-common-proration-details-f8309c02
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_bill_resource_invoicing_lines_common_proration_details

```yaml
{"title": "BillingBillResourceInvoicingLinesCommonProrationDetails", "type": "object", "properties": {"credited_items": {"description": "For a credit proration `line_item`, the original debit line_items to which the credit proration applies.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/billing_bill_resource_invoicing_lines_common_credited_items"}]}}, "description": "", "x-expandableFields": ["credited_items"]}
```
