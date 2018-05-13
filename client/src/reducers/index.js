import { APP_LOAD, LOGIN } from "../constants/actionTypes";

const defaultState = {
  token: null,
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
    case LOGIN:
      return {
        ...state,
        token: action.payload.token
      };
    default:
      return state;
  }
};
