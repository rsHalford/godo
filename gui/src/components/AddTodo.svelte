<script>
import { createEventDispatcher } from 'svelte';
import { login } from "../stores/login.js";

const dispatch = createEventDispatcher();

function openTodoForm() {
  document.getElementById("todoForm").style.display = "block";
}

function closeTodoForm() {
  document.getElementById("todoForm").style.display = "none";
}

let newTodo = {};
let jsonNewTodo;

function addTodo() {
  newTodo = {
    title: document.getElementById("todoTitle").value,
    body: document.getElementById("todoBody").value,
    priority: document.getElementById("todoPriority").checked,
  }
  jsonNewTodo = JSON.stringify(newTodo);
  let xhr = new XMLHttpRequest();
  xhr.open("POST", $login.url);
  xhr.setRequestHeader("Accept", "*/*");
  xhr.setRequestHeader("Content-Type", "application/json");
  xhr.setRequestHeader("Authorization", "Basic " + btoa($login.username + ":" + $login.password));
  xhr.onreadystatechange = function () {
    if (this.readyState == 4 && this.status == 200) {
      console.log(this.status);
      console.log(this.responseText);
    }
  };
  xhr.send(jsonNewTodo);
  dispatch('refresh');
  closeTodoForm();
}
</script>

<div class="godo">
  <button class="logo-button" on:click={openTodoForm}><span class="plus"> </span><span class="go">Go</span>Do</button>
  <div class="overlay" id="todoForm">
    <div class="todo-form">
      <form class="form-container">
        <div class="header">
          <div class="priority">
            <label>
              <input id="todoPriority" type="checkbox"/>
              <i class="fas fa-star"></i>
            </label>
          </div>
          <h1>Go<span>Do</span></h1>
          <button class="close-form" type="button" on:click={closeTodoForm}><i class="fas fa-times-circle"></i></button>
        </div>
        <div class="inputs">
          <label for="title"></label>
          <input id="todoTitle" type="text" placeholder="Name your todo"/>
          <label for="body"></label>
          <textarea id="todoBody" class="input" placeholder="Fill in the body of your todo using markdown" wrap="off"></textarea>
        </div>
        <div class="buttons">
          <button class="form-button add" type="button" on:click={addTodo}>ADD</button>
          <button class="form-button clear" type="reset">CLEAR</button>
        </div>
      </form>
    </div>
  </div>
</div>

<style lang="scss">
.godo {
  z-index: 1;

  .logo-button {
    width: 100%;
    cursor: pointer;
    margin-bottom: min(1vw, 7.5px);
    padding: min(1.8vw, 13.5px);
    font-size: min(5.333vw, 40px);
    font-family: "Lexend", sans-serif;
    font-weight: 400;
    color: #00add8;
    background-color: #00385c;
    border: none;
    border-radius: min(0.75vw, 5.625px);
    transition: opacity 0.2s ease-in-out;

    .go {
      color: #fff;
    }

    &:hover {

      .plus {
        padding-left: min(3.333vw, 25px);

        &:after {
          opacity: 1;
          left: 0;
        }
      }
    }
  
    .plus {
      color: #fff;
      position: relative;
      transition: 0.2s;

      &:after {
        content: '\002b';
        position: absolute;
        opacity: 0;
        top: -5%;
        left: min(-2vw, -15px);
        transition: 0.2s;
      }
    }
  
    &:active {
      background-color: #01476e;
    }
  }

  .overlay {
    height: 100vh;
    width: 100vw;
    position: absolute;
    top: 0;
    left: 0;
    display: none;
    backdrop-filter: blur(8px);

    .todo-form {
      display: block;
      position: fixed;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      border-radius: min(0.75vw, 5.625px) 0 0 min(0.75vw, 5.625px);

      form {
        width: min(90vw, 675px);
        padding: min(1.8vw, 13.5px);
        color: #00385c;
        background-color: #00add8;
        border-radius: min(0.75vw, 5.625px);

        .header {
          display: grid;
          grid-template-columns: min-content auto min-content;
          grid-template-areas: "priority title close";
          margin-bottom: min(1vw, 7.5px);

          .priority {
            grid-area: priority;
            margin: auto 0;

            label {
              display: block;
              position: relative;
              cursor: pointer;
              font-size: min(4vw, 30px);
              text-align: center;
              -webkit-user-select: none;
              -moz-user-select: none;
              -ms-user-select: none;
              user-select: none;

              input[type="checkbox"] {
                display: none;

                &:checked ~ .fa-star {
                  display: inline-block;
                  color: gold;
                }
              }

              i {
                transition: 0.2s ease-in-out;

                &:hover {
                  color: #2e4e6f;
                }
              }
            }
          }

          h1 {
            grid-area: title;
            text-align: center;
            font-size: min(6vw, 45px);
            font-family: "Lexend", sans-serif;
            font-weight: 400;
            color: #fff;

            span {
              color: #00385c;
            }
          }

          .close-form {
            height: min(4vw, 30px);
            width: min(4vw, 30px);
            grid-area: close;
            border: none;
            border-radius: min(0.75vw, 5.625px);
            outline: none;
            cursor: pointer;
            font-size: min(3vw, 22.5px);
            color: #00385c;
            background-color: #00add8;
            transition: 0.2s ease-in-out;

            &:hover {
              color: #fff;
            }
          }
        }

        .inputs {
          display: grid;
          grid-template-areas:
            "title"
            "body";

          input[type=text] {
            width: 100%;
            margin: min(1vw, 7.5px) 0;
            padding: min(1.8vw, 13.5px);
            outline: none;
            border: none;
            border-radius: min(0.75vw, 5.625px);
            font-size: min(2.25vw, 16.875px);
            background-color: #fff;
          }

          .input {
            height: 50vh;
            width: 100%;
            outline: none;
            margin: min(1vw, 7.5px) 0;
            padding: min(1.8vw, 13.5px);
            border: none;
            border-radius: min(0.75vw, 5.625px);
            overflow: auto;
            resize: none;
            font-size: min(2.25vw, 16.875px);
            color: #00385c;
            background-color: #fff;
          }
        }

        .buttons {
          display: grid;
          grid-template-columns: auto min-content min-content;
          grid-template-areas: ". add clear";
          grid-gap: min(2vw, 15px);
          margin: min(1vw, 7.5px) 0 min(0.5vw, 3.75px);

          .form-button {
            position: relative;
            top: 0;
            width: 100%;
            cursor: pointer;
            font-size: min(3vw, 22.5px);
            font-family: "Lexend", sans-serif;
            font-weight: 400;
            padding: min(2vw, 15px);
            color: #00385c;
            border: none;
            border-radius: min(0.75vw, 5.625px);
            background-color: #fff;
            box-shadow: 0 min(0.5vw, 3.75px) 0 #00385c;
            transition: all 0.1s linear 0s;

            &:active {
              top: min(0.5vw, 3.75px);
              box-shadow: none;
            }
          }

          .add {
            grid-area: add;
            width: min(20vw, 150px);
            background-color: #fff;

            &:hover {
              background-color: #efefef;
            }
          }

          .clear {
            grid-area: clear;
            width: min(20vw, 150px);
            background-color: #bbb;

            &:hover {
              background-color: #ccc;
            }
          }
        }
      }
    }
  }
}
</style>
