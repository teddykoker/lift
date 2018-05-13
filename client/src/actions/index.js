import { APP_LOAD, LOGIN } from "../constants/actionTypes";

export const onLoad = (token, user) => ({
  type: APP_LOAD,
  token,
  user
});

export const login = payload => ({
  type: LOGIN,
  payload
});
