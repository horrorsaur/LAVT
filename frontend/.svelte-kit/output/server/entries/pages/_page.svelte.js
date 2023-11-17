import { c as create_ssr_component, d as add_attribute, e as escape } from "../../chunks/index.js";
const logo = "/_app/immutable/assets/logo-universal.157a874a.png";
const _page_svelte_svelte_type_style_lang = "";
const css = {
  code: "#logo.svelte-140lxdh.svelte-140lxdh{display:block;width:50%;height:50%;margin:auto;padding:10% 0 0;background-position:center;background-repeat:no-repeat;background-size:100% 100%;background-origin:content-box}.result.svelte-140lxdh.svelte-140lxdh{height:20px;line-height:20px;margin:1.5rem auto}.input-box.svelte-140lxdh .btn.svelte-140lxdh{width:60px;height:30px;line-height:30px;border-radius:3px;border:none;margin:0 0 0 20px;padding:0 8px;cursor:pointer}.input-box.svelte-140lxdh .btn.svelte-140lxdh:hover{background-image:linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);color:#333333}.input-box.svelte-140lxdh .input.svelte-140lxdh{border:none;border-radius:3px;outline:none;height:30px;line-height:30px;padding:0 10px;background-color:rgba(240, 240, 240, 1);-webkit-font-smoothing:antialiased}.input-box.svelte-140lxdh .input.svelte-140lxdh:hover{border:none;background-color:rgba(255, 255, 255, 1)}.input-box.svelte-140lxdh .input.svelte-140lxdh:focus{border:none;background-color:rgba(255, 255, 255, 1)}",
  map: null
};
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let resultText = "Please enter your name below 👇";
  let name;
  $$result.css.add(css);
  return `<main><h1>Welcome to the Unofficial Wails.io SvelteKit Template!</h1>
  <p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a>
    to read the documentation
  </p>
  <img alt="Wails logo" id="logo"${add_attribute("src", logo, 0)} class="svelte-140lxdh">
  <div class="result svelte-140lxdh" id="result">${escape(resultText)}</div>
  <div class="input-box svelte-140lxdh" id="input"><input autocomplete="off" class="input svelte-140lxdh" id="name" type="text"${add_attribute("value", name, 0)}>
    <button class="btn svelte-140lxdh">Greet</button></div>
</main>`;
});
export {
  Page as default
};
