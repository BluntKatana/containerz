
module.exports = {
  extends: ['plugin:svelte/recommended'],
  parser: '@typescript-eslint/parser',
  parserOptions: {
    // project from folder the cjs file is in
    parser: '@typescript-eslint/parser',
    project: './tsconfig.json',
    // @ts-ignore
    tsconfigRootDir: __dirname,
    extraFileExtensions: ['.svelte'] // This is a required setting in `@typescript-eslint/parser` v4.24.0.
  },
  overrides: [
    {
      files: ['*.svelte'],
      parser: 'svelte-eslint-parser',
      // Parse the `<script>` in `.svelte` as TypeScript by adding the following configuration.
      parserOptions: {
        parser: '@typescript-eslint/parser'
      }
    }
  ]
};