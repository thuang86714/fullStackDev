// Importing necessary plugins and configs directly
import eslintPluginReact from "eslint-plugin-react";

export default {
  files: ["**/*.ts", "**/*.tsx"],
  parser: "@typescript-eslint/parser",
  plugins: {
    "@typescript-eslint": {},
    "react": {}
  },
  extends: [
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:react/recommended"
  ],
  languageOptions: {
    ecmaVersion: 2020,
    sourceType: "module",
    ecmaFeatures: {
      jsx: true,
    },
  },
  settings: {
    react: {
      version: "detect",
    },
  },
};
