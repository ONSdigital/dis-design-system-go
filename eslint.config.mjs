import { configs, plugins } from 'eslint-config-airbnb-extended';
import globals from 'globals';

const clean = (obj) => Object.fromEntries(Object.entries(obj).map(([k, v]) => [k.trim(), v]));

export default [
  {
    ignores: ['build/**', 'dist/**', 'megalinter-reports/**'],
  },
  plugins.importX,
  plugins.stylistic,
  ...configs.base.recommended,
  // This is the configuration for Node.js files, which applies to all JS files in the root directory
  {
    files: ['*.js'],
    languageOptions: {
      sourceType: 'commonjs',
      globals: {
        ...clean(globals.node),
      },
    },
  },
  // This is the main configuration for the project, which applies to all JS files in the assets directory
  {
    files: ['assets/**/*.js'],
    languageOptions: {
      globals: {
        ...clean(globals.browser),
        ...clean(globals.es2021),
      },
      ecmaVersion: 'latest',
      sourceType: 'module',
      parserOptions: {
        ecmaVersion: 'latest',
      },
    },
    rules: {
      '@stylistic/max-len': 'off',
      'no-console': [
        'error',
        {
          allow: [
            'warn',
            'error',
          ],
        },
      ],
    },
  },
];
