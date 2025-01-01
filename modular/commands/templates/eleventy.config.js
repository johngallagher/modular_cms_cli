import pluginWebc from "@11ty/eleventy-plugin-webc";
import MarkdownIt from "markdown-it";
import DOMPurify from "isomorphic-dompurify";
import Stripe from 'stripe';
import fs from 'fs';
import path from 'path';
import { IdAttributePlugin } from "@11ty/eleventy";
import { eleventyImageTransformPlugin } from "@11ty/eleventy-img";

export default function (eleventyConfig) {
  eleventyConfig.addPlugin(
    eleventyImageTransformPlugin,
    {
      formats: ["webp", "jpeg"],
      defaultAttributes: {
        loading: "lazy",
        decoding: "async",
        sizes: "auto",
      },
    }
  );
  eleventyConfig.addPassthroughCopy({ "src/images/favicon": "/" });
  eleventyConfig.ignores.add("./README.md");
  eleventyConfig.addWatchTarget(
    "src/components/**/*.css",
    "src/models/**/*.ts",
  );
  eleventyConfig.addPlugin(pluginWebc, {
    components: ["src/components/**/*.webc"],
  });
  eleventyConfig.addPlugin(IdAttributePlugin);
  eleventyConfig.setServerOptions({
    watch: ['_site/**/*.css'],
  });
  eleventyConfig.addPassthroughCopy("src/fonts");
  eleventyConfig.addFilter("markdown", function (content, allowedTags = []) {
    const unsafeHtml = new MarkdownIt({ html: true }).render(content || "");
    return DOMPurify.sanitize(unsafeHtml, { ALLOWED_TAGS: ['span', 'strong', 'em', ...allowedTags] });
  });
  eleventyConfig.addFilter("stripe", function () {
    const stripe = new Stripe(process.env.STRIPE_SECRET_KEY);
    return stripe;
  });

  eleventyConfig.addJavaScriptFunction("read", function (name) {
    const dataPath = path.join(process.cwd(), 'data', `${name}.json`);
    return JSON.parse(fs.readFileSync(dataPath, 'utf8'));
  });

  eleventyConfig.addJavaScriptFunction("priceAscending", function (a, b) {
    return a.price - b.price;
  });
  eleventyConfig.addJavaScriptFunction("priceDescending", function (a, b) {
    return b.price - a.price;
  });

  eleventyConfig.addJavaScriptFunction("nextTitle", function(question, questions) {
    if (!question || !questions || !Array.isArray(questions)) {
      return '';
    }
    const nextQuestion = questions.find(q => q && parseInt(q.id) === (parseInt(question.id) + 1));
    return nextQuestion?.title || 'Submit and Finish';
  });

  eleventyConfig.addJavaScriptFunction("prevTitle", function(question, questions) {
    if (!question || !questions || !Array.isArray(questions)) {
      return '';
    }
    const prevQuestion = questions.find(q => q && parseInt(q.id) === (parseInt(question.id) - 1));
    return prevQuestion?.title || 'Start';
  });

  return {
    markdownTemplateEngine: "webc",
    htmlTemplateEngine: "webc",
    dir: {
      input: 'src',
      output: '_site',
      includes: "includes",
      layouts: "layouts",
      data: "data",
    },
  };
};