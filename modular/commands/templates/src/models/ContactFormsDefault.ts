import { Model } from '@stackbit/types';

export const ContactFormsDefault: Model = {
  type: 'object',
  name: 'ContactFormsDefault',
  fields: [
    { name: 'heading', label: 'Heading', type: 'text' },
    { name: 'subheading', label: 'Subheading', type: 'text' },
    { name: 'submit_text', label: 'Submit Text', type: 'text' },
  ]
};