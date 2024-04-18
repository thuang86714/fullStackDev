module.exports = {
  env: {
    browser: true, // Frontend code runs in a browser
    es2021: true,
    node: true, // For build tools running in Node.js
  },
  extends: [
    'eslint:recommended',
    'plugin:react/recommended', // React-specific linting rules
    'plugin:@typescript-eslint/recommended', // TypeScript-specific linting rules
    'plugin:react-hooks/recommended', // Enforces React hooks rules
  ],
  parser: '@typescript-eslint/parser', // Specifies the TypeScript parser for ESLint
  parserOptions: {
    ecmaVersion: 12, // Allows for parsing modern ECMAScript features
    sourceType: 'module', // Support syntax for ES Modules
    ecmaFeatures: {
      jsx: true, // Support for JSX
    },
  },
  settings: {
    react: {
      version: 'detect', // Automatically detect the React version
    },
  },
  plugins: [
    'react', // React-specific ESLint plugin
    '@typescript-eslint', // TypeScript-specific ESLint plugin
    'react-hooks', // Adds linting for React hooks
  ],
  rules: {
    // Customize or override ESLint rules here
    'react/react-in-jsx-scope': 'off', // Not needed with React 17 JSX Transform
    'react/prop-types': 'off', // Not needed in TS projects where we use interfaces instead of PropTypes
    'no-unused-vars': 'warn', // Highlight unused variables as warnings
    '@typescript-eslint/no-unused-vars': ['warn', {
      vars: 'all',
      args: 'after-used',
      ignoreRestSiblings: false,
      argsIgnorePattern: '^_'
    }], // Fine-tune the rule for TypeScript
    'no-empty-function': 'off',
    '@typescript-eslint/no-empty-function': ['warn'], // Warn about empty functions
  },
  globals: {
    // Define globals if necessary
  }
};
