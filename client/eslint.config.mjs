import { dirname } from "path";
import { fileURLToPath } from "url";
import { FlatCompat } from "@eslint/eslintrc";
import js from "@eslint/js";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const compat = new FlatCompat({
  baseDirectory: __dirname,
  recommendedConfig: js.configs.recommended,
});

const eslintConfig = [
  ...compat.extends("next/core-web-vitals"),
  ...compat.extends("eslint-config-prettier"),
  {
    files: ["**/*.{js,mjs,ts,tsx}"],
    plugins: {
      prettier: (await import("eslint-plugin-prettier")).default,
    },
    rules: {
      "prettier/prettier": [
        "error",
        {
          semi: true,
          singleQuote: true,
          trailingComma: "none",
          tabWidth: 2,
          useTabs: false,
          printWidth: 80,
          bracketSpacing: true,
          arrowParens: "avoid",
        },
      ],
    },
  },
];

export default eslintConfig;
