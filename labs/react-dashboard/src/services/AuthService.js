import Cookies from "js-cookie";

const AUTH_TOKEN_KEY = "TOKEN";
const AUTH_TOKEN_EXPIRY_DAYS = 1;

const AuthService = {
  storeToken(token) {
    Cookies.set(AUTH_TOKEN_KEY, token, { expires: AUTH_TOKEN_EXPIRY_DAYS });
    return true;
  },

  getToken() {
    return Cookies.get(AUTH_TOKEN_KEY);
  },

  isAuthorized() {
    return Boolean(Cookies.get(AUTH_TOKEN_KEY));
  },

  clearToken() {
    Cookies.remove(AUTH_TOKEN_KEY);
  },
};

export default AuthService;
