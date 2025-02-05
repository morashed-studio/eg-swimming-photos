// some constants used as state values keys for consistency
// and for reflecting the application concept in code 
const SELECTED_NAV_OPTION = "sn";

NAV_OPTIONS = {
  HOME: 0,
  MY_CART: 1,
  CONTACT_US: 2,
}

class StateManager {
  _stateMap = {};
  _listeners = {};

  get(key) {
    return this._stateMap[key];
  }

  set(key, value) {
    this._stateMap[key] = value;
    if (this._listeners[key]) {
      for (const func of this._listeners[key]) func();
    }
  }

  addListener(key, func) {
    if (!this._listeners[key]) {
      this._listeners[key] = [];
    }
    this._listeners[key].push(func);
  }
}

const AppState = new StateManager();

AppState.addListener(SELECTED_NAV_OPTION, () => {
  const btns = document.getElementsByClassName("nav-btn");
  for (const btn of btns) {
    if (NAV_OPTIONS[btn.dataset.option] == AppState.get(SELECTED_NAV_OPTION)) {
      btn.classList.add("bg-secondary", "color-primary");
    }
    else {
      btn.classList.remove("bg-secondary", "color-primary");
    }
  }
})

document.addEventListener("DOMContentLoaded", () => {
  AppState.set(SELECTED_NAV_OPTION, NAV_OPTIONS.HOME);
});
