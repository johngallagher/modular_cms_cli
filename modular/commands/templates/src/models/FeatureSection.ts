import { Model } from '@stackbit/types';

export const FeatureSection: Model = {
  type: 'object',
  name: 'FeatureSection',
  fields: [
    {
      type: 'boolean',
      name: 'hide_from_nav',
      label: 'Hide From Navigation',
      default: true,
      required: true
    },
    {
      type: 'enum',
      name: 'view',
      label: 'View',
      default: 'card-list',
      required: true,
      options: [
        { value: 'cta-list', label: 'CTA List' },
        { value: 'card-list', label: 'Card List' },
        { value: 'icons', label: 'Icons' }
      ]
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
            type: 'enum',
            name: 'icon',
            label: 'Icon',
            required: false,
            options: [
              { value: 'solid-bed', label: 'Bed' },
            ]
          }
        ]
      }
    }
  ]
};

