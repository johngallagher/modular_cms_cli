import { Model } from '@stackbit/types';

export const FeatureSectionsCtaList: Model = {
  type: 'object',
  name: 'FeatureSectionsCtaList',
  fields: [
    {
      type: 'boolean',
      name: 'hide_from_nav',
      label: 'Hide From Navigation',
      default: true
    },
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
      type: 'list',
      name: 'features',
      label: 'Features',
      required: true,
      items: {
        type: 'object',
        fields: [
          {
            type: 'string',
            name: 'heading',
            label: 'Heading',
            required: true
          },
          {
            type: 'string',
            name: 'summary',
            label: 'Summary',
            required: true
          },
          {
            type: 'string',
            name: 'icon',
            label: 'Icon',
            required: true
          }
        ]
      }
    }
  ]
};
