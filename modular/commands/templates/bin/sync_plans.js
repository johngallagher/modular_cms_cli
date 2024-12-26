#!/usr/bin/env node

import fs from 'fs';
import path from 'path';
import 'dotenv/config'
import Stripe from 'stripe';

async function syncPlans() {
  try {
    const stripe = new Stripe(process.env.STRIPE_SECRET_KEY);
    
    // Get all products and payment links
    const productsFromStripe = (await stripe.products.list({
      expand: ['data.default_price']
    })).data;
    
    const paymentLinks = (await stripe.paymentLinks.list({
      active: true,
      expand: ['data.line_items']
    })).data;

    // Create a map of product IDs to their payment links
    const productToPaymentLink = {};
    paymentLinks.forEach(link => {
      if (link.line_items?.data?.[0]?.price?.product) {
        const productId = link.line_items.data[0].price.product;
        productToPaymentLink[productId] = link;
      }
    });

    // Filter and transform products that have payment links
    const plans = productsFromStripe
      .filter(product => productToPaymentLink[product.id])
      .map(product => {
        const paymentLink = productToPaymentLink[product.id];
        const price = paymentLink.line_items.data[0].price;
        
        return {
          id: product.id,
          name: product.metadata.name,
          subheading: product.metadata.description,
          price: price.unit_amount / 100,
          features: product.marketing_features?.map(f => f.name) || [],
          cta: {
            text: "Get started",
            url: paymentLink.url
          },
          product_line: {
            id: product.metadata['product_line.id'] || "",
            name: product.metadata['product_line.name'] || ""
          }
        };
      });

    // Create data directory if it doesn't exist
    const dataDir = path.join(process.cwd(), 'data');
    if (!fs.existsSync(dataDir)) {
      fs.mkdirSync(dataDir);
    }

    fs.writeFileSync(
      path.join(dataDir, 'plans.json'),
      JSON.stringify(plans, null, 2)
    );

    console.log('Successfully synced plans to plans.json');
  } catch (error) {
    console.error('Error syncing plans:', error);
    process.exit(1);
  }
}

syncPlans(); 