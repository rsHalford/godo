<script>
import { todo } from "../stores/todo.js";
import { login } from "../stores/login.js";
import { createEventDispatcher } from 'svelte';

const dispatch = createEventDispatcher();

function saveTodo() {
  // TODO: Refresh sidebar items in case title changes
  let updatedTodo = {
    ID: $todo.ID,
    title: $todo.title,
    body: $todo.body,
    priority: $todo.priority,
    status: $todo.status,
  }
  let jsonTodo = JSON.stringify(updatedTodo);
  var xhr = new XMLHttpRequest();
  xhr.open("PUT", $login.url+"/"+$todo.ID);
  xhr.setRequestHeader("Accept", "*/*");
  xhr.setRequestHeader("Authorization", "Basic " + btoa($login.username + ":" + $login.password));
  xhr.send(jsonTodo);
  dispatch('refresh');
}

function deleteTodo() {
  if (confirm('Confirm deletion of "'+$todo.title+'"?')) {
    // TODO: Refresh sidebar items in case title changes
    var xhr = new XMLHttpRequest();
    xhr.open("DELETE", $login.url+"/"+$todo.ID);
    xhr.setRequestHeader("Accept", "*/*");
    xhr.setRequestHeader("Authorization", "Basic " + btoa($login.username + ":" + $login.password));
    xhr.send();
    dispatch('refresh');
  }
}
</script>

<header class="toolbar">
  <div class="title">
    <input bind:value={$todo.title} type="text" placeholder="Name your todo"/>
  </div>
  <div class="tools">
    <div class="prioritise">
      <label>
        <input type="checkbox" bind:checked={$todo.priority}/>
        <i class="fas fa-star unchecked"></i>
        <i class="fas fa-star checked"></i>
      </label>
    </div>
    <div class="done">
      <label>
        <input type="checkbox" bind:checked={$todo.status}/>
        <i class="fas fa-check-circle unchecked"></i>
        <i class="fas fa-check-circle checked"></i>
      </label>
    </div>
    <div class="save">
      <button on:click={saveTodo}><i class="fas fa-save"></i></button>
    </div>
    <div class="delete">
      <button on:click={deleteTodo}><i class="fas fa-trash-alt"></i></button>
    </div>
  </div>
</header>

<style lang="scss">
.toolbar {
  flex-shrink: 0.06;
  display: grid;
  grid-template-columns: 1fr 0.6fr;
  grid-template-areas: "input tools";
  background-color: #00add8;

  .title {
    margin: min(1vw, 7.5px) min(0.5vw, 3.75px) min(1vw, 7.5px) min(1vw, 7.5px);
  
    input {
      grid-area: input;
      width: 100%;
      font-size: min(2.3vw, 17.25px);
      outline: none;
      border: none;
      border-radius: min(0.75vw, 5.625px);
      padding: min(1.4vw, 10.5px);
      background: #fff;
      color: #00385c;
    }
  }

  .tools {
    grid-area: tools;
    display: flex;
    flex-flow: row nowrap;
    justify-content: flex-end;

    .prioritise {
      grid-area: prioritise;
      margin: min(1vw, 7.5px) min(0.5vw, 3.75px);
    }

    .done {
      grid-area: done;
      margin: min(1vw, 7.5px) min(0.5vw, 3.75px);
    }

    .save {
      grid-area: save;
      margin: min(1vw, 7.5px) min(0.5vw, 3.75px);
    }

    .delete {
      grid-area: delete;
      margin: min(1vw, 7.5px) min(1vw, 7.5px) min(1vw, 7.5px) min(0.5vw, 3.75px);
    }

    div {

      label {
        display: block;
        position: relative;
        top: 0;
        width: min(7vw, 52.5px);
        padding: min(1.133vw, 8.5px) min(1.4vw, 10.5px);
        cursor: pointer;
        font-size: min(2.8vw, 21px);
        text-align: center;
        color: #00385c;
        background-color: #fff;
        border-radius: min(0.75vw, 5.625px);
        border: none;
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
        user-select: none;
        box-shadow: 0 min(0.5vw, 3.75px) 0 #00385c;
        transition: all 0.1s linear 0s;

        &:hover {
          background-color: #efefef;
        }

        &:active {
          top: min(0.5vw, 3.75px);
          box-shadow: none;
        }

        .checked {
          display: none;
        }

        input[type="checkbox"] {
          display: none;

          &:checked ~ .fa-star.checked {
            display: inline-block;
            color: gold;
          }

          &:checked ~ .fa-check-circle.checked {
            display: inline-block;
            color: green;
          }

          &:checked ~ .unchecked {
            display: none;
            color: #fff;
          }
        }
      }

      button {
        position: relative;
        top: 0;
        width: min(7vw, 52.5px);
        padding: min(1vw, 7.5px) min(1.4vw, 10.5px);
        cursor: pointer;
        font-size: min(2.8vw, 21px);
        color: #00385c;
        background-color: #fff;
        border-radius: min(0.75vw, 5.625px);
        border: none;
        box-shadow: 0 min(0.5vw, 3.75px) 0 #00385c;
        transition: all 0.1s linear 0s;

        &:hover {
          background-color: #efefef;
        }

        &:active {
          top: min(0.5vw, 3.75px);
          box-shadow: none;
        }
      }
    }
  }
}
</style>
