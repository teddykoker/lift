import { APP_LOAD, LOGIN, SIGNUP, LOGOUT } from "../constants/actionTypes";
import api from "../api";

export const onLoad = (token, user) => ({
  type: APP_LOAD,
  token,
  user
});

export const login = payload => ({
  type: LOGIN,
  payload
});

export const signup = payload => ({
  type: SIGNUP,
  payload
});

export const logout = () => {
  localStorage.removeItem("jwt");
  api.setToken(null);
  return {
    type: LOGOUT
  };
};
