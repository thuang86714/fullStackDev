module.exports = {
  env: {
    node: true,  // Specifies that the code is running in a Node.js environment
    es2021: true // Use modern ECMAScript features
  },
  extends: [
    'eslint:recommended',  // Start with a set of recommended rules by ESLint
  ],
  parserOptions: {
    ecmaVersion: 12,  // Allows for the parsing of modern ECMAScript features
    sourceType: 'module'  // Enable ECMAScript modules
  },
  rules: {
    // You can customize your rules here. For example:
    'no-unused-vars': ['warn', { vars: 'all', args: 'after-used', ignoreRestSiblings: false }],
    'no-console': 'off',  // Since you're likely using console in the backend development
  }
};
