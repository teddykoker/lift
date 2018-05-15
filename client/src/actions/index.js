import { APP_LOAD, LOGIN, SIGNUP } from "../constants/actionTypes";

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
