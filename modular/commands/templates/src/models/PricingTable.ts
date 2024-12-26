import { Model } from '@stackbit/types';
import fs from 'fs';
import path from 'path';

const dataPath = path.join(process.cwd(), 'data', 'plans.json');
const allPlans = JSON.parse(fs.readFileSync(dataPath, 'utf8'));
const uniquePlans = allPlans
  .filter((plan, index, self) => index === self.findIndex(p => p.product_line.id === plan.product_line.id))
  .filter(plan => plan.product_line.id);

export const PricingTable: Model = {
  type: 'object',
  name: 'PricingTable',
  fields: [
    {
      type: 'string',
      name: 'heading',
      label: 'Heading',
      required: true
    },
    {
      type: 'string',
      name: 'subheading',
      label: 'Subheading',
      required: true
    },
    {
      type: 'enum',
      name: 'product_line_id',
      label: 'Product Line',
      required: true,
      description: 'The product line to use for the pricing table. This is used to filter the products in the Stripe API.',
      options: uniquePlans.map(plan => ({
        label: plan.product_line.name,
        value: plan.product_line.id
      }))
    },
  ]
};
