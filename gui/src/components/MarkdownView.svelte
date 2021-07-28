<script>
import marked from 'marked';
import { onDestroy } from 'svelte';
import { todo } from "../stores/todo.js";

// TODO: Add markdown codeblock syntax highlighting
/*
marked.setOptions({
  highlight: function(code, lang, callback) {
    require('pygmentize-bundled') ({ lang: lang, format: 'html'  }, code, function (err, result) {
      callback(err, result.toString());
    });
  }
});
*/
let placeholder;
const unsubscribe = todo.subscribe(todo => {
  placeholder = todo.body;
})

onDestroy(unsubscribe);
</script>

<div class="markdown">
  <textarea bind:value={$todo.body} id="textArea" class="input" placeholder={placeholder} wrap="off"></textarea>
  <div class="output"><div class="marked">{@html marked($todo.body)}</div></div>
</div>

<style lang="scss">
.markdown {
  height: 94.2vh;
  display: grid;
  grid-template-columns: 1fr 1fr;
  grid-template-areas: "input output";
  border-radius: min(0.75vw, 5.625px) 0 0 min(0.75vw, 5.625px);
  background-color: #00385c;

  .input {
    grid-area: input;
    outline: none;
    border: none;
    border-radius: min(0.75vw, 5.625px);
    margin: min(1vw, 7.5px) min(0.5vw, 3.75px) min(1vw, 7.5px) min(1vw, 7.5px);
    padding: min(1.8vw, 13.5px);
    overflow: auto;
    resize: none;
    font-size: min(2.25vw, 16.875px);
    font-family: monospace;
    color: #00385c;
    background-color: #fff;
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
  margin: 0 0 min(4.25vw, 31.875px);
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
  padding-left: min(3.2vw, 24px);
  border-left: min(1vw, 7.5px) solid #00add8;
}

:global(.marked > ul, ol) {
  margin: min(2.13vw, 16px) 0;
  padding-left: min(3.2vw, 24px);
}

:global(.marked > ol > li) {
  margin-left: min(2vw, 14px);
}

:global(.marked > p > a) {
  text-decoration: none;
  color: #00add8;
}

:global(.marked > p > a:hover) {
  text-decoration: underline;
}

:global(.marked > hr) {
  border-color: #00add8;
}

:global(.marked > pre) {
  margin: min(2.13vw, 16px) 0;
  padding: min(2.13vw, 16px);
  border-radius: min(0.64vw, 4.8px);
  overflow-x: auto;
  background: #171717;
  color: #f8f8f2;
}

:global(.marked > p > code) {
  border-radius: min(0.64vw, 4.8px);
  overflow-x: auto;
  background: #171717;
  color: #f8f8f2;
  padding: min(0.3vw, 2.25px) min(0.5vw, 3.75px);
  font-size: min(2vw, 14px);
}
</style>
