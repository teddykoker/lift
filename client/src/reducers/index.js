import { APP_LOAD, LOGIN, SIGNUP } from "../constants/actionTypes";

const defaultState = {
  token: null,
  error: null,
  loaded: false
};

export default (state = defaultState, action) => {
  switch (action.type) {
    case APP_LOAD:
      return {
        ...state,
        token: action.token || null,
        loaded: true
      };
    case SIGNUP:
    case LOGIN:
      return {
        ...state,
        token: action.payload.token,
        error: action.payload.error || null
      };
    default:
      return state;
  }
};
