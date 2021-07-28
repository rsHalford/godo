<script>
import { createEventDispatcher } from 'svelte';
import { login } from "../stores/login.js";

const dispatch = createEventDispatcher();

function getFileContents() {
  document.getElementById("read-button").addEventListener('click', function () {
    let file = document.getElementById("file-input").files[0];
    let reader = new FileReader();
    reader.addEventListener('load', function(e) {
      let text = e.target.result;
      let userRegex = /username: ?\s*"(\S+)"/g;
      let passRegex = /password: ?\s*"(\S+)"/g;
      let urlRegex = /api: ?\s*"(\S+)"/g;
      let username = userRegex.exec(text)[1];
      let password = passRegex.exec(text)[1];
      let url = urlRegex.exec(text)[1];
      login.set({
        url: url,
        username: username,
        password: password,
      });
      dispatch('refresh');
    });
    reader.readAsText(file);
  });
}
</script>

<div class="login" id="loginForm">
  <div class="overlay">
    <div class="login-form">
      <form class="form-container" name="loginForm">
        <div class="header">
          <h1>Go<span>Do</span></h1>
          <h2>Select config.yaml</h2>
        </div>
        <div class="inputs">
          <input class="input" id="file-input" type="file" on:click={getFileContents}>
        </div>
        <div class="button">
          <button class="form-button" type="button" id="read-button">Submit</button>
        </div>
      </form>
    </div>
  </div>
</div>

<style lang="scss">
.login {
  z-index: 2;

  .overlay {
    height: 100vh;
    width: 100vw;
    position: absolute;
    top: 0;
    left: 0;
    backdrop-filter: blur(8px);

    .login-form {
      display: block;
      position: fixed;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      border-radius: min(0.75vw, 5.625px) 0 0 min(0.75vw, 5.625px);

      form {
        padding: min(1.8vw, 13.5px);
        color: #00add8;
        background-color: #00385c;
        border-radius: min(0.75vw, 5.625px);

        .header {
          display: grid;
          grid-template-rows: auto auto;
          grid-template-areas:
            "title"
            "subtitle";
          margin-bottom: min(1vw, 7.5px);

          h1 {
            grid-area: title;
            text-align: center;
            font-size: min(6vw, 45px);
            font-family: "Lexend", sans-serif;
            font-weight: 400;
            color: #fff;

            span {
              color: #00aad8;
            }
          }

          h2 {
            grid-area: subtitle;
            text-align: center;
            font-size: min(3vw, 22.5px);
            font-family: "Lexend", sans-serif;
            font-weight: 400;
            color: #fff;
          }
        }

        .inputs {

          .input {
            width: 100%;
            margin: min(1vw, 7.5px) auto;
            padding: min(1.8vw, 13.5px);
            outline: none;
            border: none;
            border-radius: min(0.75vw, 5.625px);
            font-size: min(2.25vw, 16.875px);
            color: #fff;
          }
        }

        .button {
          text-align: center;

          .form-button {
            top: 0;
            cursor: pointer;
            font-size: min(2.5vw, 18.75px);
            font-family: "Lexend", sans-serif;
            font-weight: 400;
            margin-top: min(1vw, 7.5px);
            padding: min(1vw, 7.5px) min(2vw, 15px);
            color: #00385c;
            border: none;
            border-radius: min(0.75vw, 5.625px);
            background-color: #fff;
            box-shadow: 0 min(0.5vw, 3.75px) 0 #001a3a;
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
  }
}
</style>
