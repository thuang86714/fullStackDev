module.exports = {
  env: {
    browser: true,
    es2021: true,
    node: true,
  },
  extends: [
    'eslint:recommended',
    'plugin:@typescript-eslint/recommended',
  ],
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaVersion: 12,
    sourceType: 'module',
    ecmaFeatures: {
      jsx: true,  // If you use JSX, keep this
    },
  },
  plugins: [
    '@typescript-eslint',
  ],
  rules: {
    // Specific rule adjustments
    'no-undef': 'warn',  // Changing undefined variables to warnings or disabling
    'no-empty': 'warn',  // Allowing empty blocks as warnings
    '@typescript-eslint/no-unused-vars': ['warn', { argsIgnorePattern: '^_' }], // Ignore unused vars starting with _
    'no-useless-escape': 'off', // Disabling useless escapes if they're intentional
  },
  globals: {
    define: 'readonly',
    $: 'readonly',
    Cookies: 'readonly',
    window: 'readonly',  // Add window if you're using it globally
    document: 'readonly' // Add document if you're directly manipulating it
  }
};
