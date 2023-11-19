

export const index = 1;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/fallbacks/error.svelte.js')).default;
export const imports = ["_app/immutable/nodes/1.50aaaa6e.js","_app/immutable/chunks/index.6dba6488.js","_app/immutable/chunks/singletons.ed7ca84a.js"];
export const stylesheets = [];
export const fonts = [];
