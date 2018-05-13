import { APP_LOAD } from "../constants/actionTypes";

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
    default:
      return state;
  }
};
