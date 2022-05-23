import { toast } from "react-toastify";
import { LIGHT_MODE, DARK_MODE } from "../constants";

const checkTheme = () => {
  if (
    window.localStorage.theme === "dark" ||
    (!("theme" in window.localStorage) &&
      window.matchMedia("(prefers-color-scheme: dark)").matches)
  ) {
    document.documentElement.classList.add("dark");
    return "dark";
  } else {
    document.documentElement.classList.remove("dark");
    return "light";
  }
};

interface IThemeState {
  theme: "dark" | "light" | (() => "dark" | "light");
}

const initialState: IThemeState = {
  theme: checkTheme(),
};

const UiReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case DARK_MODE:
      window.localStorage.theme = "dark";
      document.documentElement.classList.add("dark");
      toast.success("Dark mode set");
      return { ...state, theme: "dark" };

    case LIGHT_MODE:
      window.localStorage.theme = "light";
      document.documentElement.classList.remove("dark");
      toast.success("Light mode set");
      return { ...state, theme: "light" };

    default:
      return state;
  }
};

export default UiReducer;
