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
  {
    languageOptions: {
      globals: globals.browser,
      parser: "@typescript-eslint/parser",  // Correct parser setting for TypeScript
      ecmaFeatures: {
        jsx: true,  // JSX support for React
      }
    },
  },
  ...compat.extends("standard-with-typescript"),
  ...tseslint.configs.recommended,
  {
    ...pluginReactConfig,
    settings: {
      react: {
        version: "detect"  // Automatically detect React version for eslint-plugin-react
      }
    }
  },
];
