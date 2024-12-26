export const Page = {
  name: 'Page',
  type: 'page',
  urlPath: '/{slug}',
  filePath: 'src/{slug}.md',
  hideContent: false,
  fields: [
    {
      type: 'list',
      name: 'blocks',
      label: 'Blocks',
      description: 'The blocks to be arranged in the page',
      items: {
        type: 'model',
        models: [
          'MarketingHeroCoverImageWithCtas',
          'FeatureBlock',
          'PricingTable',
          'FaqSectionsAccordion',
          'ContactDefaultForm',
        ]
      }
    }
  ]
};