

export const index = 1;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/fallbacks/error.svelte.js')).default;
export const imports = ["_app/immutable/nodes/1.a1cb6bee.js","_app/immutable/chunks/index.63dde02a.js","_app/immutable/chunks/singletons.8424f29c.js"];
export const stylesheets = [];
export const fonts = [];
