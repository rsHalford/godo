<script>
import { onMount } from "svelte";
import { onDestroy } from 'svelte';
import { todo } from "./stores/todo.js";
import { login } from "./stores/login.js";
import AddTodo from './components/AddTodo.svelte';
import FilterTodos from './components/FilterTodos.svelte';
import Login from './components/Login.svelte';
import MarkdownView from './components/MarkdownView.svelte';
import Toolbar from './components/Toolbar.svelte';

let todos = [];
let isDone = false;
let isPriority = false;
let url;
let username;
let password;

const unsubscribe = login.subscribe(login => {
  url = login.url;
  username = login.username;
  password = login.password;
})

function getTodos(url, username, password) {
  var xhr = new XMLHttpRequest();
  xhr.open("GET", url);
  xhr.setRequestHeader("Accept", "*/*");
  xhr.setRequestHeader("Authorization", "Basic " + btoa(username + ":" + password));
  xhr.onreadystatechange = function () {
    if (this.readyState == 4 && this.status == 200) {
      let json = this.responseText;
      todos = JSON.parse(json);
    }
  };
  xhr.send();
}

// TODO: onMount if login info exists getTodos()
onMount(() => getTodos(url, username, password));

function refreshTodos() {
  getTodos(url, username, password);
}

function selectTodo(todoObject) {
  todo.set(todoObject);
}

function toggleDone() {
  isDone = !isDone;
  console.log(isDone);
}

function togglePriority() {
  isPriority = !isPriority;
  console.log(isPriority);
}

onDestroy(unsubscribe);
</script>

<div class="container">
{#if todos == ""}
  <Login on:refresh={refreshTodos}/>
{:else}
  <div class="sidebar">
    <AddTodo on:refresh={refreshTodos}/>
    <FilterTodos on:done={toggleDone} on:priority={togglePriority}/>
    {#if isDone == true && isPriority == true}
      <ul>
      {#each todos as todo}
        {#if todo.status == true && todo.priority == true}
          <li><button class="priority done" on:click={() => selectTodo(todo)}>{todo.title}</button></li>
        {/if}
      {/each}
      </ul>
    {:else if isDone == true && isPriority == false}
      <ul>
      {#each todos as todo}
        {#if todo.status == true && todo.priority == true}
        <li><button class="priority done" on:click={() => selectTodo(todo)}>{todo.title}</button></li>
        {:else if todo.status == true}
        <li><button class="done" on:click={() => selectTodo(todo)}>{todo.title}</button></li>
        {/if}
      {/each}
      </ul>
    {:else if isDone == false && isPriority == true}
      <ul>
      {#each todos as todo}
        {#if todo.status == false && todo.priority == true}
        <li><button class="priority" on:click={() => selectTodo(todo)}>{todo.title}</button></li>
        {/if}
      {/each}
      </ul>
    {:else}
      <ul>
      {#each todos as todo}
        {#if todo.status == false && todo.priority == true}
        <li><button class="priority" on:click={() => selectTodo(todo)}>{todo.title}</button></li>
        {:else if todo.status == false}
        <li><button on:click={() => selectTodo(todo)}>{todo.title}</button></li>
        {/if}
      {/each}
      </ul>
    {/if}
  </div>
  <div class="main">
    <Toolbar on:refresh={refreshTodos}/>
    <MarkdownView />
  </div>
{/if}
</div>

<style lang="scss">
.container {
  height: 100vh;
  width: 100vw;
  display: grid;
  grid-template-columns: min(30vw, 225px) auto;
  grid-template-areas:
    "sidebar main";

  .main {
    grid-area: main;
    display: flex;
    flex-flow: column nowrap;
    justify-content: flex-start;
    background-color: #00add8;
  }

  .sidebar {
    grid-area: sidebar;
    overflow: auto;
    padding: min(1vw, 7.5px) min(1vw, 7.5px) 0 min(1vw, 7.5px);
    display: flex;
    flex-flow: column nowrap;
    justify-content: flex-start;
    background-color: #00add8;
  
    ul {
      li {
        button {
          width: 100%;
          cursor: pointer;
          text-overflow: ellipsis;
          font-size: min(2.25vw, 16.875px);
          color: #00385c;
          background-color: #fff;
          margin-bottom: min(1vw, 7.5px);
          padding: min(1.8vw, 13.5px);
          border-radius: min(0.75vw, 5.625px);
          border: none;
  
          &:hover {
            background-color: #efefef;
          }
        }
  
        .priority {
          border: min(0.4vw, 3px) solid gold;
        }
  
        .done {
          background-color: #bbb;
  
          &:hover {
            background-color: #ccc;
          }
        }
      }
    }
  }
}
</style>
