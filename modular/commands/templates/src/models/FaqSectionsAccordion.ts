import { Model } from '@stackbit/types';

export const FaqSectionsAccordion: Model = {
  type: 'object',
  name: 'FaqSectionsAccordion',
  fields: [
    { name: 'heading', label: 'Heading', type: 'text' },
    { name: 'subheading', label: 'Subheading', type: 'text' },
  ]
};