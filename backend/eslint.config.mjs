import globals from "globals";
import pluginReactConfig from "eslint-plugin-react/configs/recommended.js";
import tseslintPlugin from '@typescript-eslint/eslint-plugin';

import path from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const { configs: tseslintConfigs } = tseslintPlugin;

export default [
  {
    languageOptions: {
      globals: globals.browser,
      parser: "@typescript-eslint/parser",
      ecmaFeatures: {
        jsx: true
      }
    }
  },
  ...tseslintConfigs.recommended,
  {
    ...pluginReactConfig,
    settings: {
      react: {
        version: "detect"
      }
    }
  },
];
