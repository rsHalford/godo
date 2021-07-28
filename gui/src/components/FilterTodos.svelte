<script>
import { createEventDispatcher } from 'svelte';

const dispatch = createEventDispatcher();
const done = () => dispatch('done');
const priority = () => dispatch('priority');

function filterDropdown() {
  let isFilter = document.getElementById("filterCheckbox").checked;
  if (isFilter === true) {
    document.getElementById("filterTitle").style.marginBottom = 0;
    document.getElementById("filterTitle").style.borderRadius = "min(0.75vw, 5.625px) min(0.75vw, 5.625px) 0 0";
    document.getElementById("optionsMenu").style.display = "block";
  } else {
    document.getElementById("filterTitle").style.marginBottom = "min(1vw, 7.5px)";
    document.getElementById("filterTitle").style.borderRadius = "min(0.75vw, 5.625px)";
    document.getElementById("optionsMenu").style.display = "none";
  }
}
</script>

<div class="filter">
  <div class="filter-title" id="filterTitle">
    <label>
      <input id="filterCheckbox" type="checkbox" on:click={filterDropdown}/>
      FILTER
      <i class="fas fa-caret-right"></i>
    </label>
  </div>
  <div class="options" id="optionsMenu">
    <div class="priority">
      <label>
        <input type="checkbox" on:click={priority}/>
        <i class="fas fa-star unchecked"></i>
        <i class="fas fa-star checked"></i>
        &nbsp;PRIORITY
      </label>
    </div>
    <div class="done">
      <label>
        <input type="checkbox" on:click={done} />
        <i class="fas fa-check-circle unchecked"></i>
        <i class="fas fa-check-circle checked"></i>
        &nbsp;DONE
      </label>
    </div>
  </div>
</div>

<style lang="scss">
.filter {
  position: relative;
  display: inline-block;
  
  .filter-title {
    width: 100%;
    margin-bottom: min(1vw, 7.5px);
    padding: min(1.8vw, 13.5px) min(3vw, 22.5px) min(1.8vw, 13.5px);
    font-size: min(3vw, 22.5px);
    color: #00add8;
    background-color: #00385c;
    border: none;
    border-radius: min(0.75vw, 5.625px);
    transition: opacity 0.5s ease-in-out;

    .fa-caret-right {
      margin-left: min(7.6vw, 57px);
      transition: 0.2s ease-in-out;
    }

    &:hover {
      .fa-caret-right {
        color: #fff;
      }
    }

    label {
      display: block;
      cursor: pointer;
      font-family: "Lexend", sans-serif;
      font-weight: 400;
      -webkit-user-select: none;
      -moz-user-select: none;
      -ms-user-select: none;
      user-select: none;

      input[type="checkbox"] {
        display: none;

        &:checked ~ .fa-caret-right {
          display: inline-block;
          transform: rotate(90deg);
        }
      }
    }
  }

  .options {
    display: none;

    .priority {
      border-radius: 0;
      padding: 0 min(4.5vw, 33.75px) min(2.2vw, 16.5px);
    }

    .done {
      padding: 0 min(4.653vw, 34.8975px) min(2.2vw, 16.5px);
      margin-bottom: min(1vw, 7.5px);
      border-radius: 0 0 min(0.75vw, 5.625px) min(0.75vw, 5.625px);
    }

    div {
      display: inline-block;
      width: 100%;
      font-size: min(2.5vw, 18.75px);
      font-family: "Lexend", sans-serif;
      font-weight: 400;
      color: #fff;
      background-color: #00385c;
      border: none;
      border-radius: min(0.75vw, 5.625px);

      label {
        display: block;
        cursor: pointer;
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
        user-select: none;

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

        &:hover {
          
          .fa-star {
            color: gold;
          }

          .fa-check-circle {
            color: green;
          }
        }
      }
    }
  }
}
</style>
