// @ts-expect-error
import { defaultRenderers } from "svelte-markdown/src/markdown-parser"
import CodeBlock from "$lib/components/CodeBlock.svelte"
let renderers = defaultRenderers
renderers.code = CodeBlock
export default renderers
