import { Model } from '@stackbit/types';

export const MarketingHeroCoverImageWithCtas: Model = {
  type: 'object',
  name: 'MarketingHeroCoverImageWithCtas',
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
      type: 'object',
      name: 'left',
      required: true,
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
          type: 'object',
          name: 'cta',
          fields: [
            {
              type: 'string',
              name: 'text',
              label: 'Button Text',
              required: true
            },
            {
              type: 'string',
              name: 'url',
              label: 'Button URL',
              required: true
            }
          ]
        }
      ]
    },
    {
      type: 'object',
      name: 'right',
      required: true,
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
          type: 'object',
          name: 'cta',
          fields: [
            {
              type: 'string',
              name: 'text',
              label: 'Button Text',
              required: true
            },
            {
              type: 'string',
              name: 'url',
              label: 'Button URL',
              required: true
            }
          ]
        }
      ]
    },
    {
      type: 'object',
      name: 'image',
      fields: [
        {
          type: 'string',
          name: 'url',
          label: 'Image URL',
          required: true
        }
      ]
    }
  ]
};
