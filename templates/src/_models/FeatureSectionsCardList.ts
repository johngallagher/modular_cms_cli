import { Model } from '@stackbit/types';

export const FeatureSectionsCardList: Model = {
  type: 'object',
  name: 'FeatureSectionsCardList',
  fields: [
    {
      type: 'string',
      name: 'heading',
      label: 'Heading',
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
          }
        ]
      }
    }
  ]
};