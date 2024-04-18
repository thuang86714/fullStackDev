import globals from "globals";
import { configs as tseslintConfigs } from "@typescript-eslint/eslint-plugin";
import pluginReactConfig from "eslint-plugin-react/configs/recommended.js";

import path from "path";
import { fileURLToPath } from "url";

// mimic CommonJS variables -- not needed if using CommonJS
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

export default [
  {languageOptions: {
    globals: globals.browser,
    parser: "@typescript-eslint/parser", // This explicitly sets the parser for TypeScript.
    ecmaFeatures: {
      jsx: true  // Support for JSX syntax
    }
  }},
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
