import { APP_LOAD } from "../constants/actionTypes";

export const onLoad = token => ({
  type: APP_LOAD,
  token
});
