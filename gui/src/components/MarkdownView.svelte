<script>
import marked from 'marked';
import highlight from 'highlight.js';
import { onDestroy } from 'svelte';
import { todo } from "../stores/todo.js";

// TODO: Add markdown codeblock syntax highlighting
marked.setOptions({
  renderer: new marked.Renderer(),
  highlight: function(code, lang) {
    const hljs = highlight;
    const language = hljs.getLanguage(lang) ? lang : 'plaintext';
    return hljs.highlight(code, { language }).value;
  },
  smartLists: true,
  smartypants: true,
});

let placeholder;
const unsubscribe = todo.subscribe(todo => {
  placeholder = todo.body;
})

let preview = false;
function togglePreview() {
  preview = !preview;
}

onDestroy(unsubscribe);
</script>

{#if preview}
  <div class="markdown visible">
    <div class="input-container">
      <textarea bind:value={$todo.body} id="textArea" class="input" placeholder={placeholder} wrap="off"></textarea>
      <button class="toggle" on:click={togglePreview}>HIDE PREVIEW</button>
    </div>
    <div class="output"><div class="marked">{@html marked($todo.body)}</div></div>
  </div>
{:else}
  <div class="markdown hidden">
    <div class="input-container">
      <textarea bind:value={$todo.body} id="textArea" class="input" placeholder={placeholder} wrap="off"></textarea>
      <button class="toggle" on:click={togglePreview}>SHOW PREVIEW</button>
    </div>
  </div>
{/if}

<style lang="scss">
$bg: #282828;
$fg: #ebdbb2;

$dark_red: #cc241d;
$dark_green: #98971a;
$dark_yellow: #d79921;
$dark_blue: #458588;
$dark_purple: #b16286;
$dark_aqua: #689d6a;
$dark_orange: #d65d0e;

$bright_red: #fb4934;
$bright_green: #b8bb26;
$bright_yellow: #fabd2f;
$bright_blue: #83a598;
$bright_purple: #d3869b;
$bright_aqua: #8ec07c;
$bright_orange: #fe8019;

$bg0_h: #1d2021;
$bg0: $bg;
$bg0_s: #32302f;
$bg1: #3c3836;
$bg2: #504945;
$bg3: #665c54;
$bg4: #7c6f64;
$gray: #928374;
$fg4: #a89984;
$fg3: #bdae93;
$fg2: #d5c4a1;
$fg1: $fg;
$fg0: #fbf1c7;


.markdown {
  height: 94.2vh;
  display: grid;
  border-radius: min(0.75vw, 5.625px) 0 0 min(0.75vw, 5.625px);
  background-color: #00385c;

  &.visible {
    grid-template-columns: 1fr 1fr;
    grid-template-areas: "input output";

    .input-container {
      margin: min(1vw, 7.5px) min(0.5vw, 3.75px) min(1vw, 7.5px) min(1vw, 7.5px);
    }
  }

  &.hidden {
    grid-template-columns: 1fr;
    grid-template-areas: "input";

    .input-container {
      margin: min(1vw, 7.5px);
    }
  }

  .input-container {
    grid-area: input;
    display: grid;
    grid-template-rows: auto min-content;
    grid-template-areas:
      "textarea"
      "toggle";

    .input {
      grid-area: textarea;
      outline: none;
      border: none;
      border-radius: min(0.75vw, 5.625px);
      margin-bottom: min(1vw, 7.5px);
      padding: min(1.8vw, 13.5px);
      overflow: auto;
      resize: none;
      font-size: min(2.25vw, 16.875px);
      font-family: monospace;
      color: #00385c;
      background-color: #fff;
    }

    .toggle {
      grid-area: toggle;
      position: relative;
      top: 0;
      width: min(35vw, 262.5px);
      margin: 0 auto;
      cursor: pointer;
      font-family: "Lexend", sans-serif;
      font-size: min(3vw, 22.5px);
      color: #00385c;
      background-color: #fff;
      outline: none;
      border: none;
      border-radius: min(0.75vw, 5.625px);
      box-shadow: 0 min(0.5vw, 3.75px) 0 #000;
      transition: all 50ms linear 0s;

      &:hover {
        background-color: #efefef;
      }

      &:active {
        top: min(0.5vw, 3.75px);
        box-shadow: none;
      }
    }
  }

  .output {
    grid-area: output;
    border-radius: min(0.75vw, 5.625px);
    margin: min(1vw, 7.5px) min(1vw, 7.5px) min(1vw, 7.5px) min(0.5vw, 3.75px);
    padding: min(1.8vw, 13.5px);
    overflow: auto;
    font-family: sans-serif;
    font-size: min(2.25vw, 16.875px);
    color: #fff;
    background-color: #00385c;
  }
}

:global(.marked > *) {
  /*margin: 0 0 min(4.25vw, 31.875px);*/
  margin: MIN(4.27vw, 32px) 0;
}

:global(.marked > h1) {
  margin: MIN(4vw, 30px) 0 MIN(2vw, 15px);
}

:global(.marked > h2) {
  margin: MIN(4vw, 30px) 0 MIN(2vw, 15px);
  font-size: MIN(3.2vw, 24px);
}

:global(.marked > h3) {
  margin: MIN(4vw, 30px) 0 MIN(2vw, 15px);
}

:global(.marked > h4) {
  margin: MIN(4vw, 30px) 0 MIN(2vw, 15px);
}

:global(.marked > h5) {
  margin: MIN(4vw, 30px) 0 MIN(2vw, 15px);
}

:global(.marked > h6) {
  margin: MIN(4vw, 30px) 0 MIN(2vw, 15px);
}
  
:global(.marked > p > img) {
  display: block;
  max-width: 100%;
  margin:  min(4.25vw, 31.875px) auto;
  object-fit: cover;
  border-radius: min(0.64vw, 4.8px);
}

:global(.marked > blockquote) {
  color: #00add8;
  margin-bottom: MIN(2vw, 15px);
  padding-left: min(3.2vw, 24px);
  border-left: min(1vw, 7.5px) solid #00add8;
}

:global(.marked > ul) {
  /*margin: min(2.13vw, 16px) 0;*/
  margin: MIN(2.13vw, 16px) 0 MIN(2.13vw, 16px) MIN(3.2vw, 24px);
  padding-left: min(3.2vw, 24px);
}

:global(.marked > ol) {
  /*margin: min(2.13vw, 16px) 0;*/
  margin: MIN(2.13vw, 16px) 0 MIN(2.13vw, 16px) MIN(3.2vw, 24px);
  padding-left: min(3.2vw, 24px);
}

:global(.marked > * > li) {
  margin-bottom: MIN(2.13vw, 16px);
}

:global(.marked > * > li > ul, .marked > * > li > ol) {
  margin: 0 0 MIN(2.13vw, 16px) MIN(3.2vw, 24px);
}

:global(.marked > p > a) {
  text-decoration: none;
  color: #00add8;
}

:global(.marked > p > a:hover) {
  text-decoration: underline;
}

:global(.marked > hr) {
  height: MIN(0.4vw, 3px);
  margin: MIN(8vw, 60px) 20%;
  border-color: #00add8;
  border: 0;
}

:global(.marked > pre) {
  margin: min(2.13vw, 16px) 0;
  padding: min(2.13vw, 16px);
  border-radius: min(0.64vw, 4.8px);
  overflow-x: auto;
  background: $bg0_h;
  color: $fg0;
}

:global(.marked > pre code.hljs) {
  display: block;
  overflow-x: auto;
  padding: 1em;
}

:global(.marked code.hljs) {
  padding: 3px 5px;
}

:global(.marked .hljs) {
  background: $bg0_h;
  color: $fg0;
}

:global(.marked .hljs-comment, .hljs-quote) {
  color: $fg3;
  font-style: italic;
}

:global(.marked .hljs-keyword, .hljs-name, .hljs-section, .hljs-selector-tag, .hljs-title) {
  color: $bright_red;
}

:global(.marked .hljs-template-variable, .hljs-variable) {
  color: $bright_red;
}

:global(.marked .hljs-regexp, .hljs-selector-attr, .hljs-selector-pseudo, .hljs-string) {
  color: $bright_green;
}

:global(.marked .hljs-bullet, .hljs-link, .hljs-literal, .hljs-meta, .hljs-params, .hljs-symbol) {
  color: $bright_orange;
}

:global(.marked  .hljs-number) {
  color: $bright_purple;
}

:global(.marked .hljs-attr, .hljs-built_in, .hljs-doctag, .hljs-title, .hljs-type) {
  color: $bright_green;
}

:global(.marked .hljs-attribute, .hljs-subst) {
  color: $bg0_h;
}

:global(.marked .hljs-formula) {
  background-color: $fg;
  font-style: italic;
}

:global(.marked .hljs-selector-class, .hljs-selector-id) {
  color: $bright_orange;
}

:global(.marked .hljs-addition) {
  background-color: $bright_green;
}

:global(.marked .hljs-deletion) {
  background-color: $bright_red;
}

:global(.marked .hljs-doctag, .hljs-strong) {
  font-weight: 700;
}

:global(.marked .hljs-emphasis) {
  font-style: italic;
}

:global(.marked > p > code) {
  border-radius: min(0.64vw, 4.8px);
  overflow-x: auto;
  background: #171717;
  color: #f8f8f2;
  margin: MIN(4vw, 30px) 0; /*new*/
  padding: min(0.3vw, 2.25px) min(0.5vw, 3.75px);
  font-size: min(2vw, 14px);
}
</style>
