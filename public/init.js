const AppState = new StateManager();

document.addEventListener("DOMContentLoaded", () => {
  AppState.set(SELECTED_NAV_OPTION, 0);
});
