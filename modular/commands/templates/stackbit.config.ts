import { defineStackbitConfig } from '@stackbit/types';
import { GitContentSource } from '@stackbit/cms-git';
import { Page } from './src/models/page';
import { MarketingHeroCoverImageWithCtas } from './src/models/MarketingHeroCoverImageWithCtas';
import { PricingTable } from './src/models/PricingTable';
import { FaqSectionsAccordion } from './src/models/FaqSectionsAccordion';
import { ContactFormsDefault } from './src/models/ContactFormsDefault';
import { FeatureSection } from './src/models/FeatureSection';

export default defineStackbitConfig({
    stackbitVersion: '~0.7.0',
    ssgName: 'custom',
    devCommand: './node_modules/.bin/tailwindcss -i src/tailwind.css -c tailwind.config.js -o _site/styles.css && ./node_modules/.bin/eleventy --serve --port {PORT} --incremental',
    experimental: {
        ssg: {
            name: 'eleventy',
            logPatterns: {
                up: ['Server at']
            }
        }
    },
    useESM: true,
    nodeVersion: '20',
    contentSources: [
        new GitContentSource({
            rootPath: __dirname,
            contentDirs: ['src'],
            models: [
                Page,
                MarketingHeroCoverImageWithCtas,
                FeatureSection,
                PricingTable,
                FaqSectionsAccordion,
                ContactFormsDefault
            ],
            assetsConfig: {
                referenceType: 'static',
                staticDir: 'public',
                uploadDir: 'images',
                publicPath: '/'
            }
        })
    ],
    modelExtensions: [
        { name: 'Page', type: 'page', urlPath: '/{slug}' },
    ],
    pagesDir: 'src',
    sidebarButtons: []
});