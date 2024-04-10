import path from "path";
import { fileURLToPath } from "url";

// Corrected imports
import eslintPluginReact from "eslint-plugin-react";
import eslintPluginTs from "@typescript-eslint/eslint-plugin";

// Mimic CommonJS variables
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

export default {
  root: true,
  parser: "@typescript-eslint/parser", // Use the TypeScript parser
  parserOptions: {
    ecmaVersion: 2020, // Allow modern ECMAScript features
    sourceType: "module", // Allow `import` statements
    ecmaFeatures: {
      jsx: true, // Allow JSX syntax
    },
    project: "./tsconfig.json", // Specify the project's tsconfig
  },
  plugins: [
    "react", // Assumes eslint-plugin-react is installed
    "@typescript-eslint", // Assumes @typescript-eslint/eslint-plugin is installed
  ],
  extends: [
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended", // Use recommended rules from @typescript-eslint/eslint-plugin
    "plugin:react/recommended", // Use recommended rules from eslint-plugin-react
    // Add or replace with other configurations as needed
  ],
  settings: {
    react: {
      version: "detect", // Automatically detect the React version
    },
  },
  rules: {
    // Specify any custom rules
  },
};
