---
title: tax_product_resource_tax_association_transaction_attempts
page_id: schema-tax-product-resource-tax-association-transaction-attempts-29538da8
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_product_resource_tax_association_transaction_attempts

```yaml
{"title": "TaxProductResourceTaxAssociationTransactionAttempts", "required": ["source", "status"], "type": "object", "properties": {"committed": {"$ref": "#/components/schemas/tax_product_resource_tax_association_transaction_attempts_resource_committed"}, "errored": {"$ref": "#/components/schemas/tax_product_resource_tax_association_transaction_attempts_resource_errored"}, "source": {"maxLength": 5000, "type": "string", "description": "The source of the tax transaction attempt. This is either a refund or a payment intent."}, "status": {"maxLength": 5000, "type": "string", "description": "The status of the transaction attempt. This can be `errored` or `committed`."}}, "description": "", "x-expandableFields": ["committed", "errored"]}
```
