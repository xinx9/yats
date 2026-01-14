import react from "@vitejs/plugin-react"
import * as path from "node:path"
import { defineConfig } from "vitest/config"
import packageJson from "./package.json" with { type: "json" }
// @ts-expect-error not a module
import sass from 'vite-plugin-sass'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [react(), sass()],

    server: {
        open: true,
    },

    test: {
        root: import.meta.dirname,
        name: packageJson.name,
        environment: "jsdom",

        typecheck: {
            enabled: true,
            tsconfig: path.join(import.meta.dirname, "tsconfig.json"),
        },

        globals: true,
        watch: false,
        setupFiles: ["./src/setupTests.ts"],
    },
})
