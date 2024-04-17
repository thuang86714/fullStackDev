import globals from "globals";
import tseslint from "typescript-eslint";
import pluginReactConfig from "eslint-plugin-react/configs/recommended.js";

import path from "path";
import { fileURLToPath } from "url";
import { FlatCompat } from "@eslint/eslintrc";
import pluginJs from "@eslint/js";

// mimic CommonJS variables -- not needed if using CommonJS
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);
const compat = new FlatCompat({baseDirectory: __dirname, recommendedConfig: pluginJs.configs.recommended});

export default [
  {languageOptions: {
    globals: globals.browser,
    parser: "@typescript-eslint/parser", // This explicitly sets the parser for TypeScript.
    ecmaFeatures: {
      jsx: true  // Support for JSX syntax
    }
  }},
  ...compat.extends("standard-with-typescript"),
  ...tseslint.configs.recommended,
  {
    // Integrating React plugin configuration into the flat config
    ...pluginReactConfig,
    settings: {
      react: {
        version: "detect"  // Automatically detect the React version
      }
    }
  },
];
