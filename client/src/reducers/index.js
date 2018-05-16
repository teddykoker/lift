import { APP_LOAD, LOGIN, SIGNUP } from "../constants/actionTypes";

const defaultState = {
  user: null,
  error: null,
  loaded: false
};

export default (state = defaultState, action) => {
  switch (action.type) {
    case APP_LOAD:
      return {
        ...state,
        user: action.user || null,
        loaded: true
      };
    case SIGNUP:
    case LOGIN:
      return {
        ...state,
        user: action.payload || null,
        error: action.payload.error || null
      };
    default:
      return state;
  }
};
